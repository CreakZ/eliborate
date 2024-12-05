package bootch

import (
	"eliborate/internal/models/dto"
	"eliborate/pkg/bootch/parse"
	"sync"
)

func GetBookWithChitaiGorod(wg *sync.WaitGroup, isbn string, books chan dto.BookInfo, errs chan error) {
	defer wg.Done()

	book, err := parse.ParseBookInfoFromChitaiGorod(isbn)
	if err != nil {
		errs <- err
		return
	}
	books <- book
}

func GetBookWithLivelib(wg *sync.WaitGroup, isbn string, books chan dto.BookInfo, errs chan error) {
	defer wg.Done()

	book, err := parse.ParseBookInfoFromLivelib(isbn)
	if err != nil {
		errs <- err
		return
	}
	books <- book
}

func GetBookWithGoogleAPI(wg *sync.WaitGroup, isbn string, books chan dto.BookInfo, errs chan error) {
	defer wg.Done()

	book, err := parse.ParseBookInfoFromGoogleBookApi(isbn)
	if err != nil {
		errs <- err
		return
	}
	books <- book
}
