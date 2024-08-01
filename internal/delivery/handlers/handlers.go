package handlers

import (
	"github.com/gin-gonic/gin"
)

type BookHandlers interface {
	GetBookByISBN(c *gin.Context)
	CreateBook(c *gin.Context)
	UpdateBookInfo(c *gin.Context)
	UpdateBookPlacement(c *gin.Context)
	DeleteBook(c *gin.Context)

	GetBooks(c *gin.Context)
	GetBooksByRack(c *gin.Context)
	GetBooksByTextSearch(c *gin.Context)
}
