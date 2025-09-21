package repository

import (
	"context"
	"eliborate/internal/errs"
	"eliborate/internal/models/entity"

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
	_, err := c.db.ExecContext(ctx, `INSERT INTO categories name VALUES $1`)
	return err
}

func (c categoryRepo) GetAll(ctx context.Context) ([]entity.Category, error) {
	rows, err := c.db.QueryContext(ctx, `SELECT id, name from categories`)
	if err != nil {
		return []entity.Category{}, err
	}

	var categories []entity.Category

	for rows.Next() {
		var category entity.Category
		if err = rows.Scan(&category.ID, &category.Name); err != nil {
			return []entity.Category{}, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (c categoryRepo) Update(ctx context.Context, id int, newName string) error {
	res, err := c.db.ExecContext(ctx, `UPDATE categories SET name = $1 WHERE id = $2`, newName, id)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return errs.ErrEntityNotFound
	}
	return nil
}

func (c categoryRepo) Delete(ctx context.Context, id int) error {
	res, err := c.db.ExecContext(ctx, `DELETE FROM categories WHERE id = $1`, id)
	if err != nil {
		return err
	}
	if affected, _ := res.RowsAffected(); affected == 0 {
		return errs.ErrEntityNotFound
	}
	return nil
}
