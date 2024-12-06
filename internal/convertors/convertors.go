package convertors

import (
	"database/sql"
	domain "eliborate/internal/models/domain"
	dto "eliborate/internal/models/dto"
	"reflect"

	"github.com/lib/pq"
)

func ToDomainBookInfo(book dto.BookInfo) domain.BookInfo {
	var coverUrls pq.StringArray
	coverUrls.Scan(book.CoverUrls)

	var desc sql.NullString
	if book.Description != nil {
		desc.String = *book.Description
		desc.Valid = true
	}

	return domain.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: desc,
		Category:    book.Category,
		CoverUrls:   coverUrls,
	}
}

func UpdateBookInfoToMap(book dto.UpdateBookInfo) map[string]interface{} {
	values := make(map[string]interface{}, 1)

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

func ToDomainBookPlacement(book dto.BookPlacement) domain.BookPlacement {
	return domain.BookPlacement{
		Rack:  book.Rack,
		Shelf: book.Shelf,
	}
}

func ToDomainBook(book dto.Book) domain.Book {
	return domain.Book{
		ID:            book.ID,
		BookPlacement: ToDomainBookPlacement(book.BookPlacement),
	}
}

func ToDtoBookInfo(book domain.BookInfo) dto.BookInfo {
	var desc *string
	if book.Description.Valid {
		desc = &book.Description.String
	}

	return dto.BookInfo{
		Title:       book.Title,
		Authors:     book.Authors,
		Description: desc,
		Category:    book.Category,
		CoverUrls:   book.CoverUrls,
	}
}

func ToDtoBookPlacement(book domain.BookPlacement) dto.BookPlacement {
	return dto.BookPlacement{
		Rack:  book.Rack,
		Shelf: book.Shelf,
	}
}

func ToDomainBookCreate(book dto.BookCreate) domain.BookCreate {
	return domain.BookCreate{
		BookInfo:      ToDomainBookInfo(book.BookInfo),
		BookPlacement: ToDomainBookPlacement(book.BookPlacement),
	}
}

func ToDtoBookCreate(book domain.BookCreate) dto.BookCreate {
	return dto.BookCreate{
		BookInfo:      ToDtoBookInfo(book.BookInfo),
		BookPlacement: ToDtoBookPlacement(book.BookPlacement),
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

func ToDtoUser(user domain.User) dto.User {
	return dto.User{
		ID:         user.ID,
		UserCreate: ToDtoUserCreate(user.UserCreate),
	}
}

func ToDomainUser(user dto.User) domain.User {
	return domain.User{
		ID:         user.ID,
		UserCreate: ToDomainUserCreate(user.UserCreate),
	}
}

func ToDtoAdminUser(user domain.AdminUser) dto.AdminUser {
	return dto.AdminUser{
		ID:              user.ID,
		AdminUserCreate: ToDtoAdminUserCreate(user.AdminUserCreate),
	}
}

func ToDomainAdminUser(user dto.AdminUser) domain.AdminUser {
	return domain.AdminUser{
		ID:              user.ID,
		AdminUserCreate: ToDomainAdminUserCreate(user.AdminUserCreate),
	}
}

func ToDomainBookSearch(book dto.BookSearch) domain.BookSearch {
	return domain.BookSearch{
		ID:          book.ID,
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		Category:    book.Category,
	}
}

func ToDtoBookSearch(book domain.BookSearch) dto.BookSearch {
	return dto.BookSearch{
		ID:          book.ID,
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		Category:    book.Category,
	}
}

func DomainBookSearchFromBookCreate(bookId int, book domain.BookCreate) domain.BookSearch {
	var bookDesc string
	if book.Description.Valid {
		bookDesc = book.Description.String
	}
	return domain.BookSearch{
		ID:          bookId,
		Title:       book.Title,
		Authors:     book.Authors,
		Description: bookDesc,
		Category:    book.Category,
	}
}

func MeiliDocumentFromBookSearch(primaryKey int, bookSearch domain.BookSearch) map[string]interface{} {
	switch reflect.TypeOf(bookSearch).Kind() {
	case reflect.Struct:
		fields := reflect.VisibleFields(reflect.TypeOf(bookSearch))
		values := reflect.ValueOf(bookSearch)

		doc := make(map[string]interface{})
		for i, field := range fields {
			if !values.Field(i).CanInterface() {
				continue
			}
			tag, ok := field.Tag.Lookup("search")
			if ok {
				doc[tag] = values.Field(i).Interface()
			}
		}
		return doc

	default:
		return map[string]interface{}{}
	}
}

// doc is expected to be map[string]interface{}
func BookFromMeiliDocument(doc interface{}) domain.BookSearch {
	docMap, ok := doc.(map[string]interface{})
	if !ok {
		return domain.BookSearch{}
	}

	var book domain.BookSearch
	bookValue := reflect.ValueOf(&book).Elem()
	bookType := bookValue.Type()

	for i := range bookType.NumField() {
		field := bookType.Field(i)
		searchTag := field.Tag.Get("search")

		if value, ok := docMap[searchTag]; ok {
			fieldValue := bookValue.Field(i)

			if fieldValue.CanSet() {
				if field.Type == reflect.TypeOf([]string{}) {
					interfaceSlice, ok := value.([]interface{})
					if ok {
						strSlice := make([]string, len(interfaceSlice))
						for j, v := range interfaceSlice {
							strSlice[j], _ = v.(string)
						}
						fieldValue.Set(reflect.ValueOf(strSlice))
						continue
					}
				}
				fieldValue.Set(reflect.ValueOf(value).Convert(field.Type))
			}
		}
	}
	return book
}

func BooksFromMeiliDocuments(docs []interface{}) []domain.BookSearch {
	books := make([]domain.BookSearch, len(docs))
	for i, doc := range docs {
		books[i] = BookFromMeiliDocument(doc)
	}

	return books
}
