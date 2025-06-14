package service

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/models/domain"
	"eliborate/internal/repository"
	"eliborate/pkg/logging"
	"fmt"
)

var ErrWrongCategory = fmt.Errorf("wrong category name")

type bookService struct {
	repo repository.BookRepo
	log  *logging.Log
}

func InitBookService(repo repository.BookRepo, log *logging.Log) BookService {
	return bookService{
		repo: repo,
		log:  log,
	}
}

func (b bookService) CreateBook(ctx context.Context, book domain.BookCreate) (int, error) {
	bookEntity := convertors.DomainBookCreateToEntity(book)

	bookID, err := b.repo.CreateBook(ctx, bookEntity)
	if err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("create book %v", err.Error()))
		return 0, err
	}

	return bookID, nil
}

func (b bookService) GetBookById(ctx context.Context, id int) (domain.Book, error) {
	book, err := b.repo.GetBookById(ctx, id)
	if err != nil {
		return domain.Book{}, err
	}
	return convertors.EntityBookToDomain(book), nil
}

func (b bookService) GetBooks(ctx context.Context, offset, limit int) ([]domain.Book, error) {
	booksRaw, err := b.repo.GetBooks(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	books := make([]domain.Book, 0, len(booksRaw))
	for _, book := range booksRaw {
		books = append(books, convertors.EntityBookToDomain(book))
	}
	return books, nil
}

func (b bookService) GetBooksTotalCount(ctx context.Context) (int, error) {
	count, err := b.repo.GetBooksTotalCount(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (b bookService) GetBooksByRack(ctx context.Context, rack, offset, limit int) ([]domain.Book, error) {
	booksRaw, err := b.repo.GetBooksByRack(ctx, rack, offset, limit)
	if err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("get books by rack failed: %v", err))
		return nil, err
	}

	books := make([]domain.Book, len(booksRaw))
	for i := range booksRaw {
		books[i] = convertors.EntityBookToDomain(booksRaw[i])
	}
	return books, nil
}

func (b bookService) GetBooksByTextSearch(ctx context.Context, text string, offset, limit int) ([]domain.BookSearch, error) {
	booksDomain, err := b.repo.GetBooksByTextSearch(ctx, text, offset, limit)
	if err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("get books by fulltext search failed: %v", err))
		return nil, err
	}

	books := make([]domain.BookSearch, len(booksDomain))
	for i := range booksDomain {
		books[i] = convertors.EntityBookSearchToDomain(booksDomain[i])
	}
	return books, nil
}

func (b bookService) UpdateBookInfo(ctx context.Context, id int, book domain.UpdateBookInfo) error {
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
