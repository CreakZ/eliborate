package repository

import (
	"context"
	domain "eliborate/internal/models/domain"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func InitUserRepo(db *sqlx.DB) UserRepo {
	return userRepo{
		db: db,
	}
}

func (u userRepo) Create(ctx context.Context, user domain.UserCreate) (int, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users (login, name, password) VALUES ($1, $2, $3) RETURNING id;`

	row := tx.QueryRowContext(ctx, query, user.Login, user.Name, user.Password)

	var id int
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (u userRepo) UpdatePassword(ctx context.Context, id int, password string) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `UPDATE clients SET password=$1 WHERE id=$2`, password, id); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()

		return err
	}

	return nil
}

func (u userRepo) Delete(ctx context.Context, id int) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.ExecContext(ctx, `DELETE FROM clients WHERE id=$1`, id); err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
