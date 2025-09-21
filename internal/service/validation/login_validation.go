package validation

import (
	"eliborate/internal/errs"
	"errors"
	"regexp"
)

var (
	errWrongLoginLength        = errors.New("login length must be between 3 and 32 characters")
	errLoginContainsWrongChars = errors.New("login can only contain letters, numbers, and underscores")
)

func ValidateLogin(login string) error {
	length := len(login)
	if length < 3 || length > 32 {
		return errs.NewValidationError("login", errWrongLoginLength)
	}
	if matched, _ := regexp.MatchString(`[^a-zA-Z0-9_]`, login); matched {
		return errs.NewValidationError("login", errLoginContainsWrongChars)
	}
	return nil
}
