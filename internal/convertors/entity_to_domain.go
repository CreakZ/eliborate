package convertors

import (
	domain "eliborate/internal/models/domain"
	"eliborate/internal/models/entity"
)

func EntityBookInfoToDomain(book entity.BookInfo) domain.BookInfo {
	var desc *string
	if book.Description.Valid {
		desc = &book.Description.String
	}
	return domain.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: desc,
		Category:    book.Category,
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
	return domain.Book{
		ID:            book.ID,
		BookInfo:      EntityBookInfoToDomain(book.BookInfo),
		BookPlacement: EntityBookPlacementToDomain(book.BookPlacement),
	}
}

func EntityAdminUserInfoToDomain(user entity.AdminUserInfo) domain.AdminUserInfo {
	return domain.AdminUserInfo{
		Login: user.Login,
	}
}

func EntityUserInfoToDomain(user entity.UserInfo) domain.UserInfo {
	return domain.UserInfo{
		Login: user.Login,
		Name:  user.Name,
	}
}

func EntityAdminUserCreateToDomain(user entity.AdminUserCreate) domain.AdminUserCreate {
	return domain.AdminUserCreate{
		AdminUserInfo: EntityAdminUserInfoToDomain(user.AdminUserInfo),
		Password:      user.Password,
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
		ID:              user.ID,
		AdminUserCreate: EntityAdminUserCreateToDomain(user.AdminUserCreate),
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
