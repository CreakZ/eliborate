package repository

import (
	"context"
	domain "yurii-lib/internal/models/domain"
	"yurii-lib/pkg/errs"

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
	var count int
	rows, err := u.db.QueryContext(ctx, `SELECT COUNT(*) FROM users WHERE login = $1`, user.Login)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if err = rows.Scan(&count); err != nil {
		return 0, err
	}

	if count != 0 {
		return 0, errs.ErrUserAlreadyExists
	}

	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, nil
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

func (u userRepo) CheckByLogin(ctx context.Context, login string) (bool, error) {
	var exists bool

	err := u.db.GetContext(ctx, &exists, `--sql
	SELECT EXISTS (
		SELECT 1 FROM users WHERE login = $1
	);`, login)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u userRepo) GetPassword(ctx context.Context, id int) (string, error) {
	row := u.db.QueryRowContext(ctx, `SELECT password FROM users WHERE id=$1`, id)

	var password string
	if err := row.Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}

func (u userRepo) UpdatePassword(ctx context.Context, id int, password string) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, `UPDATE users SET password=$1 WHERE id=$2`, password, id)
	if err != nil {
		tx.Rollback()

		return err
	}

	affected, _ := res.RowsAffected()

	if affected != 1 {
		tx.Rollback()

		return errs.ErrNoRowsAffected
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

	res, err := tx.ExecContext(ctx, `DELETE FROM admin_users WHERE id=$1`, id)
	if err != nil {
		tx.Rollback()

		return err
	}

	affected, _ := res.RowsAffected()

	if affected != 1 {
		tx.Rollback()

		return errs.ErrNoRowsAffected
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()

		return err
	}

	return nil
}
