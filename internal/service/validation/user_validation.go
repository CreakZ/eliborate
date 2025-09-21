package validation

import (
	"eliborate/internal/errs"
	"eliborate/internal/models/domain"
	"errors"
)

var errEmptyName = errors.New("empty 'name' provided")

func ValidateUserCreate(user domain.UserCreate) error {
	if user.Name == "" {
		return errs.NewValidationError("name", errEmptyName)
	}
	if err := ValidateLogin(user.Login); err != nil {
		return err
	}
	if err := ValidatePassword(user.Password); err != nil {
		return err
	}
	return nil
}
