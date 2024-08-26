package format_test

import (
	"testing"
	"yurii-lib/pkg/utils/format"
)

func TestFormatISBN(t *testing.T) {
	isbns := []string{
		"5-218-00723-4",
		"5218007234",
		"3-123-54531-1",
		"1-34-34343-5",
		"134343435",
		"1234567890",
		"978-5-227-02547-0",
		"9785227025470",
	}

	expect := []string{
		"9785218007234",
		"9785218007234",
		"9783123545311",
		"",
		"",
		"9781234567890",
		"9785227025470",
		"9785227025470",
	}

	for i := range isbns {
		res, err := format.FormatISBN(isbns[i])
		if err != nil {
			t.Log(err.Error())
		}

		if res != expect[i] {
			t.Errorf("isbn %s err: %s", isbns[i], err.Error())
		}
	}
}
