package repository

import (
	"context"
	"yurii-lib/pkg/errs"

	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func InitCategoryRepo(db *sqlx.DB) CategoryRepo {
	return categoryRepo{
		db: db,
	}
}

func (c categoryRepo) CreateCategory(ctx context.Context, categoryName string) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, `INSERT INTO categories name VALUES $1`)
	if err != nil {
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		return errs.ErrNoRowsAffected
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()

		return err
	}

	return nil
}

func (c categoryRepo) GetCategoryNameIfExists(ctx context.Context, name string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS (
		SELECT 1 FROM categories WHERE name = $1
	)`

	if err := c.db.GetContext(ctx, &exists, query, name); err != nil {
		return exists, err
	}

	return exists, nil
}

func (c categoryRepo) GetAllCategories(ctx context.Context) ([]string, error) {
	rows, err := c.db.QueryContext(ctx, `SELECT name from categories`, nil)
	if err != nil {
		return []string{}, err
	}

	var (
		categories []string
		category   string
	)

	for rows.Next() {
		if err = rows.Scan(&category); err != nil {
			return []string{}, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (c categoryRepo) DeleteCategory(ctx context.Context, name string) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, `DELETE FROM categories WHERE name = $1`, name)
	if err != nil {
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		return errs.ErrNoRowsAffected
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()

		return err
	}

	return nil
}
