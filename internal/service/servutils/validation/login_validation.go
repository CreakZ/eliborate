package validation

import (
	"errors"
	"regexp"
)

var ErrWrongLoginValue ValidationError = errors.New("wrong 'login' value")

func ValidateLogin(login string) ValidationError {
	length := len(login)
	if length < 3 || length > 32 {
		return wrap(ErrWrongLoginValue, "login length must be between 3 and 32 characters")
	}

	if matched, _ := regexp.MatchString(`[^a-zA-Z0-9_]`, login); matched {
		return wrap(ErrWrongLoginValue, "login can only contain letters, numbers, and underscores")
	}

	return nil
}
