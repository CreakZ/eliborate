package validation

import (
	"eliborate/internal/errs"
	"errors"
)

var errEmptySearchQuery = errors.New("empty 'search_query' provided")

func ValidateSearchQueryPtr(searchQuery *string) error {
	if searchQuery != nil && *searchQuery == "" {
		return errs.NewValidationError("search_query", errEmptySearchQuery)
	}
	return nil
}
