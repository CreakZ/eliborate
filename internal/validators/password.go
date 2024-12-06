package validators

import (
	"fmt"
	"unicode"
)

// Password should contain at least 1 digit and upper-case letter and have length 8 to 64 symbols

var (
	ErrShortPassword    = fmt.Errorf("password is too short")
	ErrInsecurePassword = fmt.Errorf("provided password is insecure")
)

func newPasswordSecurityError(message string) error {
	return fmt.Errorf("provided password is insecure: %s", message)
}

func IsPasswordValid(password string) error {
	if plen := len(password); plen < 8 {
		return newPasswordSecurityError("password is too short")
	} else if plen > 64 {
		return newPasswordSecurityError("password is too long")
	}

	if !containsDigits(password) {
		return newPasswordSecurityError("password should contain digits")
	}

	if !containsUpperCase(password) {
		return newPasswordSecurityError("password should contain uppercase letters")
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
