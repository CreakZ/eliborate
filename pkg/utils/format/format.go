package format

import (
	"fmt"
	"strings"
)

var ErrWrongISBN = fmt.Errorf("wrong isbn provided")

// FormatISBN formats ISBN by deleting all hyphens ('-').
// If provided string is supposedly ISBN-10, FormatISBN adds '978' at the beginning.
// If provided string is neither ISBN-10 nor ISBN-13, FormatISBN returns empty string and ErrWrongISBN
func FormatISBN(isbn string) (string, error) {
	clean := strings.ReplaceAll(isbn, "-", "")

	switch len(clean) {
	case 10:
		return "978" + clean, nil
	case 13:
		return clean, nil
	default:
		return "", ErrWrongISBN
	}
}
