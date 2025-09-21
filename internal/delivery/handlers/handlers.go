package handlers

import (
	"github.com/gin-gonic/gin"
)

type BookHandlers interface {
	CreateBook(c *gin.Context)

	GetBookById(c *gin.Context)
	GetBooks(c *gin.Context)

	UpdateBookInfo(c *gin.Context)
	UpdateBookPlacement(c *gin.Context)

	DeleteBook(c *gin.Context)
}

type CategoryHandlers interface {
	Create(c *gin.Context)

	GetAll(c *gin.Context)

	Update(c *gin.Context)

	Delete(c *gin.Context)
}

type UserHandlers interface {
	Create(c *gin.Context)

	// GetPassword(c *gin.Context)

	UpdatePassword(c *gin.Context)

	Delete(c *gin.Context)
}

type AdminUserHandlers interface {
	// GetPassword(c *gin.Context)

	UpdatePassword(c *gin.Context)
}

type PublicHandlers interface {
	LoginAdminUser(c *gin.Context)
	LoginUser(c *gin.Context)
}
