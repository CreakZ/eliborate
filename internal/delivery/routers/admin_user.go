package routers

import (
	"eliborate/internal/delivery/handlers"
	"eliborate/internal/delivery/middleware"
	"eliborate/internal/repository"
	"eliborate/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAdminUsersRouter(group *gin.RouterGroup, db *sqlx.DB, middleware middleware.Middleware) {
	adminRepo := repository.InitAdminUserRepo(db)
	adminService := service.InitAdminUserService(adminRepo)
	adminHandlers := handlers.InitAdminUserHandlers(adminService)

	group.PATCH("", middleware.BearerAuthMiddleware(), adminHandlers.UpdatePassword)
}
