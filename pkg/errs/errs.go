package errs

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNoRowsAffected = fmt.Errorf("no rows affected")

	ErrEmptyTitle    = fmt.Errorf("empty title")
	ErrEmptyCategory = fmt.Errorf("empty category")
	ErrEmptyAuthors  = fmt.Errorf("empty authors")
	ErrWrongRack     = fmt.Errorf("rack value less than 1")
	ErrWrongShelf    = fmt.Errorf("shelf value less than 1")
)

func MergeErrors(caller string, errSlice []string) error {
	base := []string{fmt.Sprintf("%s: multiple errors occured:", caller)}
	base = append(base, errSlice...)

	return errors.New(strings.Join(base, "\n\t"))
}
