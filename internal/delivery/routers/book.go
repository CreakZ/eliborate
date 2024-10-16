package routers

import (
	"yurii-lib/internal/delivery/handlers"
	"yurii-lib/internal/delivery/middleware"
	"yurii-lib/internal/repository"
	"yurii-lib/internal/service"
	"yurii-lib/pkg/lgr"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func InitBooksRouter(
	group *gin.RouterGroup,
	db *sqlx.DB,
	cache *redis.Client,
	storage *s3.S3,
	log *lgr.Log,
	middleware middleware.Middleware,
) {
	bookRepo := repository.InitBookRepo(db, storage)
	bookService := service.InitBookService(bookRepo, log)
	bookHandlers := handlers.InitBookHandlers(bookService, cache)

	group.POST("", bookHandlers.CreateBook)

	group.GET("/:isbn", bookHandlers.GetBookByISBN)

	group.PUT("/update/info", bookHandlers.UpdateBookInfo)
	group.PUT("/update/placement", bookHandlers.UpdateBookPlacement)

	group.DELETE("", bookHandlers.DeleteBook)

	group.POST("/racks", bookHandlers.GetBooksByRack)
	group.GET("", bookHandlers.GetBooks)
	group.GET("/search", bookHandlers.GetBooksByTextSearch)
}
