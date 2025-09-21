package validation

import (
	"eliborate/internal/errs"
	"errors"
)

var errWrongPageValue = errors.New("'page' value should not be less than 1")

func ValidatePage(page int) error {
	if page < 1 {
		return errs.NewValidationError("page", errWrongPageValue)
	}
	return nil
}
