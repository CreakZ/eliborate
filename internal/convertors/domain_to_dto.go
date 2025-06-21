package convertors

import (
	domain "eliborate/internal/models/domain"
	dto "eliborate/internal/models/dto"
)

func DomainCredentialsToDto(credentials domain.Credentials) dto.Credentials {
	return dto.Credentials{
		Login:    credentials.Login,
		Password: credentials.Password,
	}
}

func DomainBookInfoToDto(book domain.BookInfo) dto.BookInfo {
	return dto.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		CoverUrls:   book.CoverUrls,
	}
}

func DomainBookPlacementToDto(book domain.BookPlacement) dto.BookPlacement {
	return dto.BookPlacement{
		Rack:  book.Rack,
		Shelf: book.Shelf,
	}
}

func DomainBookCreateToDto(book domain.BookCreate) dto.BookCreate {
	return dto.BookCreate{
		BookInfo:      DomainBookInfoToDto(book.BookInfo),
		BookPlacement: DomainBookPlacementToDto(book.BookPlacement),
	}
}

func DomainBookToDto(book domain.Book) dto.Book {
	return dto.Book{
		ID:            book.ID,
		BookInfo:      DomainBookInfoToDto(book.BookInfo),
		Category:      book.Category,
		BookPlacement: DomainBookPlacementToDto(book.BookPlacement),
	}
}

func DomainBooksToDto(books []domain.Book) []dto.Book {
	booksDto := make([]dto.Book, 0, len(books))
	for _, book := range books {
		booksDto = append(booksDto, DomainBookToDto(book))
	}
	return booksDto
}

func DomainUserInfoToDto(user domain.UserInfo) dto.UserInfo {
	return dto.UserInfo{
		Login: user.Login,
		Name:  user.Name,
	}
}

func DomainUserCreateToDto(user domain.UserCreate) dto.UserCreate {
	return dto.UserCreate{
		UserInfo: DomainUserInfoToDto(user.UserInfo),
		Password: user.Password,
	}
}

func DomainUserToDto(user domain.User) dto.User {
	return dto.User{
		ID:         user.ID,
		UserCreate: DomainUserCreateToDto(user.UserCreate),
	}
}

func DomainAdminUserToDto(user domain.AdminUser) dto.AdminUser {
	return dto.AdminUser{
		ID:          user.ID,
		Credentials: DomainCredentialsToDto(user.Credentials),
	}
}

func DomainCategoryToDto(category domain.Category) dto.Category {
	return dto.Category{
		ID:   category.ID,
		Name: category.Name,
	}
}

func DomainCategoriesToDto(categories []domain.Category) []dto.Category {
	c := make([]dto.Category, 0, len(categories))

	for _, category := range categories {
		c = append(c, DomainCategoryToDto(category))
	}

	return c
}
