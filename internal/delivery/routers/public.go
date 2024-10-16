package routers

import (
	"yurii-lib/internal/delivery/handlers"
	"yurii-lib/internal/repository"
	"yurii-lib/internal/service"
	"yurii-lib/pkg/lgr"
	"yurii-lib/pkg/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitPublicRouter(group *gin.RouterGroup, db *sqlx.DB, logger *lgr.Log, jwt jwt.JWT) {
	publicRepo := repository.InitPublicRepo(db)
	publicService := service.InitPublicService(publicRepo, logger)
	publicHandlers := handlers.InitPublicHandlers(publicService, jwt)

	group.POST("/admin", publicHandlers.LoginAdminUser)
	group.POST("/user", publicHandlers.LoginUser)
}
