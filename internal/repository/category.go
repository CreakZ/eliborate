package repository

import (
	"context"
	"eliborate/internal/errs"

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

func (c categoryRepo) Create(ctx context.Context, categoryName string) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err = tx.ExecContext(ctx, `INSERT INTO categories name VALUES $1`); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (c categoryRepo) GetAll(ctx context.Context) ([]string, error) {
	rows, err := c.db.QueryContext(ctx, `SELECT name from categories`)
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

func (c categoryRepo) Update(ctx context.Context, id int, newName string) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, `UPDATE categories SET name = $1 WHERE id = $2`, newName, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		tx.Rollback()
		return errs.ErrNoRowsAffected
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (c categoryRepo) Delete(ctx context.Context, name string) error {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `DELETE FROM categories WHERE name = $1`, name); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
