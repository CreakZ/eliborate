package handlers

import (
	"eliborate/internal/constants"
	"eliborate/internal/convertors"
	"eliborate/internal/delivery/responses"
	"eliborate/internal/errs"
	"eliborate/internal/models/dto"
	"eliborate/internal/service"
	"eliborate/internal/validators"
	"eliborate/pkg/storage"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	result := validators.ValidateBookCreate(&book)
	if !result.Ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(result.Err))
		return
	}

	bookDomain := convertors.DtoBookCreateToDomain(book)

	id, err := b.service.CreateBook(c.Request.Context(), bookDomain)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	// replace all c.Request.Context() with more suitable context
	_, err = b.cache.GetInt(constants.RedisTotalBooks)
	if errors.Is(err, redis.Nil) {
		totalBooks, err := b.service.GetBooksTotalCount(c.Request.Context())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
			return
		}
		b.cache.SetInt(constants.RedisTotalBooks, totalBooks)
	} else {
		b.cache.Incr(constants.RedisTotalBooks)
	}

	c.JSON(http.StatusCreated, responses.NewMessageResponse(fmt.Sprintf("book with id %v created successfully", id)))
}

// @Summary Get book by id
// @Description Retrieve book information by its id
// @Tags books
// @Param id path int true "Rack Number"
// @Produce json
// @Success 200 {array}  dto.Book
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books/{id} [get]
func (b bookHandlers) GetBookById(c *gin.Context) {
	idRaw := c.Param("id")
	if idRaw == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponse("no id number provided"))
		return
	}

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	book, err := b.service.GetBookById(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	c.JSON(http.StatusOK, book)
}

// @Summary Get a paginated list of books
// @Description Retrieve a list of books with pagination
// @Tags books
// @Param page query int true "Page number"
// @Param limit query int true "Books limit per page (10, 20, 50 or 100)"
// @Produce json
// @Success 200 {object} responses.BookPaginationResponse "page, limit, results (array of dto.Book), total_pages"
// @Failure 400 {object} responses.MessageResponse
// @Failure 404 {object} responses.MessageResponse
// @Failure 422 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books [get]
func (b bookHandlers) GetBooks(c *gin.Context) {
	var count int

	count, err := b.cache.GetInt(constants.RedisTotalBooks)
	if errors.Is(err, redis.Nil) {
		count, err := b.service.GetBooksTotalCount(c.Request.Context())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
			return
		}
		b.cache.SetInt(constants.RedisTotalBooks, count)
	}

	pageRaw, exists := c.GetQuery("page")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponse("no page number provided"))
		return
	}

	limitRaw, exists := c.GetQuery("limit")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponse("no limit number provided"))
		return
	}

	page, err := strconv.Atoi(pageRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	if page < 1 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, responses.NewMessageResponse("'page' shouldn't be less than 1"))
		return
	}

	limit, err := strconv.Atoi(limitRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, responses.NewMessageResponseFromErr(err))
		return
	}

	// this condition limits page size by 4 values
	if !(limit == 10 || limit == 20 || limit == 50 || limit == 100) {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			responses.NewMessageResponse(fmt.Sprintf("limit number: expected 10, 20, 50, 100, got %d", limit)),
		)
		return
	}

	if (page-1)*limit > count {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, responses.NewMessageResponse("'page' value too large"))
		return
	}

	books, err := b.service.GetBooks(c.Request.Context(), page, limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	totalPages := count/limit + 1

	booksDto := make([]dto.Book, 0, len(books))
	for _, book := range books {
		booksDto = append(booksDto, convertors.DomainBookToDto(book))
	}

	c.JSON(http.StatusOK, responses.NewBookPaginationResponse(page, totalPages, limit, booksDto))
}

// @Summary Get books by rack number
// @Description Retrieve all books located in a specific rack
// @Tags books
// @Param rack path int true "Rack Number"
// @Produce json
// @Success 200 {array}  dto.Book
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books/racks/{rack} [get]
func (b bookHandlers) GetBooksByRack(c *gin.Context) {
	rawRack := c.Param("rack")
	if rawRack == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no rack number provided"})
		return
	}

	rack, err := strconv.Atoi(rawRack)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	books, err := b.service.GetBooksByRack(c.Request.Context(), rack)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// @Summary Search books by text
// @Description Search for books matching a text query
// @Tags books
// @Param text query string true "Search Query"
// @Produce json
// @Success 200 {array}  responses.BookSearchResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /books/search [get]
func (b bookHandlers) GetBooksByTextSearch(c *gin.Context) {
	text, exists := c.GetQuery("text")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no 'text' query param provided"})
		return
	}

	if err := validators.ValidateTextQuery(text); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	books, err := b.service.GetBooksByTextSearch(c.Request.Context(), text)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	booksDto := make([]dto.BookSearch, 0, len(books))
	for _, book := range books {
		booksDto = append(booksDto, convertors.DomainBookSearchToDto(book))
	}

	c.JSON(http.StatusOK, responses.NewBookSearchResponse(booksDto))
}

// @Summary Update book information
// @Description Update details of an existing book
// @Tags books
// @Accept json
// @Produce json
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
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	var updateBook dto.UpdateBookInfo

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	if err = json.Unmarshal(body, &updateBook); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	if updateBook.Authors == nil && updateBook.Category == nil && updateBook.Description == nil &&
		len(updateBook.CoverUrls) == 0 && updateBook.Title == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponse("no parameters provided"))
		return
	}

	updateBookDomain := convertors.DtoUpdateBookInfoToDomain(updateBook)

	if err = b.service.UpdateBookInfo(c.Request.Context(), id, updateBookDomain); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	c.JSON(http.StatusOK, responses.NewMessageResponse(fmt.Sprintf("info about book with id %v updated successfully", id)))
}

// @Summary Update book placement
// @Description Update the rack and shelf placement of a book
// @Tags books
// @Accept json
// @Produce json
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
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	placement := struct {
		Rack  int `json:"rack"`
		Shelf int `json:"shelf"`
	}{
		Rack:  0,
		Shelf: 0,
	}

	if err := c.ShouldBindJSON(&placement); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	err = b.service.UpdateBookPlacement(c.Request.Context(), id, placement.Rack, placement.Shelf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	c.JSON(
		http.StatusOK,
		responses.NewMessageResponse(
			fmt.Sprintf("placement of book with id %v updated successfully", id),
		),
	)
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
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponse("no 'id' provided"))
		return
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	if id < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id value is less than 1"})
		return
	}

	if err = b.service.DeleteBook(c.Request.Context(), id); err != nil {
		switch err {
		case errs.ErrNoRowsAffected:
			c.AbortWithStatus(http.StatusNotFound)
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("book with id %d deleted successfully", id)})
}
