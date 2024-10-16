package libs_test

import (
	"testing"
	"yurii-lib/internal/requests/libs"
)

func TestCleanify(t *testing.T) {
	s := []string{
		"  some \n       ",
		"     ",
		"\n\n\n",
		"\n\nsome\n\n",
		"any     ",
		"     any",
		" any",
		"any ",
		"some",
	}

	expect := []string{
		"some",
		"",
		"",
		"some",
		"any",
		"any",
		"any",
		"any",
		"some",
	}

	for i := range s {
		res := libs.Cleanify(s[i])

		if res != expect[i] {
			t.Errorf("error in string '%s'", s[i])
		}
	}
}
