package repository

import (
	"context"
	"eliborate/internal/models/entity"

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

func (p publicRepo) GetUserByLogin(ctx context.Context, login string) (entity.User, error) {
	var user entity.User

	row := p.db.QueryRowContext(ctx, `SELECT * FROM users WHERE login = $1`, login)
	if err := row.Scan(&user.ID, &user.Name, &user.Login, &user.Password); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (p publicRepo) GetAdminUserByLogin(ctx context.Context, login string) (entity.AdminUser, error) {
	var adminUser entity.AdminUser

	row := p.db.QueryRowContext(ctx, `SELECT * FROM admins WHERE login = $1`, login)
	if err := row.Scan(&adminUser.ID, &adminUser.Login, &adminUser.Password); err != nil {
		return entity.AdminUser{}, err
	}

	return adminUser, nil
}
