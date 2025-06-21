package validation

import (
	"eliborate/internal/models/domain"
	"errors"
)

var (
	ErrWrongRackValue       ValidationError = errors.New("wrong 'rack' value")
	ErrWrongShelfValue      ValidationError = errors.New("wrong 'shelf' value")
	ErrWrongTitleValue      ValidationError = errors.New("wrong 'title' value")
	ErrWrongAuthorsValue    ValidationError = errors.New("wrong 'authors' value")
	ErrWrongCategoryIDValue ValidationError = errors.New("wrong 'category_id' value")
)

func ValidateBookInfo(book domain.BookInfo) ValidationError {
	switch {
	case book.Title == "":
		return wrap(ErrWrongTitleValue, "empty 'title' provided")
	case len(book.Authors) == 0:
		return wrap(ErrWrongAuthorsValue, "empty 'authors' provided")
	case containsEmptyStrings(book.Authors):
		return wrap(ErrWrongAuthorsValue, "'authors' contains empty strings")
	}
	return nil
}

func ValidateBookPlacement(book domain.BookPlacement) ValidationError {
	if book.Rack < 1 {
		return wrap(ErrWrongRackValue, "'rack' value should not be less than 1")
	}
	if book.Shelf < 1 {
		return wrap(ErrWrongShelfValue, "'shelf' value should not be less than 1")
	}
	return nil
}

func ValidateBookCreate(book domain.BookCreate) (err ValidationError) {
	err = ValidateBookInfo(book.BookInfo)
	if err != nil {
		return err
	}

	err = validateCategoryID(book.CategoryID)
	if err != nil {
		return err
	}

	err = ValidateBookPlacement(book.BookPlacement)
	if err != nil {
		return err
	}

	return nil
}

func validateCategoryID(categoryID int) ValidationError {
	if categoryID < 1 {
		return wrap(ErrWrongCategoryIDValue, "'category_id' should not be less than 1")
	}
	return nil
}

func containsEmptyStrings(s []string) bool {
	for i := range s {
		if s[i] == "" {
			return true
		}
	}
	return false
}
