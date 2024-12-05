package parse_test

import (
	"eliborate/pkg/bootch/parse"
	"testing"
	"time"
)

var isbns = []string{
	"1234567890123",
	"9789781233333",
	"978-0-307-26398-8",
	"97899071539191",
	"978-5-00155-719-7",
	"978-5-17-090630-7",
}

func TestParseBookInfoFromChitaiGorod(t *testing.T) {
	for _, isbn := range isbns {
		start := time.Now()
		book, err := parse.ParseBookInfoFromChitaiGorod(isbn)
		if err != nil {
			t.Errorf("isbn '%s': %s\n", isbn, err.Error())
			t.Logf("time taken: %s", time.Since(start).String())
			continue
		}
		since := time.Since(start).String()
		t.Logf("book with '%s' isbn parsed successfully in %s:\n", isbn, since)
		t.Logf("%+v", book)
	}
}
