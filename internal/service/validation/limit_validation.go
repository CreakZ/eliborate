package validation

import (
	"eliborate/internal/errs"
	"errors"
)

var errWrongLimitValue = errors.New("allowed values of 'limit': 10, 20, 50, 100")

func ValidateLimit(limit int) error {
	if !(limit == 10 || limit == 20 || limit == 50 || limit == 100) {
		return errs.NewValidationError("limit", errWrongLimitValue)
	}
	return nil
}
