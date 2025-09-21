package routers

import (
	"eliborate/internal/delivery/middleware"
	"eliborate/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/meilisearch/meilisearch-go"
)

func InitRouting(
	engine *gin.Engine,
	db *sqlx.DB,
	jwt utils.JWT,
	middleware middleware.Middleware,
	search meilisearch.IndexManager,
) {
	booksRG := engine.Group("/books")
	publicRG := engine.Group("/auth")
	userRG := engine.Group("/users")
	adminUserRG := engine.Group("/admins")
	catRG := engine.Group("/categories")

	InitBooksRouter(booksRG, db, middleware, search)
	InitPublicRouter(publicRG, db, jwt)
	InitUserRouter(userRG, db, middleware)
	InitAdminUsersRouter(adminUserRG, db, middleware)
	InitCategoryRouter(catRG, db, middleware)
}
