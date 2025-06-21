package repository

import (
	"context"
	"database/sql"
	"eliborate/internal/models/entity"
	"eliborate/internal/repository/repoutils"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
)

type bookRepo struct {
	db     *sqlx.DB
	search meilisearch.IndexManager
}

func InitBookRepo(db *sqlx.DB, search meilisearch.IndexManager) BookRepo {
	return bookRepo{
		db:     db,
		search: search,
	}
}

func (b bookRepo) CreateBook(ctx context.Context, book entity.BookCreate) (int, error) {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	query, args, err := qbuilder.
		Insert("books").
		Columns("title", "authors", "description", "category_id", "cover_urls", "rack", "shelf").
		Values(book.Title, book.Authors, book.Description, book.CategoryID, book.CoverUrls, book.Rack, book.Shelf).
		Suffix("RETURNING \"id\"").
		ToSql()
	if err != nil {
		return 0, err
	}

	row := tx.QueryRowContext(ctx, query, args...)

	var bookID int
	if err = row.Scan(&bookID); err != nil {
		tx.Rollback()
		return 0, err
	}

	bookSearch := repoutils.ConvertEntityBookSearchFromEntityBookCreate(bookID, book)

	if _, err = b.search.AddDocumentsWithContext(ctx, []entity.BookSearch{bookSearch}); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return bookID, nil
}

func (b bookRepo) GetBookById(ctx context.Context, id int) (entity.Book, error) {
	query, args, err := qbuilder.
		Select("b.id", "b.title", "b.description", "c.name", "b.authors", "b.cover_urls", "b.rack", "b.shelf").
		From("books b").
		Join("categories c ON b.category_id = c.id").
		Where(squirrel.Eq{"b.id": id}).
		ToSql()
	if err != nil {
		return entity.Book{}, err
	}

	row := b.db.QueryRowContext(ctx, query, args...)

	var book entity.Book
	err = row.Scan(&book.ID, &book.Title, &book.Description, &book.Category, &book.Authors, &book.CoverUrls, &book.Rack, &book.Shelf)
	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (b bookRepo) GetBooks(ctx context.Context, offset, limit int, rack *int, searchQuery *string) ([]entity.Book, error) {
	if searchQuery != nil {
		indices, err := b.getBooksIndicesByTextSearch(ctx, *searchQuery, offset, limit, rack)
		if err != nil {
			return []entity.Book{}, err
		}
		return b.getBooksByIndices(ctx, indices)
	}

	baseQuery := qbuilder.
		Select("b.id", "b.title", "b.description", "c.name", "b.authors", "b.cover_urls", "b.rack", "b.shelf").
		From("books b").
		Join("categories c ON b.category_id = c.id")
	if rack != nil {
		baseQuery = baseQuery.Where(squirrel.Eq{"b.rack": *rack})
	}

	query, args, err := baseQuery.
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		return []entity.Book{}, err
	}

	res, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return []entity.Book{}, err
	}
	defer res.Close()

	var books []entity.Book
	for res.Next() {
		var book entity.Book
		if err := res.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Category,
			&book.Authors,
			&book.CoverUrls,
			&book.Rack,
			&book.Shelf,
		); err != nil {
			return []entity.Book{}, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b bookRepo) GetBooksTotalCount(ctx context.Context) (int, error) {
	var totalCount int
	if err := b.db.GetContext(ctx, &totalCount, `SELECT COUNT(*) FROM books`); err != nil {
		return 0, err
	}
	return totalCount, nil
}

func (b bookRepo) getBooksIndicesByTextSearch(ctx context.Context, query string, offset, limit int, rack *int) ([]int, error) {
	searchRequest := &meilisearch.SearchRequest{
		Limit:  int64(limit),
		Offset: int64(offset),
	}
	if rack != nil {
		searchRequest.Filter = fmt.Sprintf("rack = %d", *rack)
	}

	searchResp, err := b.search.SearchWithContext(ctx, query, searchRequest)
	if err != nil {
		return nil, err
	}

	return repoutils.ConvertMeiliHitsToIntSlice(searchResp.Hits), nil
}

func (b bookRepo) getBooksByIndices(ctx context.Context, indices []int) ([]entity.Book, error) {
	if len(indices) == 0 {
		return []entity.Book{}, nil
	}

	query, args, err := qbuilder.
		Select("b.id", "b.title", "b.description", "c.name", "b.authors", "b.cover_urls", "b.rack", "b.shelf").
		From("books b").
		LeftJoin("categories c ON c.id = b.category_id").
		Where(squirrel.Eq{"b.id": indices}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := b.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.Category,
			&book.Authors,
			&book.CoverUrls,
			&book.Rack,
			&book.Shelf,
		); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b bookRepo) UpdateBookInfo(ctx context.Context, id int, fields map[string]interface{}) error {
	builder := qbuilder.Update("books")

	for field, value := range fields {
		builder = builder.Set(field, value)
	}

	query, args, err := builder.
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (b bookRepo) UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error {
	query, args, err := qbuilder.
		Update("books").
		Set("rack", rack).
		Set("shelf", shelf).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (b bookRepo) DeleteBook(ctx context.Context, id int) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, `DELETE FROM books WHERE id=$1`, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		tx.Rollback()
		return sql.ErrNoRows
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
