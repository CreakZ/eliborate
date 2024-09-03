package repository

import (
	"context"
	domain "yurii-lib/internal/models/domain"
)

type BookRepo interface {
	CreateBook(ctx context.Context, book domain.BookPlacement) (int, error)
	// TODO
	CreateCategory(ctx context.Context, category string) error

	GetBooks(ctx context.Context, page, limit int) ([]domain.Book, error)
	GetBooksByRack(ctx context.Context, rack int) ([]domain.Book, error)
	GetBooksByTextSearch(ctx context.Context, text string) ([]domain.Book, error)

	UpdateBookInfo(ctx context.Context, id int, fields map[string]interface{}) error
	UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error

	DeleteBook(ctx context.Context, id int) error
}

type UserRepo interface {
	CreateUser(ctx context.Context, user domain.UserCreate) (int, error)

	GetUserPassword(ctx context.Context, id int) (string, error)

	UpdateUserPassword(ctx context.Context, id int, password string) error

	DeleteUser(ctx context.Context, id int) error
}

// There cannot be less than 1 admin user
type AdminUserRepo interface {
	CreateAdminUser(ctx context.Context, user domain.AdminUserCreate) (int, error)

	GetAdminUserPassword(ctx context.Context, id int) (string, error)

	UpdateAdminUserPassword(ctx context.Context, id int, password string) error

	DeleteAdminUser(ctx context.Context, id int) error
}
