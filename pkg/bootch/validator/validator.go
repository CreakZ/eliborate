package validator

import (
	"strings"
	"unicode"
)

func IsValid(isbn string) bool {
	isbn = cleanIsbn(isbn)

	if strings.ContainsFunc(isbn, func(r rune) bool {
		return !unicode.IsDigit(r)
	}) {
		return false
	}

	if len(isbn) != 10 && len(isbn) != 13 {
		return false
	}

	return isIsbnSumValid(isbn)
}

func cleanIsbn(isbn string) string {
	return strings.ReplaceAll(isbn, "-", "")
}

func isIsbnSumValid(isbn string) bool {
	sum := 0
	isbnLen := len(isbn)

	switch isbnLen {
	case 10:
		for i, c := range isbn {
			charValue := int(c - '0')
			sum += charValue * (isbnLen - i)
		}
		return sum%11 == 0

	case 13:
		last := int(isbn[isbnLen-1] - '0')
		for i, c := range isbn[:isbnLen-1] {
			charValue := int(c - '0')
			if i%2 != 0 {
				charValue *= 3
			}
			sum += charValue
		}
		return (10-(sum%10))%10 == last

	default:
		return false
	}
}
