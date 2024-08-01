package routers

import (
	"yurii-lib/internal/delivery/handlers"
	"yurii-lib/internal/delivery/middleware"
	"yurii-lib/internal/repository"
	"yurii-lib/internal/service"
	"yurii-lib/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitBooksRouter(group *gin.RouterGroup, db *sqlx.DB, log *log.Log, middleware middleware.Middleware) {
	bookRepo := repository.InitBookRepo(db)
	bookService := service.InitBookService(bookRepo, log)
	bookHandlers := handlers.InitBookHandlers(bookService)

	bookRouter := group.Group("/book")

	bookRouter.POST("/create", bookHandlers.CreateBook)

	bookRouter.GET("/get", bookHandlers.GetBookByISBN)

	/*
		UpdateBookInfo has problems:
		- Passing struct as an argument implies request to update ALL row fields, including those weren't indended to update
	*/

	bookRouter.PUT("/update/info", bookHandlers.UpdateBookInfo)
	bookRouter.PUT("/update/placement", bookHandlers.UpdateBookPlacement)

	bookRouter.DELETE("/delete", bookHandlers.DeleteBook)

	libraryRouter := group.Group("/library")

	libraryRouter.POST("/racks", bookHandlers.GetBooksByRack)

	libraryRouter.GET("/books", bookHandlers.GetBooks)
	libraryRouter.GET("/search", bookHandlers.GetBooksByTextSearch)
}
