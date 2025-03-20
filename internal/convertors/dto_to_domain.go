package convertors

import (
	domain "eliborate/internal/models/domain"
	dto "eliborate/internal/models/dto"
)

func DtoBookInfoToDomain(book dto.BookInfo) domain.BookInfo {
	return domain.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		Category:    book.Category,
		CoverUrls:   book.CoverUrls,
	}
}

func DtoBookPlacementToDomain(book dto.BookPlacement) domain.BookPlacement {
	return domain.BookPlacement{
		Rack:  book.Rack,
		Shelf: book.Shelf,
	}
}

func DtoBookCreateToDomain(book dto.BookCreate) domain.BookCreate {
	return domain.BookCreate{
		BookInfo:      DtoBookInfoToDomain(book.BookInfo),
		BookPlacement: DtoBookPlacementToDomain(book.BookPlacement),
	}
}

func DtoBookToDomain(book dto.Book) domain.Book {
	return domain.Book{
		ID:            book.ID,
		BookInfo:      DtoBookInfoToDomain(book.BookInfo),
		BookPlacement: DtoBookPlacementToDomain(book.BookPlacement),
	}
}

func DtoAdminUserInfoToDomain(user dto.AdminUserInfo) domain.AdminUserInfo {
	return domain.AdminUserInfo{
		Login: user.Login,
	}
}

func DtoUserInfoToDomain(user dto.UserInfo) domain.UserInfo {
	return domain.UserInfo{
		Login: user.Login,
		Name:  user.Name,
	}
}

func DtoAdminUserCreateToDomain(user dto.AdminUserCreate) domain.AdminUserCreate {
	return domain.AdminUserCreate{
		AdminUserInfo: DtoAdminUserInfoToDomain(user.AdminUserInfo),
	}
}

func DtoUserCreateToDomain(user dto.UserCreate) domain.UserCreate {
	return domain.UserCreate{
		UserInfo: DtoUserInfoToDomain(user.UserInfo),
		Password: user.Password,
	}
}

func DtoUserToDomain(user dto.User) domain.User {
	return domain.User{
		ID:         user.ID,
		UserCreate: DtoUserCreateToDomain(user.UserCreate),
	}
}

func DtoAdminUserToDomain(user dto.AdminUser) domain.AdminUser {
	return domain.AdminUser{
		ID:              user.ID,
		AdminUserCreate: DtoAdminUserCreateToDomain(user.AdminUserCreate),
	}
}

func DtoBookSearchToDomain(book dto.BookSearch) domain.BookSearch {
	return domain.BookSearch{
		ID:          book.ID,
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		Category:    book.Category,
	}
}

func DtoUpdateBookInfoToDomain(updateBook dto.UpdateBookInfo) domain.UpdateBookInfo {
	return domain.UpdateBookInfo{
		Title:       updateBook.Title,
		Authors:     updateBook.Authors,
		Description: updateBook.Description,
		Category:    updateBook.Category,
		CoverUrls:   updateBook.CoverUrls,
	}
}
