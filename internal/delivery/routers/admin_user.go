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

func InitAdminUsersRouter(group *gin.RouterGroup, db *sqlx.DB, middleware middleware.Middleware, logger *logging.Log) {
	adminRepo := repository.InitAdminUserRepo(db)
	adminService := service.InitAdminUserService(adminRepo, logger)
	adminHandlers := handlers.InitAdminUserHandlers(adminService)

	group.PATCH("", middleware.BearerAuthMiddleware(), adminHandlers.UpdatePassword)
}
