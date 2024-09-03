package errs

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoRowsAffected = fmt.Errorf("no rows affected")

	ErrUserAlreadyExists = fmt.Errorf("user already exists")

	ErrEmptyTitle    = fmt.Errorf("empty title")
	ErrEmptyCategory = fmt.Errorf("empty category")
	ErrEmptyAuthors  = fmt.Errorf("empty authors")
	ErrWrongRack     = fmt.Errorf("rack value less than 1")
	ErrWrongShelf    = fmt.Errorf("shelf value less than 1")

	ErrLastAdminUser = fmt.Errorf("there cannot be less than 1 admin user")
)

func MergeErrors(caller string, errSlice []string) error {
	base := []string{fmt.Sprintf("%s: multiple errors occured:", caller)}
	base = append(base, errSlice...)

	return errors.New(strings.Join(base, "\n\t"))
}
