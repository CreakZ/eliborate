package service

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/models/domain"
	"eliborate/internal/repository"
	"eliborate/internal/service/servutils"
	"eliborate/internal/service/servutils/validation"
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
	err := validation.ValidateBookCreate(book)
	if err != nil {
		return 0, err
	}

	bookEntity := convertors.DomainBookCreateToEntity(book)

	bookID, err := b.repo.CreateBook(ctx, bookEntity)
	if err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("create book %v", err.Error()))
		return 0, err
	}

	return bookID, nil
}

func (b bookService) GetBookById(ctx context.Context, id int) (domain.Book, error) {
	if err := validation.ValidateID(id); err != nil {
		return domain.Book{}, err
	}

	book, err := b.repo.GetBookById(ctx, id)
	if err != nil {
		return domain.Book{}, err
	}

	return convertors.EntityBookToDomain(book), nil
}

func (b bookService) GetBooks(ctx context.Context, page, limit int, rack *int, searchQuery *string) ([]domain.Book, error) {
	if err := validation.ValidatePage(page); err != nil {
		return []domain.Book{}, err
	}
	if err := validation.ValidateLimit(limit); err != nil {
		return []domain.Book{}, err
	}
	if err := validation.ValidateRackPtr(rack); err != nil {
		return []domain.Book{}, err
	}
	if err := validation.ValidateSearchQueryPtr(searchQuery); err != nil {
		return []domain.Book{}, err
	}

	offset := servutils.CountOffset(page, limit)

	booksRaw, err := b.repo.GetBooks(ctx, offset, limit, rack, searchQuery)
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

func (b bookService) UpdateBookInfo(ctx context.Context, id int, book domain.UpdateBookInfo) error {
	if err := validation.ValidateID(id); err != nil {
		return err
	}

	bookConv := convertors.UpdateBookInfoToMap(book)

	if err := b.repo.UpdateBookInfo(ctx, id, bookConv); err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("update book info %v", err.Error()))
		return err
	}
	return nil
}

func (b bookService) UpdateBookPlacement(ctx context.Context, id, rack, shelf int) error {
	if err := validation.ValidateID(id); err != nil {
		return err
	}

	if err := b.repo.UpdateBookPlacement(ctx, id, rack, shelf); err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("update book placement %v", err.Error()))
		return err
	}
	return nil
}

func (b bookService) DeleteBook(ctx context.Context, id int) error {
	if err := validation.ValidateID(id); err != nil {
		return err
	}

	if err := b.repo.DeleteBook(ctx, id); err != nil {
		b.log.InfoLogger.Info().Msg(fmt.Sprintf("delete book %v", err.Error()))
		return err
	}

	return nil
}
