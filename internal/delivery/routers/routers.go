package routers

import (
	"yurii-lib/internal/delivery/middleware"
	"yurii-lib/pkg/lgr"
	"yurii-lib/pkg/utils/jwt"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

func InitRouting(
	engine *gin.Engine,
	db *sqlx.DB,
	cache *redis.Client,
	storage *s3.S3,
	logger *lgr.Log,
	jwt jwt.JWT,
	middleware middleware.Middleware,
) {
	booksRG := engine.Group("/books")
	publicRG := engine.Group("/public")
	userRG := engine.Group("/user")
	adminUserRG := engine.Group("/admin")

	InitBooksRouter(booksRG, db, cache, storage, logger, middleware)
	InitPublicRouter(publicRG, db, logger, jwt)
	InitUserRouter(userRG, db, logger)
	InitAdminUsersRouter(adminUserRG, db, logger)
}
