package routers

import (
	"eliborate/internal/delivery/middleware"
	"eliborate/pkg/logging"
	"eliborate/pkg/storage"
	"eliborate/pkg/utils"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
)

func InitRouting(
	engine *gin.Engine,
	db *sqlx.DB,
	cache *storage.RedisCache,
	storage *s3.S3,
	logger *logging.Log,
	jwt utils.JWT,
	middleware middleware.Middleware,
	search meilisearch.IndexManager,
) {
	booksRG := engine.Group("/books")
	publicRG := engine.Group("/public")
	userRG := engine.Group("/users")
	adminUserRG := engine.Group("/admins")
	catRG := engine.Group("/categories")

	InitBooksRouter(booksRG, db, cache, logger, middleware, search)
	InitPublicRouter(publicRG, db, logger, jwt)
	InitUserRouter(userRG, db, middleware, logger)
	InitAdminUsersRouter(adminUserRG, db, middleware, logger)
	InitCategoryRouter(catRG, db, middleware, logger)
}
