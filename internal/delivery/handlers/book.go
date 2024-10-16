package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"yurii-lib/internal/models/dto"
	"yurii-lib/internal/requests"
	"yurii-lib/internal/service"
	"yurii-lib/internal/validators"
	"yurii-lib/pkg/errs"
	"yurii-lib/pkg/utils/format"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

const TotalBooks = "total_books"

type bookHandlers struct {
	service service.BookService
	cache   *redis.Client
}

func InitBookHandlers(service service.BookService, cache *redis.Client) BookHandlers {
	return bookHandlers{
		service: service,
		cache:   cache,
	}
}

// @Summary Get a book by its ISBN
// @Description Retrieve a book's data from online libraries and book APIs using its ISBN number
// @Tags books
// @Param isbn query string true "ISBN"
// @Produce json
// @Success 200 {object} dto.Book
// @Failure 400 {object} map[string]string
// @Failure 204
// @Failure 500 {object} map[string]string
// @Router /books/{isbn} [get]
func (b bookHandlers) GetBookByISBN(c *gin.Context) {
	isbnRaw := c.Param("isbn")

	isbn, err := format.FormatISBN(isbnRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	books, err := requests.GetBookByISBN(isbn)
	if err != nil {
		switch err {
		case requests.ErrNoBooksFound:
			c.AbortWithStatus(http.StatusNoContent)
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, books)
}

// @Summary Create a new book
// @Description Create a new book entry in the system
// @Tags books
// @Accept json
// @Produce json
// @Param book body dto.BookPlacement true "Book Placement"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [post]
func (b bookHandlers) CreateBook(c *gin.Context) {
	var book dto.BookPlacement

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	valid, err := validators.ValidateBookPlacement(&book)
	if !valid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := b.service.CreateBook(c.Request.Context(), book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("book with id %v created successfully", id)})
}

// @Summary Update book information
// @Description Update details of an existing book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/update/info [put]
func (b bookHandlers) UpdateBookInfo(c *gin.Context) {
	book := struct {
		ID   int `json:"id"`
		Book dto.UpdateBookInfo
	}{
		ID:   0,
		Book: dto.UpdateBookInfo{},
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = json.Unmarshal(body, &book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if book.Book.Authors == nil && book.Book.Category == nil && book.Book.Description == nil &&
		book.Book.IsForeign == nil && book.Book.CoverURL == nil && book.Book.Title == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no parameters provided"})
		return
	}

	if err = b.service.UpdateBookInfo(c.Request.Context(), book.ID, book.Book); err != nil {
		switch err {
		case errs.ErrNoRowsAffected:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"id": book.ID, "message": err.Error()})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"id": book.ID, "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("info about book with id %v updated successfully", book.ID)})
}

// @Summary Update book placement
// @Description Update the rack and shelf placement of a book
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/update/placement [put]
func (b bookHandlers) UpdateBookPlacement(c *gin.Context) {
	placement := struct {
		ID    int `json:"id"`
		Rack  int `json:"rack"`
		Shelf int `json:"shelf"`
	}{
		ID:    0,
		Rack:  0,
		Shelf: 0,
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = json.Unmarshal(body, &placement); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = b.service.UpdateBookPlacement(c.Request.Context(), placement.ID, placement.Rack, placement.Shelf)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": placement.ID,
		"message": fmt.Sprintf("placement of book with id %v updated successfully", placement.ID)})
}

// @Summary Delete a book by its ID
// @Description Remove a book from the system using its ID
// @Tags books
// @Param id query int true "Book ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [delete]
func (b bookHandlers) DeleteBook(c *gin.Context) {
	rawID, exists := c.GetQuery("id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no id provided"})
		return
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if id < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id value is less than 1"})
		return
	}

	if err = b.service.DeleteBook(c.Request.Context(), id); err != nil {
		switch err {
		case errs.ErrNoRowsAffected:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("book with id %d does not exist", id)})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("book with id %v deleted successfully", id)})
}

// @Summary Get books by rack number
// @Description Retrieve all books located in a specific rack
// @Tags books
// @Param rack query int true "Rack Number"
// @Produce json
// @Success 200 {array} dto.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/racks [get]
func (b bookHandlers) GetBooksByRack(c *gin.Context) {
	rawRack, exists := c.GetQuery("rack")
	if !exists {
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

// @Summary Get a paginated list of books
// @Description Retrieve a list of books with pagination
// @Tags books
// @Param page query int true "Page number"
// @Param limit query int true "Limit per page"
// @Produce json
// @Success 200 {object} map[string]interface{} "page, limit, results (array of dto.Book), total_pages"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books [get]
func (b bookHandlers) GetBooks(c *gin.Context) {
	var count int

	// Проверка, существует ли значение общего количества книг в кэше
	stringCmd := b.cache.Get(c.Request.Context(), TotalBooks)
	if stringCmd.Err() != nil {
		count, err := b.service.GetBooksTotalCount(c.Request.Context())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if err = b.cache.Set(c.Request.Context(), TotalBooks, count, 0).Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	} else {
		count, _ = stringCmd.Int()
	}

	pageRaw, exists := c.GetQuery("page")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no page number provided"})
		return
	}

	limitRaw, exists := c.GetQuery("limit")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no limit number provided"})
		return
	}

	page, err := strconv.Atoi(pageRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	limit, err := strconv.Atoi(limitRaw)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if (page+1)*limit > count {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "wrong page or limit value"})
		return
	}

	// this condition limits page size by 4 values
	if !(limit == 10 || limit == 20 || limit == 50 || limit == 100) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("limit number: expected 10, 20, 50, 100, got %d", limit),
		})
		return
	}

	books, err := b.service.GetBooks(c.Request.Context(), page, limit)
	if err != nil {
		// обработка специфических ошибок
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var totalPages int = count / limit
	if count%limit != 0 {
		totalPages++
	}

	var body = struct {
		Page       int        `json:"page"`
		Limit      int        `json:"limit"`
		Results    []dto.Book `json:"results"`
		TotalPages int        `json:"total_pages"`
	}{
		Page:       page,
		Limit:      limit,
		Results:    books,
		TotalPages: totalPages,
	}

	c.JSON(http.StatusOK, body)
}

// @Summary Search books by text
// @Description Search for books matching a text query
// @Tags books
// @Param text query string true "Search Query"
// @Produce json
// @Success 200 {array} dto.Book
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /books/search [get]
func (b bookHandlers) GetBooksByTextSearch(c *gin.Context) {
	text, exists := c.GetQuery("text")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no query param"})
		return
	}

	if err := validators.ValidateTextQuery(text); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	books, err := b.service.GetBooksByTextSearch(c.Request.Context(), text)
	if err != nil {
		// switch some errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
