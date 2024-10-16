package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type publicRepo struct {
	db *sqlx.DB
}

func InitPublicRepo(db *sqlx.DB) PublicRepo {
	return publicRepo{
		db: db,
	}
}

func (p publicRepo) GetByLogin(ctx context.Context, userType, login string) (int, string, error) {
	query := "SELECT (id, password) FROM %s WHERE login=$1"

	row := p.db.QueryRowContext(ctx, fmt.Sprintf(query, userType), login)

	var (
		id       int
		password string
	)

	if err := row.Scan(&id, &password); err != nil {
		return 0, "", err
	}

	return id, password, nil
}
