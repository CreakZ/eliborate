package responses

import "eliborate/internal/models/dto"

type BookSearchResponse struct {
	TotalItems int              `json:"total_items"`
	Items      []dto.BookSearch `json:"items"`
}

func NewBookSearchResponse(items []dto.BookSearch) BookSearchResponse {
	return BookSearchResponse{
		TotalItems: len(items),
		Items:      items,
	}
}
