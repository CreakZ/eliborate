package convertors

import (
	"database/sql"
	domain "eliborate/internal/models/domain"
	"eliborate/internal/models/entity"
)

func DomainCredentialsToEntity(credentials domain.Credentials) entity.Credentials {
	return entity.Credentials{
		Login:    credentials.Login,
		Password: credentials.Password,
	}
}

func DomainBookInfoToEntity(book domain.BookInfo) entity.BookInfo {
	return entity.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		CoverUrls:   book.CoverUrls,
	}
}

func DomainBookPlacementToEntity(book domain.BookPlacement) entity.BookPlacement {
	return entity.BookPlacement{
		Rack:  book.Rack,
		Shelf: book.Shelf,
	}
}

func DomainBookCreateToEntity(book domain.BookCreate) entity.BookCreate {
	return entity.BookCreate{
		BookInfo:      DomainBookInfoToEntity(book.BookInfo),
		CategoryID:    book.CategoryID,
		BookPlacement: DomainBookPlacementToEntity(book.BookPlacement),
	}
}

func DomainBookToEntity(book domain.Book) entity.Book {
	var category sql.NullString
	if book.Category == "" {
		category.String = book.Category
		category.Valid = true
	}
	return entity.Book{
		ID:            book.ID,
		BookInfo:      DomainBookInfoToEntity(book.BookInfo),
		Category:      category,
		BookPlacement: DomainBookPlacementToEntity(book.BookPlacement),
	}
}

func DomainUserInfoToEntity(user domain.UserInfo) entity.UserInfo {
	return entity.UserInfo{
		Login: user.Login,
		Name:  user.Name,
	}
}

func DomainUserCreateToEntity(user domain.UserCreate) entity.UserCreate {
	return entity.UserCreate{
		UserInfo: DomainUserInfoToEntity(user.UserInfo),
		Password: user.Password,
	}
}

func ToEntityUser(user domain.User) entity.User {
	return entity.User{
		ID:         user.ID,
		UserCreate: DomainUserCreateToEntity(user.UserCreate),
	}
}

func DomainAdminUser(user domain.AdminUser) entity.AdminUser {
	return entity.AdminUser{
		ID:          user.ID,
		Credentials: DomainCredentialsToEntity(user.Credentials),
	}
}

func DomainUpdateBookInfoToEntity(book domain.UpdateBookInfo) entity.UpdateBookInfo {
	return entity.UpdateBookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		CategoryID:  book.CategoryID,
		CoverUrls:   book.CoverUrls,
	}
}

func DomainUpdateBookPlacementToEntity(book domain.UpdateBookPlacement) entity.UpdateBookPlacement {
	return entity.UpdateBookPlacement{
		Rack:  book.Rack,
		Shelf: book.Shelf,
	}
}
