package validation

import (
	"errors"
	"fmt"
)

type ValidationError error

func NewValidationError(message string) ValidationError {
	return errors.New(message)
}

func wrap(err ValidationError, message string) ValidationError {
	return fmt.Errorf("%w: %s", err, message)
}
