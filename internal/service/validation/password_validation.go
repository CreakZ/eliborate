package validation

import (
	"eliborate/internal/errs"
	"errors"
	"unicode"
)

// Password should contain at least 1 digit and upper-case letter and have length 8 to 64 symbols

var (
	errShortPassword                  = errors.New("password is too short")
	errLongPassword                   = errors.New("password is too long")
	errPasswordShouldContainDigits    = errors.New("password should contain digits")
	errPasswordShouldContainUppercase = errors.New("password should contain uppercase letters")
)

func ValidatePassword(password string) error {
	plen := len(password)
	if plen < 8 {
		return errs.NewValidationError("password", errShortPassword)
	}
	if plen > 64 {
		return errs.NewValidationError("password", errLongPassword)
	}

	if !containsDigits(password) {
		return errs.NewValidationError("password", errPasswordShouldContainDigits)
	}
	if !containsUpperCase(password) {
		return errs.NewValidationError("password", errPasswordShouldContainUppercase)
	}

	return nil
}

func containsUpperCase(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func containsDigits(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}
