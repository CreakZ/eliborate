package handlers

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/delivery/responses"
	"eliborate/internal/errs"
	"eliborate/internal/models/dto"
	"eliborate/internal/service"
	"eliborate/pkg/storage"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandlers struct {
	service service.BookService
	cache   *storage.RedisCache
}

func InitBookHandlers(service service.BookService, cache *storage.RedisCache) BookHandlers {
	return bookHandlers{
		service: service,
		cache:   cache,
	}
}

// @Summary Create a new book
// @Description Create a new book entry in the system
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.BookCreate true "Book Create"
// @Security BearerAuth
// @Success 200 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books [post]
func (b bookHandlers) CreateBook(c *gin.Context) {
	var book dto.BookCreate

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	bookDomain := convertors.DtoBookCreateToDomain(book)

	id, err := b.service.CreateBook(context.Background(), bookDomain)
	if err != nil {
		var vErr *errs.ValidationError
		switch {
		case errors.As(err, &vErr):
			c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse(vErr.Error()))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusCreated, responses.NewBookCreateResponse(id))
}

// @Summary Get book by id
// @Description Retrieve book information by its id
// @Tags books
// @Param id path int true "Book id"
// @Produce json
// @Success 200 {array}  dto.Book
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books/{id} [get]
func (b bookHandlers) GetBookById(c *gin.Context) {
	idRaw := c.Param("id")
	if idRaw == "" {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("no id number provided"))
		return
	}

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	book, err := b.service.GetBookById(context.Background(), id)
	if err != nil {
		var vErr *errs.ValidationError
		switch {
		case errors.As(err, &vErr):
			c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse(vErr.Error()))
		case errors.Is(err, errs.ErrEntityNotFound):
			c.JSON(http.StatusNotFound, responses.NewMessageResponse(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, convertors.DomainBookToDto(book))
}

// @Summary Get a paginated list of books
// @Description Retrieve a list of books with pagination, optional rack filtering and full-text search
// @Tags books
// @Param page query int true "Page number"
// @Param limit query int true "Books limit per page (10, 20, 50 or 100)"
// @Param rack query int false "Rack number to filter books"
// @Param search_query query string false "Full-text search query"
// @Produce json
// @Success 200 {object} responses.BookPaginationResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 422 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books [get]
func (b bookHandlers) GetBooks(c *gin.Context) {
	pageRaw := c.Query("page")
	limitRaw := c.Query("limit")

	page, err := strconv.Atoi(pageRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("wrong 'page' param format"))
		return
	}

	limit, err := strconv.Atoi(limitRaw)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("wrong 'limit' param format"))
		return
	}

	var (
		rackPtr        *int
		searchQueryPtr *string
	)

	query, ok := c.GetQuery("search_query")
	if ok {
		searchQueryPtr = &query
	}

	rackRaw, ok := c.GetQuery("rack")
	if ok {
		rack, err := strconv.Atoi(rackRaw)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.NewMessageResponse("wrong 'rack' param format"))
			return
		}
		rackPtr = &rack
	}

	books, err := b.service.GetBooks(c.Request.Context(), page, limit, rackPtr, searchQueryPtr)
	if err != nil {
		var vErr *errs.ValidationError
		switch {
		case errors.As(err, &vErr):
			c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse(vErr.Error()))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		}
		return
	}

	response := responses.NewBookPaginationResponse(convertors.DomainBooksToDto(books)).
		WithPage(page).
		WithLimit(limit)

	if rackPtr != nil {
		response = response.WithRack(*rackPtr)
	}
	if searchQueryPtr != nil {
		response = response.WithSearchQuery(*searchQueryPtr)
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Update book information
// @Description Update details of an existing book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book id"
// @Param book body dto.UpdateBookInfo true "Update Book Info"
// @Security BearerAuth
// @Success 200 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books/{id}/info [patch]
func (b bookHandlers) UpdateBookInfo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	var updateBook dto.UpdateBookInfo
	if err = c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	updateBookDomain := convertors.DtoUpdateBookInfoToDomain(updateBook)

	if err = b.service.UpdateBookInfo(c.Request.Context(), id, updateBookDomain); err != nil {
		var vErr *errs.ValidationError
		switch {
		case errors.As(err, &vErr):
			c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse(vErr.Error()))
		case errors.Is(err, errs.ErrEntityNotFound):
			c.JSON(http.StatusNotFound, responses.NewMessageResponse("book not found"))
		case errors.Is(err, errs.ErrNoDataSentToUpdate):
			c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse("no data sent to update"))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
}

// @Summary Update book placement
// @Description Update the rack and shelf placement of a book
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book id"
// @Param book body dto.UpdateBookPlacement true "pdate Book Placement"
// @Security BearerAuth
// @Success 200 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books/{id}/placement [patch]
func (b bookHandlers) UpdateBookPlacement(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	var updateBookDto dto.UpdateBookPlacement
	if err := c.ShouldBindJSON(&updateBookDto); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	updateBook := convertors.DtoUpdateBookPlacementToDomain(updateBookDto)

	err = b.service.UpdateBookPlacement(c.Request.Context(), id, updateBook)
	if err != nil {
		var vErr *errs.ValidationError
		switch {
		case errors.As(err, &vErr):
			c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse(vErr.Error()))
		case errors.Is(err, errs.ErrEntityNotFound):
			c.JSON(http.StatusNotFound, responses.NewMessageResponse("book not found"))
		case errors.Is(err, errs.ErrNoDataSentToUpdate):
			c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse("no data sent to update"))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
}

// @Summary Delete a book by its ID
// @Description Remove a book from the system using its ID
// @Tags books
// @Param id path int true "Book ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 404 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books/{id} [delete]
func (b bookHandlers) DeleteBook(c *gin.Context) {
	rawID := c.Param("id")
	if rawID == "" {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("no 'id' param provided"))
		return
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	if err = b.service.DeleteBook(c.Request.Context(), id); err != nil {
		var vErr *errs.ValidationError
		switch {
		case errors.As(err, &vErr):
			c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse(vErr.Error()))
		case errors.Is(err, errs.ErrEntityNotFound):
			c.JSON(http.StatusNotFound, responses.NewMessageResponse("book not found"))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
}
