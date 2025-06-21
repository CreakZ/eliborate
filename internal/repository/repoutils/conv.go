package repoutils

import (
	"eliborate/internal/models/entity"
)

func ConvertMeiliHitsToIntSlice(hits []any) []int {
	if len(hits) == 0 {
		return []int{}
	}

	mapSlice := convertRawHitsToMapSlice(hits)

	ints := make([]int, 0, len(mapSlice))

	for _, hit := range mapSlice {
		idAny, ok := hit["id"]
		if !ok {
			continue
		}

		id, ok := idAny.(float64)
		if !ok {
			continue
		}

		ints = append(ints, int(id))
	}

	return ints
}

func convertRawHitsToMapSlice(hits []any) []map[string]any {
	mapSlice := make([]map[string]any, 0, len(hits))

	for _, rawHit := range hits {
		hit, ok := rawHit.(map[string]any)
		if !ok {
			continue
		}
		mapSlice = append(mapSlice, hit)
	}

	return mapSlice
}

func ConvertEntityBookSearchFromEntityBookCreate(bookId int, book entity.BookCreate) entity.BookSearch {
	return entity.BookSearch{
		ID:          bookId,
		Title:       book.Title,
		Authors:     book.Authors,
		Description: book.Description,
		CategoryID:  book.CategoryID,
		Rack:        book.Rack,
	}
}
