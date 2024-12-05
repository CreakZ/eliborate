package responses

import (
	"eliborate/internal/models/dto"
)

type BookSearchResponse struct {
	TotalItems int              `json:"total_items"`
	Items      []dto.BookSearch `json:"items"`
}

type BookPaginationResponse struct {
	Page       int        `json:"page"`
	TotalPages int        `json:"total_pages"`
	Limit      int        `json:"limit"`
	Items      []dto.Book `json:"items"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func NewBookSearchResponse(items []dto.BookSearch) BookSearchResponse {
	return BookSearchResponse{
		TotalItems: len(items),
		Items:      items,
	}
}

func NewBookPaginationResponse(page, totalPages, limit int, items []dto.Book) BookPaginationResponse {
	return BookPaginationResponse{
		Page:       page,
		TotalPages: totalPages,
		Limit:      limit,
		Items:      items,
	}
}

func NewMessageResponse(message string) MessageResponse {
	return MessageResponse{
		Message: message,
	}
}

func NewMessageResponseFromErr(err error) MessageResponse {
	return MessageResponse{
		Message: err.Error(),
	}
}
