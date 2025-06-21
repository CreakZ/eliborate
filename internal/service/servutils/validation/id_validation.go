package validation

import (
	"errors"
)

var ErrWrongIDValue ValidationError = errors.New("wrong 'id' param")

func ValidateID(id int) (err ValidationError) {
	if id < 1 {
		return wrap(ErrWrongPageValue, "'id' value should not be less than 1")
	}
	return nil
}
