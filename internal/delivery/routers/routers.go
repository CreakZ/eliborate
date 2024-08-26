package routers

import (
	"yurii-lib/internal/delivery/middleware"
	"yurii-lib/pkg/log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRouting(engine *gin.Engine, db *sqlx.DB, storage *s3.S3, log *log.Log, middleware middleware.Middleware) {
	bookRouter := engine.Group("")

	InitBooksRouter(bookRouter, db, storage, log, middleware)
}
