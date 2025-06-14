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

func InitUserRouter(group *gin.RouterGroup, db *sqlx.DB, middleware middleware.Middleware, logger *logging.Log) {
	userRepo := repository.InitUserRepo(db)
	userService := service.InitUserService(userRepo, logger)
	userHandlers := handlers.InitUserHandlers(userService)

	group.POST("", userHandlers.Create)

	group.PATCH("", middleware.BearerAuthMiddleware(), userHandlers.UpdatePassword)

	group.DELETE("", middleware.BearerAuthMiddleware(), userHandlers.Delete)
}
