package service

import (
	"context"
	"yurii-lib/internal/models/dto"
)

type BookService interface {
	CreateBook(ctx context.Context, book dto.BookPlacement) (int, error)

	GetBooks(ctx context.Context, page, limit int) ([]dto.Book, error)
	GetBooksByRack(ctx context.Context, rack int) ([]dto.Book, error)
	GetBooksByTextSearch(ctx context.Context, text string) ([]dto.Book, error)

	UpdateBookInfo(ctx context.Context, id int, book dto.UpdateBookInfo) error
	UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error

	DeleteBook(ctx context.Context, id int) error
}

type UserService interface {
	CreateUser(ctx context.Context, user dto.UserCreate) (int, error)

	GetUserPassword(ctx context.Context, id int) (string, error)

	UpdateUserPassword(ctx context.Context, id int, password string) error

	DeleteUser(ctx context.Context, id int) error
}

type AdminUserService interface {
	CreateAdminUser(ctx context.Context, user dto.AdminUserCreate) (int, error)

	GetAdminUserPassword(ctx context.Context, id int) (string, error)

	UpdateAdminUserPassword(ctx context.Context, id int, password string) error

	DeleteAdminUser(ctx context.Context, id int) error
}
