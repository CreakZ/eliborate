package validators_test

import (
	"testing"
	"yurii-lib/internal/validators"
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
		if matched := validators.ValidatePassword(pass); !matched {
			t.Errorf("%s string is not valid password\n", pass)
			continue
		}
	}
}
