package convertors

import (
	"database/sql"
	"fmt"
	"strings"
	domain "yurii-lib/internal/models/domain"
	dto "yurii-lib/internal/models/dto"

	"github.com/lib/pq"
)

var categories = []string{
	"Не определено",
	"Библиотека всемирной литературы",
	"Собрание сочинений",
	"Художественная отечественная литература",
	"Художественная зарубежная литература",
	"Поэзия отечественная",
	"Поэзия зарубежная",
	"Детская отечественная литература",
	"Детская зарубежная литература",
	"Моя первая книга",
	"Книги нашего детства",
	"Моё первое собрание сочинений",
	"Сказки отечественные",
	"Сказки зарубежные",
	"Большая библиотека приключений и научной фантастики",
	"Библиотека приключений и фантастики",
	"Классическая библиотека приключений и фантастики",
	"Ретро библиотека приключений и фантастики",
	"Приключения. Отечественная литература",
	"Приключения. Зарубежная литература",
	"Фантастика. Отечественная литература",
	"Фантастика. Зарубежная литература",
	"История России. Отечественные авторы",
	"История России. Зарубежные авторы",
	"Летописи",
	"Зарубежная история",
	"Нумизматика",
	"Альтернативная история",
	"Библиотека отечественной общественной мысли",
	"Духовная литература",
	"Справочная литература",
	"Энциклопедический словарь Терра",
	"Энциклопедический словарь Брокгауза и Эфрона",
	"Энциклопедии",
	"Великие путешествия",
	"Великие полководцы",
	"Учебная литература",
	"Настольные игры",
	"Словарь русского языка",
	"Русский язык",
	"Филология",
	"География",
	"Архитектура",
	"Искусство",
	"Живопись",
	"Сборные издания",
	"Строительство",
	"Домашнее хозяйство",
	"Сад, огород",
	"Кухня",
	"Медицина",
}

func CategoryToInt(category string) int {
	for i := range categories {
		if strings.Compare(categories[i], category) == 0 {
			return i
		}
	}

	// -1 возвращается в том случае, если приведенная в запросе категория не соотвествует ни одной из перечисленных
	return -1
}

func CategoryToString(category int) string {
	return categories[category]
}

func ToDomainBookInfo(book dto.BookInfo) domain.BookInfo {
	var cover sql.NullString
	if book.CoverURL != nil {
		cover.String = *book.CoverURL
		cover.Valid = true
	}

	var desc sql.NullString
	if book.Description != nil {
		desc.String = *book.Description
		desc.Valid = true
	}

	return domain.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: desc,
		Category:    CategoryToInt(book.Category),
		IsForeign:   book.IsForeign,
		CoverURL:    cover,
	}
}

func UpdateBookInfoToMap(book dto.UpdateBookInfo) map[string]interface{} {
	values := make(map[string]interface{}, 1)

	if book.Authors != nil {
		var authors pq.StringArray

		authors.Scan(*book.Authors)

		values["authors"] = authors
	}

	if book.Category != nil {
		values["category"] = *book.Category
	}

	if book.Description != nil {
		values["description"] = *book.Description
	}

	if book.IsForeign != nil {
		values["is_foreign"] = *book.IsForeign
	}

	if book.Title != nil {
		values["title"] = *book.Title
	}

	if book.CoverURL != nil {
		values["logo"] = *book.CoverURL
	}

	fmt.Println(len(values))

	return values
}

func ToDomainBookPlacement(book dto.BookPlacement) domain.BookPlacement {
	return domain.BookPlacement{
		BookInfo: ToDomainBookInfo(book.BookInfo),
		Rack:     book.Rack,
		Shelf:    book.Shelf,
	}
}

func ToDomainBook(book dto.Book) domain.Book {
	return domain.Book{
		ID:            book.ID,
		BookPlacement: ToDomainBookPlacement(book.BookPlacement),
	}
}

func ToDtoBookInfo(book domain.BookInfo) dto.BookInfo {
	var cover *string
	if book.CoverURL.Valid {
		cover = &book.CoverURL.String
	}

	var desc *string
	if book.Description.Valid {
		desc = &book.Description.String
	}

	return dto.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: desc,
		Category:    CategoryToString(book.Category),
		IsForeign:   book.IsForeign,
		CoverURL:    cover,
	}
}

func ToDtoBookPlacement(book domain.BookPlacement) dto.BookPlacement {
	return dto.BookPlacement{
		BookInfo: ToDtoBookInfo(book.BookInfo),
		Rack:     book.Rack,
		Shelf:    book.Shelf,
	}
}

func ToDtoBook(book domain.Book) dto.Book {
	return dto.Book{
		ID:            book.ID,
		BookPlacement: ToDtoBookPlacement(book.BookPlacement),
	}
}

func ToDomainAdminUserInfo(user dto.AdminUserInfo) domain.AdminUserInfo {
	return domain.AdminUserInfo{
		Login: user.Login,
	}
}

func ToDomainUserInfo(user dto.UserInfo) domain.UserInfo {
	return domain.UserInfo{
		Login: user.Login,
		Name:  user.Name,
	}
}

func ToDtoAdminUserInfo(user domain.AdminUserInfo) dto.AdminUserInfo {
	return dto.AdminUserInfo{
		Login: user.Login,
	}
}

func ToDtoUserInfo(user domain.UserInfo) dto.UserInfo {
	return dto.UserInfo{
		Login: user.Login,
		Name:  user.Name,
	}
}

func ToDomainAdminUserCreate(user dto.AdminUserCreate) domain.AdminUserCreate {
	return domain.AdminUserCreate{
		AdminUserInfo: ToDomainAdminUserInfo(user.AdminUserInfo),
	}
}

func ToDomainUserCreate(user dto.UserCreate) domain.UserCreate {
	return domain.UserCreate{
		UserInfo: ToDomainUserInfo(user.UserInfo),
		Password: user.Password,
	}
}

func ToDtoAdminUserCreate(user domain.AdminUserCreate) dto.AdminUserCreate {
	return dto.AdminUserCreate{
		AdminUserInfo: ToDtoAdminUserInfo(user.AdminUserInfo),
		Password:      user.Password,
	}
}

func ToDtoUserCreate(user domain.UserCreate) dto.UserCreate {
	return dto.UserCreate{
		UserInfo: ToDtoUserInfo(user.UserInfo),
		Password: user.Password,
	}
}
