package bootch

import (
	"eliborate/internal/errs"
	dto "eliborate/internal/models/dto"
	"eliborate/pkg/bootch/validator"
	"sync"
)

// GetBookByISBN searches book info in 3 services (Читай Город, Google API, Livelib)
func GetBookByISBN(isbn string) ([]dto.BookInfo, error) {
	if !validator.IsValid(isbn) {
		return []dto.BookInfo{}, errs.ErrBookWrongISBN
	}

	wg := sync.WaitGroup{}

	booksChan := make(chan dto.BookInfo, 3)
	errChan := make(chan error, 3)

	wg.Add(3)

	go GetBookWithGoogleAPI(&wg, isbn, booksChan, errChan)
	go GetBookWithChitaiGorod(&wg, isbn, booksChan, errChan)
	go GetBookWithLivelib(&wg, isbn, booksChan, errChan)

	wg.Wait()

	close(booksChan)
	close(errChan)

	if len(errChan) == 3 || len(booksChan) == 0 {
		return []dto.BookInfo{}, errs.ErrBooksNotFound
	}

	books := make([]dto.BookInfo, 0, len(booksChan))
	for book := range booksChan {
		books = append(books, book)
	}

	return books, nil
}
