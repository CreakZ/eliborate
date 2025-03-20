package convertors

import (
	"eliborate/internal/models/entity"
	"reflect"
)

func EntityBookSearchFromEntityBookCreate(bookId int, book entity.BookCreate) entity.BookSearch {
	desc := ""
	if book.Description.Valid {
		desc = book.Description.String
	}
	return entity.BookSearch{
		ID:          bookId,
		Title:       book.Title,
		Authors:     book.Authors,
		Description: desc,
		Category:    book.Category,
	}
}

func MeiliDocumentFromBookSearch(primaryKey int, bookSearch entity.BookSearch) map[string]any {
	switch reflect.TypeOf(bookSearch).Kind() {
	case reflect.Struct:
		fields := reflect.VisibleFields(reflect.TypeOf(bookSearch))
		values := reflect.ValueOf(bookSearch)

		doc := make(map[string]any)
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
		return map[string]any{}
	}
}

// doc is expected to be map[string]interface{}
func bookFromMeiliDocument(doc any) entity.BookSearch {
	docMap, ok := doc.(map[string]any)
	if !ok {
		return entity.BookSearch{}
	}

	var book entity.BookSearch
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

func BooksFromMeiliDocuments(docs []any) []entity.BookSearch {
	books := make([]entity.BookSearch, len(docs))
	for i, doc := range docs {
		books[i] = bookFromMeiliDocument(doc)
	}

	return books
}
