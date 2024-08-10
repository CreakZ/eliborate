package errs

import "fmt"

var (
	ErrNoRowsAffected = fmt.Errorf("no rows affected")

	ErrEmptyTitle    = fmt.Errorf("empty title")
	ErrEmptyCategory = fmt.Errorf("empty category")
	ErrEmptyAuthors  = fmt.Errorf("empty authors")
	ErrWrongRack     = fmt.Errorf("rack value less than 1")
	ErrWrongShelf    = fmt.Errorf("shelf value less than 1")
)
