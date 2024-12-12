package routers

import (
	"eliborate/internal/delivery/handlers"
	"eliborate/internal/delivery/middleware"
	"eliborate/internal/repository"
	"eliborate/internal/service"
	"eliborate/pkg/logging"
	"eliborate/pkg/storage"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
)

func InitBooksRouter(
	rg *gin.RouterGroup,
	db *sqlx.DB,
	cache *storage.RedisCache,
	log *logging.Log,
	middleware middleware.Middleware,
	search meilisearch.IndexManager,
) {
	bookRepo := repository.InitBookRepo(db, search)
	bookService := service.InitBookService(bookRepo, log)
	bookHandlers := handlers.InitBookHandlers(bookService, cache)

	rg.POST("", middleware.Authorize(), bookHandlers.CreateBook)

	rg.GET("/:id", bookHandlers.GetBookById)
	rg.GET("/isbn/:isbn", bookHandlers.GetBookByIsbn)
	rg.GET("", bookHandlers.GetBooks)
	rg.GET("/racks/:rack", bookHandlers.GetBooksByRack)
	rg.GET("/search", bookHandlers.GetBooksByTextSearch)

	rg.PATCH("/:id/info", middleware.Authorize(), bookHandlers.UpdateBookInfo)
	rg.PATCH("/:id/placement", middleware.Authorize(), bookHandlers.UpdateBookPlacement)

	rg.DELETE("/:id", middleware.Authorize(), bookHandlers.DeleteBook)
}
