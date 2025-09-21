package validation

import "eliborate/internal/errs"

func ValidateRackPtr(rack *int) error {
	if rack != nil && *rack < 1 {
		return errs.NewValidationError("rack", errWrongRackValue)
	}
	return nil
}
