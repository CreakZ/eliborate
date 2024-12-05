package bootch_test

import (
	bootch "eliborate/pkg/bootch/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO
func TestIsValid(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		Isbn    string
		IsValid bool
	}{
		{"", false},
		{"9679991", false},
		{"-978-0-306-40615-7-", true},
		{"978500155719?", false},
		{"9785824312737", true},
		{"1234567891011", false},
		{"9781111111111", false},
		{"978-2-266-11156-0", true},
		{"978-5-17-090630-7", true},
		{"978-5-00155-719-7", true},
		{"978-0-306-40615-7", true},
	}

	for i, tt := range tests {
		if !a.Equal(tt.IsValid, bootch.IsValid(tt.Isbn)) {
			t.Errorf("isbn '%s': test %d failed", tt.Isbn, i)
		}
	}
}
