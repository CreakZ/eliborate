package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type adminUserRepo struct {
	db *sqlx.DB
}

func InitAdminUserRepo(db *sqlx.DB) AdminUserRepo {
	return adminUserRepo{
		db: db,
	}
}

func (u adminUserRepo) UpdatePassword(ctx context.Context, id int, password string) error {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `UPDATE admins SET password=$1 WHERE id=$2`, password, id)
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
