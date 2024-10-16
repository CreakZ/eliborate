package validators

import (
	"yurii-lib/internal/models/dto"
	"yurii-lib/pkg/errs"
)

func ValidateBookPlacement(book *dto.BookPlacement) (bool, error) {
	switch {
	case book.Title == "":
		return false, errs.ErrEmptyTitle
	case book.Category == "":
		return false, errs.ErrEmptyCategory
	case len(book.Authors) == 0 || emptyStrings(book.Authors):
		return false, errs.ErrEmptyAuthors
	case book.Rack < 1:
		return false, errs.ErrWrongRack
	case book.Rack < 1:
		return false, errs.ErrWrongShelf
	default:
		return true, nil
	}
}

func emptyStrings(s []string) bool {
	for i := range s {
		if s[i] == "" {
			return true
		}
	}

	return false
}
