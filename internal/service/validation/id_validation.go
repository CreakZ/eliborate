package validation

import (
	"eliborate/internal/errs"
	"errors"
)

var errWrongIDValue = errors.New("'id' value should not be less than 1")

func ValidateID(id int) error {
	if id < 1 {
		return errs.NewValidationError("id", errWrongIDValue)
	}
	return nil
}
