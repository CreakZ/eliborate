package repository

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/errs"
	"eliborate/internal/models/entity"
	"fmt"
	"strings"

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

	var categoryId int
	row := tx.QueryRowContext(ctx, `SELECT id FROM categories WHERE name = $1`, book.Category)
	if err = row.Scan(&categoryId); err != nil {
		return 0, err
	}

	query := `--sql
	INSERT INTO books (title, authors, description, category_id, cover_urls, rack, shelf)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
	`

	row = tx.QueryRowContext(ctx, query,
		book.Title, book.Authors, book.Description, categoryId, book.CoverUrls, book.Rack, book.Shelf)

	var bookID int
	if err = row.Scan(&bookID); err != nil {
		tx.Rollback()
		return 0, err
	}

	bookSearch := convertors.EntityBookSearchFromEntityBookCreate(bookID, book)

	if _, err = b.search.AddDocumentsWithContext(ctx, []entity.BookSearch{bookSearch}); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return bookID, nil
}

func (b bookRepo) GetBookById(ctx context.Context, id int) (entity.Book, error) {
	query := `SELECT b.id, b.title, b.description, c.name, b.authors, b.cover_urls, b.rack, b.shelf
	FROM books as b
	JOIN categories as c ON b.category_id = c.id
	WHERE b.id = $1`

	row := b.db.QueryRowContext(ctx, query, id)

	var book entity.Book
	err := row.Scan(&book.ID, &book.Title, &book.Description, &book.Category, &book.Authors, &book.CoverUrls, &book.Rack, &book.Shelf)
	if err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (b bookRepo) GetBookByIsbn(ctx context.Context, id int) (entity.Book, error) {
	return entity.Book{}, nil
}

func (b bookRepo) GetBooks(ctx context.Context, page, limit int, filters ...interface{}) ([]entity.Book, error) {
	offset := page * limit

	query := `--sql
	SELECT b.id, b.title, b.description, c.name, b.authors, b.cover_urls, b.rack, b.shelf
	FROM books as b
	JOIN categories as c ON b.category_id = c.id
	LIMIT $1
	OFFSET $2`

	res, err := b.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return []entity.Book{}, err
	}

	var (
		book  entity.Book
		books []entity.Book
	)

	for res.Next() {
		if err = res.Scan(&book.ID, &book.Title, &book.Description, &book.Category, &book.Authors,
			&book.CoverUrls, &book.Rack, &book.Shelf); err != nil {
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

func (b bookRepo) GetBooksByRack(ctx context.Context, rack int) ([]entity.Book, error) {
	res, err := b.db.QueryContext(ctx, `SELECT * FROM books WHERE rack=$1`, rack)
	if err != nil {
		return []entity.Book{}, err
	}

	var books []entity.Book
	var book entity.Book
	for res.Next() {
		err = res.Scan(&book.ID, &book.Title, &book.Category, &book.Description, &book.Authors,
			&book.CoverUrls, &book.Rack, &book.Shelf)
		if err != nil {
			return []entity.Book{}, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (b bookRepo) GetBooksByTextSearch(ctx context.Context, text string) ([]entity.BookSearch, error) {
	searchResp, err := b.search.Search(
		text,
		&meilisearch.SearchRequest{
			AttributesToSearchOn: []string{"title", "authors", "description"},
		},
	)
	if err != nil {
		return []entity.BookSearch{}, err
	}
	return convertors.BooksFromMeiliDocuments(searchResp.Hits), nil
}

func (b bookRepo) UpdateBookInfo(ctx context.Context, id int, fields map[string]interface{}) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var (
		vars []string
		args []interface{}
		num  = 1
	)

	queryBase := "UPDATE books SET"

	for key, value := range fields {
		vars = append(vars, fmt.Sprintf("%s=$%d", key, num))
		args = append(args, value)
		num++
	}

	values := strings.Join(vars, ", ")

	query := strings.Join([]string{queryBase, values, fmt.Sprintf("WHERE id=$%d", num)}, " ")

	args = append(args, id)

	_, execErr := tx.ExecContext(ctx, query, args...)
	if execErr != nil {
		return execErr
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (b bookRepo) UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE books SET rack=$1, shelf=$2 WHERE id=$3`, rack, shelf, id)
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

	res, execErr := tx.ExecContext(ctx, `DELETE FROM books WHERE id=$1`, id)
	if execErr != nil {
		tx.Rollback()
		return execErr
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		tx.Rollback()
		return errs.ErrNoRowsAffected
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
