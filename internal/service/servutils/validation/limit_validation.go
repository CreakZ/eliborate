package validation

import (
	"errors"
)

var ErrWrongLimitValue ValidationError = errors.New("wrong 'limit' value")

func ValidateLimit(limit int) (err ValidationError) {
	if !(limit == 10 || limit == 20 || limit == 50 || limit == 100) {
		return wrap(ErrWrongLimitValue, "allowed values of 'limit': 10, 20, 50, 100")
	}
	return nil
}
