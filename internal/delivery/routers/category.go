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

	rg.POST("", middleware.Authorize(), catHandlers.Create)
	rg.GET("", middleware.Authorize(), catHandlers.GetAll)
	rg.PATCH("", middleware.Authorize(), catHandlers.Update)
	rg.DELETE("", middleware.Authorize(), catHandlers.Delete)
}
