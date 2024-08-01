package requests_test

import (
	"fmt"
	"os"
	"testing"
	dto "yurii-lib/internal/models/dto"
	"yurii-lib/internal/requests"
)

func TestGetBookByISBN(t *testing.T) {
	isbns := []string{
		"978-5-17-154163-7",
		"9785171541637",
		"978-5-17-156045-4",
		"9785041781521",
		"978-1-23-456789-0",
		"123-4-56-789012-3",
		"123-4-56-78012-3",
		"123456789-0",
		"978-5-4224-0970-9",
	}

	type meta struct {
		Books []dto.BookInfo
		ISBN  string
	}

	var booksAndMeta []meta

	for _, isbn := range isbns {
		books, err := requests.GetBookByISBN(isbn)
		if err != nil {
			t.Errorf("an error '%s' occured while fetching book with isbn: %s", err.Error(), isbn)
		}

		booksAndMeta = append(booksAndMeta, meta{
			Books: books,
			ISBN:  isbn,
		})
	}

	file, _ := os.Create("result.txt")
	defer file.Close()

	for i, meta := range booksAndMeta {
		var str string
		for _, book := range meta.Books {
			str = fmt.Sprintf("Book #%d: isbn: %s\n\nAuthors: %v\nCategory: %s\n", i+1, meta.ISBN, book.Authors, book.Category)
			if book.Description != nil {
				str = fmt.Sprintf("%sDescription: %s", str, *book.Description)
			} else {
				str = fmt.Sprintf("%sNo description", str)
			}
			str = fmt.Sprintf("%sForeign: %v\nLogo: %v\nTitle: %s\n\n", str, book.IsForeign, book.Logo, book.Title)
		}

		file.Write([]byte(str))
	}
}
