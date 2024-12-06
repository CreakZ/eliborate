package repository

import (
	"context"
	domain "eliborate/internal/models/domain"
)

type BookRepo interface {
	CreateBook(ctx context.Context, book domain.BookCreate) (int, error)

	GetBookById(ctx context.Context, id int) (domain.Book, error)
	GetBookByIsbn(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context, page, limit int, filters ...interface{}) ([]domain.Book, error)
	GetBooksTotalCount(ctx context.Context) (int, error)
	GetBooksByRack(ctx context.Context, rack int) ([]domain.Book, error)
	GetBooksByTextSearch(ctx context.Context, text string) ([]domain.BookSearch, error)

	UpdateBookInfo(ctx context.Context, id int, fields map[string]interface{}) error
	UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error

	DeleteBook(ctx context.Context, id int) error
}

type PublicRepo interface {
	GetUserByLogin(ctx context.Context, login string) (domain.User, error)
	GetAdminUserByLogin(ctx context.Context, login string) (domain.AdminUser, error)
}

type UserRepo interface {
	Create(ctx context.Context, user domain.UserCreate) (int, error)

	UpdatePassword(ctx context.Context, id int, password string) error

	Delete(ctx context.Context, id int) error
}

type AdminUserRepo interface {
	UpdatePassword(ctx context.Context, id int, password string) error
}

type CategoryRepo interface {
	Create(ctx context.Context, categoryName string) error

	GetAll(ctx context.Context) ([]string, error)

	Update(ctx context.Context, id int, newName string) error

	Delete(ctx context.Context, name string) error
}
