package validators_test

import (
	"eliborate/internal/validators"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	passwords := []string{
		"qwerty",
		"Cr3akZzzz",
		"popka123",
		"R1zzMusic",
		"AHAHAHAHAHA",
	}

	for _, pass := range passwords {
		if err := validators.IsPasswordValid(pass); err != nil {
			t.Errorf("%s string is not valid password: %s\n", pass, err.Error())
			continue
		}
	}
}
