package validation

import (
	"eliborate/internal/errs"
	"eliborate/internal/models/domain"
	"errors"
)

var (
	errEmptyTitle           = errors.New("empty 'title' provided")
	errEmptyDescription     = errors.New("empty 'description' provided")
	errEmptyAuthors         = errors.New("empty 'authors' provided")
	errEmptyCoverUrls       = errors.New("empty 'cover_urls' provided")
	errWrongRackValue       = errors.New("'rack' value should not be less than 1")
	errWrongShelfValue      = errors.New("'shelf' value should not be less than 1")
	errWrongCategoryIDValue = errors.New("'category_id' value should not be less than 1")
)

func ValidateBookInfo(book domain.BookInfo) error {
	switch {
	case book.Title == "":
		return errs.NewValidationError("title", errEmptyTitle)
	case len(book.Authors) == 0:
		return errs.NewValidationError("authors", errEmptyAuthors)
	}
	return nil
}

func ValidateBookPlacement(book domain.BookPlacement) error {
	if book.Rack < 1 {
		return errs.NewValidationError("rack", errWrongRackValue)
	}
	if book.Shelf < 1 {
		return errs.NewValidationError("shelf", errWrongShelfValue)
	}
	return nil
}

func ValidateBookCreate(book domain.BookCreate) error {
	err := ValidateBookInfo(book.BookInfo)
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

func ValidateUpdateBookInfo(book domain.UpdateBookInfo) error {
	if book.Title != nil && *book.Title == "" {
		return errs.NewValidationError("title", errEmptyTitle)
	}
	if book.Authors != nil && len(book.Authors) == 0 {
		return errs.NewValidationError("authors", errEmptyAuthors)
	}
	if book.Description != nil && *book.Description == "" {
		return errs.NewValidationError("description", errEmptyDescription)
	}
	if book.CategoryID != nil {
		err := validateCategoryID(*book.CategoryID)
		if err != nil {
			return err
		}
	}
	if book.CoverUrls != nil && len(book.CoverUrls) == 0 {
		return errs.NewValidationError("cover_urls", errEmptyCoverUrls)
	}
	return nil
}

func validateCategoryID(categoryID int) error {
	if categoryID < 1 {
		return errs.NewValidationError("category_id", errWrongCategoryIDValue)
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
