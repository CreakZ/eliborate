package convertors

import (
	domain "eliborate/internal/models/domain"
	"eliborate/internal/models/entity"
)

func EntityCredentialsToDomain(credentials entity.Credentials) domain.Credentials {
	return domain.Credentials{
		Login:    credentials.Login,
		Password: credentials.Password,
	}
}

func EntityBookInfoToDomain(book entity.BookInfo) domain.BookInfo {
	return domain.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		CoverUrls:   book.CoverUrls,
	}
}

func EntityBookPlacementToDomain(book entity.BookPlacement) domain.BookPlacement {
	return domain.BookPlacement{
		Rack:  book.Rack,
		Shelf: book.Shelf,
	}
}

func EntityBookCreateToDomain(book entity.BookCreate) domain.BookCreate {
	return domain.BookCreate{
		BookInfo:      EntityBookInfoToDomain(book.BookInfo),
		BookPlacement: EntityBookPlacementToDomain(book.BookPlacement),
	}
}

func EntityBookToDomain(book entity.Book) domain.Book {
	var category string
	if book.Category.Valid {
		category = book.Category.String
	}
	return domain.Book{
		ID:            book.ID,
		BookInfo:      EntityBookInfoToDomain(book.BookInfo),
		Category:      category,
		BookPlacement: EntityBookPlacementToDomain(book.BookPlacement),
	}
}

func EntityUserInfoToDomain(user entity.UserInfo) domain.UserInfo {
	return domain.UserInfo{
		Login: user.Login,
		Name:  user.Name,
	}
}

func EntityUserCreateToDomain(user entity.UserCreate) domain.UserCreate {
	return domain.UserCreate{
		UserInfo: EntityUserInfoToDomain(user.UserInfo),
		Password: user.Password,
	}
}

func EntityUserToDomain(user entity.User) domain.User {
	return domain.User{
		ID:         user.ID,
		UserCreate: EntityUserCreateToDomain(user.UserCreate),
	}
}

func EntityAdminUserToDomain(user entity.AdminUser) domain.AdminUser {
	return domain.AdminUser{
		ID:          user.ID,
		Credentials: EntityCredentialsToDomain(user.Credentials),
	}
}

func EntityBookSearchToDomain(book entity.BookSearch) domain.BookSearch {
	return domain.BookSearch{
		ID:          book.ID,
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		Category:    book.Category,
	}
}

func EntityCategoryToDomain(category entity.Category) domain.Category {
	return domain.Category{
		ID:   category.ID,
		Name: category.Name,
	}
}

func EntityCategoriesToDomain(categories []entity.Category) []domain.Category {
	c := make([]domain.Category, 0, len(categories))

	for _, category := range categories {
		c = append(c, EntityCategoryToDomain(category))
	}

	return c
}
