package routers

import (
	"yurii-lib/internal/delivery/handlers"
	"yurii-lib/internal/repository"
	"yurii-lib/internal/service"
	"yurii-lib/pkg/lgr"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitAdminUsersRouter(group *gin.RouterGroup, db *sqlx.DB, logger *lgr.Log) {
	adminRepo := repository.InitAdminUserRepo(db)
	adminService := service.InitAdminUserService(adminRepo, logger)
	adminHandlers := handlers.InitAdminUserHandlers(adminService)

	// group.POST("/create", adminHandlers.Create)
	// group.POST("/create_all", adminHandlers.CreateAll)

	group.GET("", adminHandlers.GetPassword)

	group.PUT("", adminHandlers.UpdatePassword)

	// group.DELETE("/delete", adminHandlers.Delete)
}
