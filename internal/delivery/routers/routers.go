package routers

import (
	"yurii-lib/internal/delivery/middleware"
	"yurii-lib/pkg/log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouting(engine *gin.Engine, db *sqlx.DB, log *log.Log, middleware middleware.Middleware) {
	bookRouter := engine.Group("")

	InitBooksRouter(bookRouter, db, log, middleware)
}
