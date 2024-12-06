package routers

import (
	"eliborate/internal/delivery/handlers"
	"eliborate/internal/repository"
	"eliborate/internal/service"
	"eliborate/pkg/logging"
	"eliborate/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitPublicRouter(group *gin.RouterGroup, db *sqlx.DB, logger *logging.Log, jwt utils.JWT) {
	publicRepo := repository.InitPublicRepo(db)
	publicService := service.InitPublicService(publicRepo, logger)
	publicHandlers := handlers.InitPublicHandlers(publicService, jwt)

	group.POST("/admin", publicHandlers.LoginAdminUser)
	group.POST("/user", publicHandlers.LoginUser)
}
