package repository

import (
	"context"
	"eliborate/internal/models/entity"
)

type BookRepo interface {
	CreateBook(ctx context.Context, book entity.BookCreate) (int, error)

	GetBookById(ctx context.Context, id int) (entity.Book, error)
	GetBooks(ctx context.Context, offset, limit int, rack *int, searchQuery *string) ([]entity.Book, error)
	GetBooksTotalCount(ctx context.Context) (int, error)

	UpdateBookInfo(ctx context.Context, id int, updates entity.UpdateBookInfo) error
	UpdateBookPlacement(ctx context.Context, id int, updates entity.UpdateBookPlacement) error

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
	GetAll(ctx context.Context) ([]entity.Category, error)
	Update(ctx context.Context, id int, newName string) error
	Delete(ctx context.Context, id int) error
}
