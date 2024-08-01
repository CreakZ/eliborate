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

func (u userRepo) CreateAdminUser(ctx context.Context, user domain.AdminUserCreate) (int, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, nil
	}

	row := tx.QueryRowContext(ctx, `INSERT INTO admin_users VALUES ($1, $2) RETURNING id;`, user.Login, user.Password)

	var id int
	if err = row.Scan(&id); err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}

func (u userRepo) CreateUser(ctx context.Context, user domain.UserCreate) (int, error) {
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

func (u userRepo) GetAdminUserPassword(ctx context.Context, id int) (string, error) {
	row := u.db.QueryRowContext(ctx, `SELECT password FROM admin_users WHERE id=$1`, id)

	var password string
	if err := row.Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}

func (u userRepo) GetUserPassword(ctx context.Context, id int) (string, error) {
	row := u.db.QueryRowContext(ctx, `SELECT password FROM users WHERE id=$1`, id)

	var password string
	if err := row.Scan(&password); err != nil {
		return "", err
	}

	return password, nil
}

func (u userRepo) UpdateAdminUserPassword(ctx context.Context, id int, password string) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, `UPDATE admin_users SET password=$1 WHERE id=$2`, password, id)
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

func (u userRepo) UpdateUserPassword(ctx context.Context, id int, password string) error {
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

func (u userRepo) DeleteAdminUser(ctx context.Context, id int) error {
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

func (u userRepo) DeleteUser(ctx context.Context, id int) error {
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
