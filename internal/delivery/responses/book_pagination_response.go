package responses

import "eliborate/internal/models/dto"

type BookPaginationResponse struct {
	Page        int        `json:"page,omitempty"`
	Limit       int        `json:"limit,omitempty"`
	TotalPages  int        `json:"total_pages,omitempty"`
	SearchQuery string     `json:"search_query,omitempty"`
	Rack        int        `json:"rack,omitempty"`
	Books       []dto.Book `json:"books"`
}

func NewBookPaginationResponse(books []dto.Book) *BookPaginationResponse {
	return &BookPaginationResponse{
		Books: books,
	}
}

func (b *BookPaginationResponse) WithPage(page int) *BookPaginationResponse {
	b.Page = page
	return b
}

func (b *BookPaginationResponse) WithLimit(limit int) *BookPaginationResponse {
	b.Limit = limit
	return b
}

func (b *BookPaginationResponse) WithTotalPages(totalPages int) *BookPaginationResponse {
	b.TotalPages = totalPages
	return b
}

func (b *BookPaginationResponse) WithRack(rack int) *BookPaginationResponse {
	b.Rack = rack
	return b
}

func (b *BookPaginationResponse) WithSearchQuery(searchQuery string) *BookPaginationResponse {
	b.SearchQuery = searchQuery
	return b
}
