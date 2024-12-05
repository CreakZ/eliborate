package errs

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoRowsAffected = fmt.Errorf("no rows affected")

	ErrUserAlreadyExists = fmt.Errorf("user already exists")
	ErrLastAdminUser     = fmt.Errorf("there cannot be less than 1 admin user")

	ErrBookEmptyTitle    = fmt.Errorf("empty title")
	ErrBookEmptyCategory = fmt.Errorf("empty category")
	ErrBookEmptyAuthors  = fmt.Errorf("empty authors")
	ErrBookWrongRack     = fmt.Errorf("rack value less than 1")
	ErrBookWrongShelf    = fmt.Errorf("shelf value less than 1")
	ErrBookWrongISBN     = fmt.Errorf("wrong isbn provided")
	ErrBooksNotFound     = fmt.Errorf("no books found")
)

func MergeErrors(caller string, errSlice []string) error {
	base := []string{fmt.Sprintf("%s: multiple errors occured:", caller)}
	base = append(base, errSlice...)

	return errors.New(strings.Join(base, "\n\t"))
}
