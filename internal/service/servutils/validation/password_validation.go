package validation

import (
	"errors"
	"unicode"
)

// Password should contain at least 1 digit and upper-case letter and have length 8 to 64 symbols

var (
	ErrShortPassword    = errors.New("password is too short")
	ErrInsecurePassword = errors.New("provided password is insecure")
)

func ValidatePassword(password string) ValidationError {
	if plen := len(password); plen < 8 {
		return ErrShortPassword
	} else if plen > 64 {
		return wrap(ErrInsecurePassword, "password is too long")
	}

	if !containsDigits(password) {
		return wrap(ErrInsecurePassword, "password should contain digits")
	}

	if !containsUpperCase(password) {
		return wrap(ErrInsecurePassword, "password should contain uppercase letters")
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
