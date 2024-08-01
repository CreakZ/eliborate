package validators

import (
	"fmt"
	"yurii-lib/internal/models/dto"
)

var (
	ErrEmptyTitle    = fmt.Errorf("empty title")
	ErrEmptyCategory = fmt.Errorf("empty category")
	ErrEmptyAuthors  = fmt.Errorf("empty authors")
	ErrWrongRack     = fmt.Errorf("rack value less than 1")
	ErrWrongShelf    = fmt.Errorf("shelf value less than 1")
)

func ValidBookPlacement(book *dto.BookPlacement) (bool, error) {
	switch {
	case book.Title == "":
		return false, ErrEmptyTitle
	case book.Category == "":
		return false, ErrEmptyCategory
	case len(book.Authors) == 0 || emptyStrings(book.Authors):
		return false, ErrEmptyAuthors
	case book.Rack < 1:
		return false, ErrWrongRack
	case book.Rack < 1:
		return false, ErrWrongShelf
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
