package service

import (
	"context"
	"eliborate/internal/models/domain"
)

type BookService interface {
	CreateBook(ctx context.Context, book domain.BookCreate) (int, error)

	GetBookById(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context, page, limit int, rack *int, searchQuery *string) ([]domain.Book, error)
	GetBooksTotalCount(ctx context.Context) (int, error)

	UpdateBookInfo(ctx context.Context, id int, updates domain.UpdateBookInfo) error
	UpdateBookPlacement(ctx context.Context, id int, updates domain.UpdateBookPlacement) error

	DeleteBook(ctx context.Context, id int) error
}

type PublicService interface {
	GetUserByLogin(ctx context.Context, login string) (domain.User, error)
	GetAdminUserByLogin(ctx context.Context, login string) (domain.AdminUser, error)
}

type UserService interface {
	Create(ctx context.Context, user domain.UserCreate) (int, error)

	UpdatePassword(ctx context.Context, id int, password string) error

	Delete(ctx context.Context, id int) error
}

type AdminUserService interface {
	UpdatePassword(ctx context.Context, id int, password string) error
}

type CategoryService interface {
	Create(ctx context.Context, categoryName string) error
	GetAll(ctx context.Context) ([]domain.Category, error)
	Update(ctx context.Context, id int, newName string) error
	Delete(ctx context.Context, id int) error
}
