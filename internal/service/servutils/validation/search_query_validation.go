package validation

import "errors"

var ErrEmptySearchQuery = errors.New("empty 'search_query' provided")

func ValidateSearchQueryPtr(searchQuery *string) ValidationError {
	if searchQuery == nil {
		return nil
	}
	if *searchQuery == "" {
		return ErrEmptySearchQuery
	}
	return nil
}
