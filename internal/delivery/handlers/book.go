package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"yurii-lib/internal/format"
	"yurii-lib/internal/models/dto"
	"yurii-lib/internal/requests"
	"yurii-lib/internal/service"
	"yurii-lib/internal/validators"
	"yurii-lib/pkg/errs"

	"github.com/gin-gonic/gin"
)

type bookHandlers struct {
	service service.BookService
}

func InitBookHandlers(service service.BookService) BookHandlers {
	return bookHandlers{
		service: service,
	}
}

// Book handlers
func (b bookHandlers) GetBookByISBN(c *gin.Context) {
	isbnRaw, exists := c.GetQuery("isbn")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no isbn"})
		return
	}

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

func (b bookHandlers) CreateBook(c *gin.Context) {
	var book dto.BookPlacement

	body, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(body, &book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "json unmarshal error"})
		return
	}

	valid, err := validators.ValidBookPlacement(&book)
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
		book.Book.IsForeign == nil && book.Book.Logo == nil && book.Book.Title == nil {
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

// Library handlers
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

func (b bookHandlers) GetBooks(c *gin.Context) {
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

	// this condition limits page size by 4 values
	if !(limit == 10 || limit == 20 || limit == 50 || limit == 100) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("limit number: expected 10, 20, 50, 100, got %d", limit),
		})
		return
	}

	books, err := b.service.GetBooks(c.Request.Context(), page, limit)
	if err != nil {
		// handle specific errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (b bookHandlers) GetBooksByTextSearch(c *gin.Context) {
	text, exists := c.GetQuery("text")
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no query param"})
		return
	}

	// text validation

	books, err := b.service.GetBooksByTextSearch(c.Request.Context(), text)
	if err != nil {
		// switch some errors
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
