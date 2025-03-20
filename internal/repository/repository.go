package repository

import (
	"context"
	"eliborate/internal/models/entity"
)

type BookRepo interface {
	CreateBook(ctx context.Context, book entity.BookCreate) (int, error)

	GetBookById(ctx context.Context, id int) (entity.Book, error)
	GetBookByIsbn(ctx context.Context, id int) (entity.Book, error)
	GetBooks(ctx context.Context, page, limit int, filters ...interface{}) ([]entity.Book, error)
	GetBooksTotalCount(ctx context.Context) (int, error)
	GetBooksByRack(ctx context.Context, rack int) ([]entity.Book, error)
	GetBooksByTextSearch(ctx context.Context, text string) ([]entity.BookSearch, error)

	UpdateBookInfo(ctx context.Context, id int, fields map[string]interface{}) error
	UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error

	DeleteBook(ctx context.Context, id int) error
}

type PublicRepo interface {
	GetUserByLogin(ctx context.Context, login string) (entity.User, error)
	GetAdminUserByLogin(ctx context.Context, login string) (entity.AdminUser, error)
}

type UserRepo interface {
	Create(ctx context.Context, user entity.UserCreate) (int, error)

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
