package service

import (
	"context"
	"eliborate/internal/models/dto"
)

type BookService interface {
	CreateBook(ctx context.Context, book dto.BookCreate) (int, error)

	GetBookById(ctx context.Context, id int) (dto.Book, error)
	GetBookByIsbn(ctx context.Context, isbn string) (dto.Book, error)
	GetBooks(ctx context.Context, page, limit int, filters ...interface{}) ([]dto.Book, error)
	GetBooksTotalCount(ctx context.Context) (int, error)
	GetBooksByRack(ctx context.Context, rack int) ([]dto.Book, error)
	GetBooksByTextSearch(ctx context.Context, text string) ([]dto.BookSearch, error)

	UpdateBookInfo(ctx context.Context, id int, book dto.UpdateBookInfo) error
	UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error

	DeleteBook(ctx context.Context, id int) error
}

type PublicService interface {
	GetUserByLogin(ctx context.Context, login string) (dto.User, error)
	GetAdminUserByLogin(ctx context.Context, login string) (dto.AdminUser, error)
}

type UserService interface {
	Create(ctx context.Context, user dto.UserCreate) (int, error)

	UpdatePassword(ctx context.Context, id int, password string) error

	Delete(ctx context.Context, id int) error
}

type AdminUserService interface {
	UpdatePassword(ctx context.Context, id int, password string) error
}

type CategoryService interface {
	Create(ctx context.Context, categoryName string) error

	GetAll(ctx context.Context) ([]string, error)

	Update(ctx context.Context, id int, newName string) error

	Delete(ctx context.Context, name string) error
}
