package validation

import "eliborate/internal/errs"

func ValidateShelfPtr(shelf *int) error {
	if shelf != nil && *shelf < 1 {
		return errs.NewValidationError("shelf", errWrongShelfValue)
	}
	return nil
}
