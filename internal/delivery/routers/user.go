package routers

import (
	"yurii-lib/internal/delivery/handlers"
	"yurii-lib/internal/repository"
	"yurii-lib/internal/service"
	"yurii-lib/pkg/lgr"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitUserRouter(group *gin.RouterGroup, db *sqlx.DB, logger *lgr.Log) {
	userRepo := repository.InitUserRepo(db)
	userService := service.InitUserService(userRepo, logger)
	userHandlers := handlers.InitUserHandlers(userService)

	group.POST("", userHandlers.Create)

	group.GET("", userHandlers.GetPassword)

	group.PUT("", userHandlers.UpdatePassword)

	group.DELETE("", userHandlers.Delete)
}
