package handlers

import (
	"context"
	"eliborate/internal/constants"
	"eliborate/internal/convertors"
	"eliborate/internal/delivery/responses"
	"eliborate/internal/errs"
	"eliborate/internal/models/dto"
	"eliborate/internal/service"
	"eliborate/internal/validators"
	"eliborate/pkg/storage"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

	result := validators.ValidateBookCreate(&book)
	if !result.Ok {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(result.Err.Error()))
		return
	}

	bookDomain := convertors.DtoBookCreateToDomain(book)

	_, err := b.service.CreateBook(context.Background(), bookDomain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	_, err = b.cache.GetInt(constants.RedisTotalBooks)
	if errors.Is(err, redis.Nil) {
		totalBooks, err := b.service.GetBooksTotalCount(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
			return
		}
		b.cache.SetInt(constants.RedisTotalBooks, totalBooks)
	} else {
		b.cache.Incr(constants.RedisTotalBooks)
	}

	c.JSON(http.StatusCreated, responses.NewSuccessMessageResponse())
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
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
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
// @Param text query string false "Full-text search query"
// @Produce json
// @Success 200 {object} responses.BookPaginationResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 422 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books [get]
func (b bookHandlers) GetBooks(c *gin.Context) {
	// should hide all the strategy inside service layer later
	pageRaw := c.Query("page")
	limitRaw := c.Query("limit")
	text := c.Query("text")
	rackRaw := c.Query("rack")

	page, err := strconv.Atoi(pageRaw)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("invalid 'page' param"))
		return
	}

	limit, err := strconv.Atoi(limitRaw)
	if err != nil || !(limit == 10 || limit == 20 || limit == 50 || limit == 100) {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("invalid 'limit' param (allowed: 10, 20, 50, 100)"))
		return
	}

	offset := (page - 1) * limit
	ctx := c.Request.Context()

	if text != "" {
		if err := validators.ValidateTextQuery(text); err != nil {
			c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
			return
		}

		books, err := b.service.GetBooksByTextSearch(ctx, text, offset, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
			return
		}

		booksDto := convertors.DomainBooksSearchToDto(books)

		c.JSON(http.StatusOK, responses.NewBookSearchResponse(booksDto))
		return
	}

	if rackRaw != "" {
		rack, err := strconv.Atoi(rackRaw)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.NewMessageResponse("invalid 'rack' param"))
			return
		}

		books, err := b.service.GetBooksByRack(ctx, rack, offset, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
			return
		}

		booksDto := convertors.DomainBooksToDto(books)

		c.JSON(
			http.StatusOK,
			responses.NewBookPaginationResponse(booksDto).
				WithPage(page).
				WithLimit(limit).
				WithRack(rack),
		)
		return
	}

	count, err := b.cache.GetInt(constants.RedisTotalBooks)
	if errors.Is(err, redis.Nil) {
		count, err = b.service.GetBooksTotalCount(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
			return
		}
		b.cache.SetInt(constants.RedisTotalBooks, count)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	if offset > count {
		c.JSON(http.StatusUnprocessableEntity, responses.NewMessageResponse("'page' value too large"))
		return
	}

	books, err := b.service.GetBooks(ctx, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	totalPages := (count + limit - 1) / limit

	booksDto := convertors.DomainBooksToDto(books)

	c.JSON(
		http.StatusOK,
		responses.NewBookPaginationResponse(booksDto).
			WithPage(page).
			WithTotalPages(totalPages).
			WithLimit(limit),
	)
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

	if updateBook.Authors == nil && updateBook.Description == nil &&
		len(updateBook.CoverUrls) == 0 && updateBook.Title == nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("no parameters provided"))
		return
	}

	updateBookDomain := convertors.DtoUpdateBookInfoToDomain(updateBook)

	if err = b.service.UpdateBookInfo(c.Request.Context(), id, updateBookDomain); err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
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
// @Param book body dto.BookPlacement true "Book Placement"
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

	placement := dto.BookPlacement{}

	if err := c.ShouldBindJSON(&placement); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	err = b.service.UpdateBookPlacement(c.Request.Context(), id, placement.Rack, placement.Shelf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
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
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("no 'id' provided"))
		return
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	if id < 1 {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("id value is less than 1"))
		return
	}

	if err = b.service.DeleteBook(c.Request.Context(), id); err != nil {
		switch err {
		case errs.ErrNoRowsAffected:
			c.JSON(http.StatusNotFound, responses.NewMessageResponse("book not found"))
		default:
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		}
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
}
