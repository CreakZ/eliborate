package routers

import (
	"yurii-lib/internal/delivery/handlers"
	"yurii-lib/internal/delivery/middleware"
	"yurii-lib/internal/repository"
	"yurii-lib/internal/service"
	"yurii-lib/pkg/log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

func InitBooksRouter(group *gin.RouterGroup, db *sqlx.DB, cache *redis.Client, storage *s3.S3, log *log.Log, middleware middleware.Middleware) {
	bookRepo := repository.InitBookRepo(db, storage)
	bookService := service.InitBookService(bookRepo, log)
	bookHandlers := handlers.InitBookHandlers(bookService, cache)

	bookRouter := group.Group("/book")

	bookRouter.POST("/create", bookHandlers.CreateBook)

	bookRouter.GET("/get", bookHandlers.GetBookByISBN)

	bookRouter.PUT("/update/info", bookHandlers.UpdateBookInfo)
	bookRouter.PUT("/update/placement", bookHandlers.UpdateBookPlacement)

	bookRouter.DELETE("/delete", bookHandlers.DeleteBook)

	libraryRouter := group.Group("/library")

	libraryRouter.POST("/racks", bookHandlers.GetBooksByRack)

	libraryRouter.GET("/books", bookHandlers.GetBooks)
	libraryRouter.GET("/search", bookHandlers.GetBooksByTextSearch)
}
