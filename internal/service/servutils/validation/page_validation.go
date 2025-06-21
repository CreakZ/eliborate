package validation

import (
	"errors"
)

var ErrWrongPageValue ValidationError = errors.New("wrong 'limit' value")

func ValidatePage(page int) (err ValidationError) {
	if page < 1 {
		return wrap(ErrWrongPageValue, "'page' value should not be less than 1")
	}
	return nil
}
