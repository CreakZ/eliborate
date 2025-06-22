package convertors

import (
	"database/sql"
	domain "eliborate/internal/models/domain"
	"eliborate/internal/models/entity"

	"github.com/lib/pq"
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

func UpdateBookInfoToMap(book domain.UpdateBookInfo) map[string]any {
	values := make(map[string]any, 1)

	if len(book.Authors) != 0 {
		var authors pq.StringArray
		authors.Scan(book.Authors)
		values["authors"] = authors
	}

	if book.CategoryID != 0 {
		values["category_id"] = book.CategoryID
	}

	if book.Description != nil {
		values["description"] = *book.Description
	}

	if book.Title != nil {
		values["title"] = *book.Title
	}

	if len(book.CoverUrls) != 0 {
		var coverUrls pq.StringArray
		coverUrls.Scan(book.CoverUrls)
		values["cover_urls"] = coverUrls
	}

	return values
}
