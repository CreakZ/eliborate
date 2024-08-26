package repository

import (
	"context"
	"fmt"
	"strings"
	domain "yurii-lib/internal/models/domain"
	"yurii-lib/pkg/errs"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
)

type bookRepo struct {
	db  *sqlx.DB
	svc *s3.S3
}

func InitBookRepo(db *sqlx.DB, svc *s3.S3) BookRepo {
	return bookRepo{
		db:  db,
		svc: svc,
	}
}

func (b bookRepo) CreateBook(ctx context.Context, book domain.BookPlacement) (int, error) {
	// Попытка поместить файл в S3-хранилище
	/*
		var imgKey sql.NullString

		if book.Cover.Valid {
			resp, err := http.Get(book.Cover.String)
			if err != nil {
				return 0, err
			}

			// handle error
			body, _ := io.ReadAll(resp.Body)

			img := bytes.NewReader(body)

			imgID := uuid.New().String()

			key := fmt.Sprintf("%s.webp", imgID)

			_, err = b.svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
				Bucket: aws.String(viper.GetString(config.S3ImageBucketName)),
				Key:    aws.String(key),
				Body:   img,
			})
			if err != nil {
				return 0, err
			}

			imgKey.String = key
		}
	*/

	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO books (title, authors, description, category, is_foreign, logo, rack, shelf) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	row := tx.QueryRowContext(ctx, query,
		book.Title, book.Authors, book.Description, book.Category, book.IsForeign, book.CoverURL, book.Rack, book.Shelf)

	var bookID int
	if err = row.Scan(&bookID); err != nil {
		if errTxRb := tx.Rollback(); errTxRb != nil {
			return 0, errs.MergeErrors("create book", []string{err.Error(), errTxRb.Error()})
		}

		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return bookID, nil
}

// TODO
func (b bookRepo) CreateCategory(ctx context.Context, category string) error {
	return nil
}

func (b bookRepo) GetBooks(ctx context.Context, page, limit int) ([]domain.Book, error) {
	offset := page * limit

	res, err := b.db.QueryContext(ctx, `SELECT * FROM books LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return []domain.Book{}, err
	}

	var (
		book  domain.Book
		books []domain.Book
	)

	for res.Next() {
		if err = res.Scan(&book.ID, &book.Title, &book.Description, &book.Category,
			&book.Authors, &book.IsForeign, &book.CoverURL, &book.Rack, &book.Shelf); err != nil {
			return []domain.Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (b bookRepo) GetBooksByRack(ctx context.Context, rack int) ([]domain.Book, error) {
	res, err := b.db.QueryContext(ctx, `SELECT * FROM books WHERE rack=$1`, rack)
	if err != nil {
		return []domain.Book{}, err
	}

	var books []domain.Book
	var book domain.Book
	for res.Next() {
		err = res.Scan(&book.ID, &book.Title, &book.Category, &book.Description, &book.Authors,
			&book.IsForeign, &book.CoverURL, &book.Rack, &book.Shelf)
		if err != nil {
			// scan error
			return []domain.Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

// TODO
func (b bookRepo) GetBooksByTextSearch(ctx context.Context, text string) ([]domain.Book, error) {
	query := `SELECT * FROM books 
	WHERE title % $1 OR title LIKE '%$1%' OR 
	EXISTS (
		SELECT 1 FROM unnest(authors) AS author 
		WHERE author % $1 OR author LIKE '%$1%'
	);`

	rows, err := b.db.QueryContext(ctx, query, text)
	if err != nil {
		return []domain.Book{}, err
	}

	var (
		book  domain.Book
		books []domain.Book
	)

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Category, &book.Authors,
			&book.IsForeign, &book.CoverURL, &book.Rack, &book.Shelf)
		if err != nil {
			return []domain.Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (b bookRepo) UpdateBookInfo(ctx context.Context, id int, fields map[string]interface{}) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		// tx begin error
		return err
	}

	var (
		vars []string
		args []interface{}
		num  = 1
	)

	queryBase := `UPDATE books SET`

	for key, value := range fields {
		vars = append(vars, fmt.Sprintf("%s=$%d", key, num))
		args = append(args, value)
		num += 1
	}

	values := strings.Join(vars, ", ")

	query := strings.Join([]string{queryBase, values, fmt.Sprintf("WHERE id=$%d", num)}, " ")

	args = append(args, id)

	res, execErr := tx.ExecContext(ctx, query, args...)
	if execErr != nil {
		// query error
		return execErr
	}

	if affected, _ := res.RowsAffected(); affected != 1 {
		// not a single row affected error
		return errs.ErrNoRowsAffected
	}

	if err = tx.Commit(); err != nil {
		// tx commit error
		return err
	}

	return nil
}

// Modify this method
func (b bookRepo) UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		// tx begin error
		return err
	}

	res, qErr := tx.ExecContext(ctx, `UPDATE library SET rack=$1, shelf=$2 WHERE book_id=$3`, rack, shelf, id)
	if qErr != nil {
		// query error
		return qErr
	}

	if affected, _ := res.RowsAffected(); affected != 1 {
		// not a single row affected error
		return errs.ErrNoRowsAffected
	}

	if err = tx.Commit(); err != nil {
		// tx commit error
		return err
	}

	return nil
}

func (b bookRepo) DeleteBook(ctx context.Context, id int) error {
	tx, err := b.db.BeginTx(ctx, nil)
	if err != nil {
		// tx begin error
		return err
	}

	res, execErr := tx.ExecContext(ctx, `DELETE FROM books WHERE id=$1`, id)
	if execErr != nil {
		tx.Rollback()

		return execErr
	}

	if affected, _ := res.RowsAffected(); affected != 1 {
		tx.Rollback()

		return errs.ErrNoRowsAffected
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()

		return err
	}

	return nil
}
