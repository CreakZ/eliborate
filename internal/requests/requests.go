package requests

import (
	"fmt"
	"sync"
	dto "yurii-lib/internal/models/dto"
	"yurii-lib/internal/requests/libs"
)

const reqAmount = 2

var ErrNoBooksFound = fmt.Errorf("no books found")

// GetBookByISBN searches book info in 2 services (Читай Город, Google API)
func GetBookByISBN(isbn string) ([]dto.BookInfo, error) {
	wg := sync.WaitGroup{}

	booksChan := make(chan dto.BookInfo, reqAmount)
	errChan := make(chan error, reqAmount)

	wg.Add(reqAmount)

	go libs.GetBookWithGoogleAPI(&wg, isbn, booksChan, errChan)
	go libs.GetBookWithChitaiGorod(&wg, isbn, booksChan, errChan)

	wg.Wait()

	close(booksChan)
	close(errChan)

	if len(errChan) == reqAmount || len(booksChan) == 0 {
		return []dto.BookInfo{}, ErrNoBooksFound
	}

	var books []dto.BookInfo

	for book := range booksChan {
		books = append(books, book)
	}

	return books, nil
}
