package routers

import (
	"eliborate/internal/delivery/handlers"
	"eliborate/internal/delivery/middleware"
	"eliborate/internal/repository"
	"eliborate/internal/service"
	"eliborate/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitCategoryRouter(rg *gin.RouterGroup, db *sqlx.DB, middleware middleware.Middleware, logger *logging.Log) {
	catRepo := repository.InitCategoryRepo(db)
	catService := service.NewCategoryService(catRepo)
	catHandlers := handlers.NewCategoryHandlers(catService)

	rg.POST("", middleware.BearerAuthMiddleware(), catHandlers.Create)
	rg.GET("", middleware.BearerAuthMiddleware(), catHandlers.GetAll)
	rg.PATCH("/:id", middleware.BearerAuthMiddleware(), catHandlers.Update)
	rg.DELETE("/:id", middleware.BearerAuthMiddleware(), catHandlers.Delete)
}
