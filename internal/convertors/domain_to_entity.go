package convertors

import (
	"database/sql"
	domain "eliborate/internal/models/domain"
	"eliborate/internal/models/entity"

	"github.com/lib/pq"
)

func DomainBookInfoToEntity(book domain.BookInfo) entity.BookInfo {
	desc := sql.NullString{}
	if book.Description != nil {
		desc.String = *book.Description
		desc.Valid = true
	}
	return entity.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: desc,
		Category:    book.Category,
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
		BookPlacement: DomainBookPlacementToEntity(book.BookPlacement),
	}
}

func DomainBookToEntity(book domain.Book) entity.Book {
	return entity.Book{
		ID:            book.ID,
		BookInfo:      DomainBookInfoToEntity(book.BookInfo),
		BookPlacement: DomainBookPlacementToEntity(book.BookPlacement),
	}
}

func DomainAdminUserInfoToEntity(user domain.AdminUserInfo) entity.AdminUserInfo {
	return entity.AdminUserInfo{
		Login: user.Login,
	}
}

func DomainUserInfoToEntity(user domain.UserInfo) entity.UserInfo {
	return entity.UserInfo{
		Login: user.Login,
		Name:  user.Name,
	}
}

func DomainAdminUserCreateToEntity(user domain.AdminUserCreate) entity.AdminUserCreate {
	return entity.AdminUserCreate{
		AdminUserInfo: DomainAdminUserInfoToEntity(user.AdminUserInfo),
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
		ID:              user.ID,
		AdminUserCreate: DomainAdminUserCreateToEntity(user.AdminUserCreate),
	}
}

func DomainBookSearchToEntity(book domain.BookSearch) entity.BookSearch {
	return entity.BookSearch{
		ID:          book.ID,
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		Category:    book.Category,
	}
}

func UpdateBookInfoToMap(book domain.UpdateBookInfo) map[string]any {
	values := make(map[string]any, 1)

	if len(book.Authors) != 0 {
		var authors pq.StringArray
		authors.Scan(book.Authors)
		values["authors"] = authors
	}

	if book.Category != nil {
		values["category"] = *book.Category
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
