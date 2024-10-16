package validators

import (
	"fmt"
	"unicode"
)

// Пароль должен содержать цифры и заглавные буквы, а также иметь длину от 8 до 50 символов

var (
	ErrShortPassword    = fmt.Errorf("password is too short")
	ErrInsecurePassword = fmt.Errorf("provided password is insecure")
)

func ValidatePassword(password string) bool {
	if len(password) < 8 || len(password) > 50 {
		return false
	}

	if !(containsDigits(password) && containsUpperCase(password)) {
		return false
	}

	return true
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
