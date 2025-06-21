package validation

func ValidateRackPtr(rack *int) ValidationError {
	if rack == nil {
		return nil
	}
	if *rack < 1 {
		return wrap(ErrWrongRackValue, "'rack' value should not be less than 1")
	}
	return nil
}
