package service

import (
	"context"
	"fmt"
	"yurii-lib/internal/convertors"
	"yurii-lib/internal/models/dto"
	"yurii-lib/internal/repository"
	"yurii-lib/pkg/log"
)

var ErrWrongCategory = fmt.Errorf("wrong category name")

type bookService struct {
	repo repository.BookRepo
	log  *log.Log
}

func InitBookService(repo repository.BookRepo, log *log.Log) BookService {
	return bookService{
		repo: repo,
		log:  log,
	}
}

func (b bookService) CreateBook(ctx context.Context, book dto.BookPlacement) (int, error) {
	bookConv := convertors.ToDomainBookPlacement(book)

	if bookConv.Category == -1 {
		b.log.InfoLogger.Info().Msg("")
		return 0, ErrWrongCategory
	}

	userID, err := b.repo.CreateBook(ctx, bookConv)
	if err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("create book %v", err.Error()))
		return 0, err
	}

	return userID, nil
}

func (b bookService) GetBooks(ctx context.Context, page, limit int) ([]dto.Book, error) {
	booksRaw, err := b.repo.GetBooks(ctx, page, limit)
	if err != nil {
		return []dto.Book{}, err
	}

	books := make([]dto.Book, len(booksRaw))

	for i := range booksRaw {
		books[i] = convertors.ToDtoBook(booksRaw[i])
	}

	return books, nil
}

func (b bookService) GetBooksByRack(ctx context.Context, rack int) ([]dto.Book, error) {
	booksRaw, err := b.repo.GetBooksByRack(ctx, rack)
	if err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("get book by rack %v", err.Error()))
		return []dto.Book{}, err
	}

	var books = make([]dto.Book, len(booksRaw))

	for i := range booksRaw {
		books[i] = convertors.ToDtoBook(booksRaw[i])
	}

	return books, err
}

func (b bookService) GetBooksByTextSearch(ctx context.Context, text string) ([]dto.Book, error) {
	booksRaw, err := b.repo.GetBooksByTextSearch(ctx, text)
	if err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("get book by fulltext search %s", err.Error()))
		return []dto.Book{}, err
	}

	var books = make([]dto.Book, len(booksRaw))

	for i := range booksRaw {
		books[i] = convertors.ToDtoBook(booksRaw[i])
	}

	return books, nil
}

func (b bookService) UpdateBookInfo(ctx context.Context, id int, book dto.UpdateBookInfo) error {
	bookConv := convertors.UpdateBookInfoToMap(book)

	if err := b.repo.UpdateBookInfo(ctx, id, bookConv); err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("update book info %v", err.Error()))
		return err
	}

	return nil
}

func (b bookService) UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error {
	if err := b.repo.UpdateBookPlacement(ctx, id, rack, shelf); err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("update book placement %v", err.Error()))
		return err
	}

	return nil
}

func (b bookService) DeleteBook(ctx context.Context, id int) error {
	if err := b.repo.DeleteBook(ctx, id); err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("delete book %v", err.Error()))
		return err
	}

	return nil
}
