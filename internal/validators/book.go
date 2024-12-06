package validators

import (
	"eliborate/internal/errs"
	"eliborate/internal/models/dto"
	"fmt"
)

type Result struct {
	Ok  bool
	Err error
}

func NewOkResult() Result {
	return Result{
		Ok:  true,
		Err: nil,
	}
}

func NewErrResult(message string) Result {
	return Result{
		Ok:  false,
		Err: fmt.Errorf(message),
	}
}

func NewErrResultFromErr(err error) Result {
	return Result{
		Ok:  false,
		Err: err,
	}
}

func ValidateBookPlacement(book *dto.BookPlacement) Result {
	if book.Rack < 1 {
		return NewErrResultFromErr(errs.ErrBookWrongRack)
	}
	if book.Shelf < 1 {
		return NewErrResultFromErr(errs.ErrBookWrongShelf)
	}
	return NewOkResult()
}

func ValidateBookInfo(book *dto.BookInfo) Result {
	switch {
	case book.Title == "":
		return NewErrResultFromErr(errs.ErrBookEmptyTitle)
	case len(book.Authors) == 0 || containsEmptyStrings(book.Authors):
		return NewErrResultFromErr(errs.ErrBookEmptyAuthors)
	case book.Category == "":
		return NewErrResultFromErr(errs.ErrBookEmptyCategory)
	}
	return NewOkResult()
}

func ValidateBookCreate(book *dto.BookCreate) Result {
	res := ValidateBookInfo(&book.BookInfo)
	if !res.Ok {
		return res
	}
	res = ValidateBookPlacement(&book.BookPlacement)
	if !res.Ok {
		return res
	}
	return NewOkResult()
}

func containsEmptyStrings(s []string) bool {
	for i := range s {
		if s[i] == "" {
			return true
		}
	}
	return false
}
