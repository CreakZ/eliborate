package errs

import (
	"fmt"
)

type ValidationError struct {
	field string
	err   error
}

func NewValidationError(field string, err error) error {
	return &ValidationError{
		field: field,
		err:   err,
	}
}

func (e *ValidationError) Error() string {
	if e.err == nil {
		return fmt.Sprintf("%s validation error", e.field)
	}
	return fmt.Sprintf("%s: %s", e.field, e.err.Error())
}
