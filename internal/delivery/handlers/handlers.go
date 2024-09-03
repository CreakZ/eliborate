package handlers

import (
	"github.com/gin-gonic/gin"
)

type BookHandlers interface {
	CreateBook(c *gin.Context)

	GetBookByISBN(c *gin.Context)

	GetBooks(c *gin.Context)
	GetBooksByRack(c *gin.Context)
	GetBooksByTextSearch(c *gin.Context)

	UpdateBookInfo(c *gin.Context)
	UpdateBookPlacement(c *gin.Context)

	DeleteBook(c *gin.Context)
}

type UserHandlers interface {
	Create(c *gin.Context)
}

type AdminUserHandlers interface {
	CreateAll(c *gin.Context)
	Create(c *gin.Context)

	Update(c *gin.Context)
}
