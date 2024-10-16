package service

import (
	"context"
	"yurii-lib/internal/models/dto"
)

type BookService interface {
	CreateBook(ctx context.Context, book dto.BookPlacement) (int, error)

	GetBooks(ctx context.Context, page, limit int) ([]dto.Book, error)
	GetBooksTotalCount(ctx context.Context) (int, error)
	GetBooksByRack(ctx context.Context, rack int) ([]dto.Book, error)
	GetBooksByTextSearch(ctx context.Context, text string) ([]dto.Book, error)

	UpdateBookInfo(ctx context.Context, id int, book dto.UpdateBookInfo) error
	UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error

	DeleteBook(ctx context.Context, id int) error
}

type PublicService interface {
	GetByLogin(ctx context.Context, userType, login string) (int, string, error)
}

type UserService interface {
	Create(ctx context.Context, user dto.UserCreate) (int, error)

	CheckByLogin(ctx context.Context, login string) (bool, error)

	GetPassword(ctx context.Context, id int) (string, error)

	UpdatePassword(ctx context.Context, id int, password string) error

	Delete(ctx context.Context, id int) error
}

type AdminUserService interface {
	// Create(ctx context.Context, user dto.AdminUserCreate) (int, error)

	GetPassword(ctx context.Context, id int) (string, error)

	UpdatePassword(ctx context.Context, id int, password string) error

	// Delete(ctx context.Context, id int) error
}
