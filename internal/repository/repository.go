package repository

import (
	"context"
	domain "yurii-lib/internal/models/domain"
)

type BookRepo interface {
	CreateBook(ctx context.Context, book domain.BookPlacement) (int, error)

	CreateCategory(ctx context.Context, category string) error

	GetBooks(ctx context.Context, page, limit int) ([]domain.Book, error)
	GetBooksTotalCount(ctx context.Context) (int, error)
	GetBooksByRack(ctx context.Context, rack int) ([]domain.Book, error)
	GetBooksByTextSearch(ctx context.Context, text string) ([]domain.Book, error)

	UpdateBookInfo(ctx context.Context, id int, fields map[string]interface{}) error
	UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error

	DeleteBook(ctx context.Context, id int) error
}

type PublicRepo interface {
	GetByLogin(ctx context.Context, userType, login string) (int, string, error)
}

type UserRepo interface {
	Create(ctx context.Context, user domain.UserCreate) (int, error)

	CheckByLogin(ctx context.Context, login string) (bool, error)

	GetPassword(ctx context.Context, id int) (string, error)

	UpdatePassword(ctx context.Context, id int, password string) error

	Delete(ctx context.Context, id int) error
}

type AdminUserRepo interface {
	// Create(ctx context.Context, user domain.AdminUserCreate) (int, error)

	GetPassword(ctx context.Context, id int) (string, error)

	UpdatePassword(ctx context.Context, id int, password string) error

	// Delete(ctx context.Context, id int) error
	// DeleteAll(ctx context.Context) error
}

type CategoryRepo interface {
	CreateCategory(ctx context.Context, categoryName string) error

	GetCategoryNameIfExists(ctx context.Context, name string) (bool, error) // вспомогательный метод для проверки существования
	GetAllCategories(ctx context.Context) ([]string, error)

	DeleteCategory(ctx context.Context, name string) error
}
