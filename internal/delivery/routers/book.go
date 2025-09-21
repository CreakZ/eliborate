package routers

import (
	"eliborate/internal/delivery/handlers"
	"eliborate/internal/delivery/middleware"
	"eliborate/internal/repository"
	"eliborate/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
)

func InitBooksRouter(
	rg *gin.RouterGroup,
	db *sqlx.DB,
	middleware middleware.Middleware,
	search meilisearch.IndexManager,
) {
	bookRepo := repository.InitBookRepo(db, search)
	bookService := service.InitBookService(bookRepo)
	bookHandlers := handlers.InitBookHandlers(bookService)

	rg.POST("", middleware.BearerAuthMiddleware(), bookHandlers.CreateBook)

	rg.GET("/:id", bookHandlers.GetBookById)
	rg.GET("", bookHandlers.GetBooks)

	rg.PATCH("/:id/info", middleware.BearerAuthMiddleware(), bookHandlers.UpdateBookInfo)
	rg.PATCH("/:id/placement", middleware.BearerAuthMiddleware(), bookHandlers.UpdateBookPlacement)

	rg.DELETE("/:id", middleware.BearerAuthMiddleware(), bookHandlers.DeleteBook)
}
