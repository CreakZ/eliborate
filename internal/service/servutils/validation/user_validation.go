package validation

import (
	"eliborate/internal/models/domain"
	"errors"
)

var (
	ErrWrongNameValue     ValidationError = errors.New("wrong 'name' value")
	ErrWrongPasswordValue ValidationError = errors.New("wrong 'password' value")
)

func ValidateUserCreate(user domain.UserCreate) ValidationError {
	if user.Name == "" {
		return wrap(ErrWrongNameValue, "empty 'name' provided")
	}
	if err := ValidateLogin(user.Login); err != nil {
		return err
	}
	if err := ValidatePassword(user.Password); err != nil {
		return err
	}
	return nil
}
