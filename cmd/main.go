// @title Eliborate API Documentation
// @version 1.0
// @description Swagger OpenAPI documentation for Eliborate service

// @contact.name Maxim Rusakov
// @contact.email shejustwannagethigh@yandex.ru

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"eliborate/internal/delivery/middleware"
	"eliborate/internal/delivery/routers"
	"eliborate/pkg/config"
	"eliborate/pkg/logging"
	"eliborate/pkg/storage"
	"eliborate/pkg/utils"
	"fmt"

	_ "eliborate/docs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	server := gin.Default()

	server.StaticFile("/docs", "./docs/swagger.json")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Init logger
	logger, infoFile, errorFile := logging.InitLogger()
	defer infoFile.Close()
	defer errorFile.Close()
	logger.InfoLogger.Info().Msg("Logger initialized successfully")

	// Init env configuration
	config.InitConfig()
	logger.InfoLogger.Info().Msg("Config initialized successfully")

	// Init cors configuration
	corsCfg := config.NewCorsConfig(
		viper.GetString(config.AccessControlAllowOrigin),
		viper.GetString(config.AccessControlAllowMethods),
		viper.GetString(config.AccessControlAllowHeaders),
	)

	// Init db conn
	db, err := storage.NewPostgresConn(
		viper.GetString(config.PostgresUser),
		viper.GetString(config.PostgresPassword),
		viper.GetString(config.PostgresDBName),
		viper.GetString(config.PostgresHost),
		viper.GetInt(config.PostgresPort),
	)
	if err != nil {
		panic(err)
	}
	logger.InfoLogger.Info().Msg("Postgres initialized successfully")

	// Init redis client
	cache := storage.NewRedisCacheManager()

	// Init jwt utils
	jwtUtil := utils.InitJWTUtil()
	logger.InfoLogger.Info().Msg("JWT initialized successfully")

	// Init middleware
	middleW := middleware.InitMiddleware(jwtUtil, logger, corsCfg)
	logger.InfoLogger.Info().Msg("Middleware initialized successfully")

	// Use CORS middleware
	server.Use(middleW.CorsMiddleware())

	// Init Meilisearch client
	search, err := storage.NewMeiliClient(
		viper.GetString(config.MeiliHost),
		viper.GetInt(config.MeiliPort),
		viper.GetString(config.MeiliIndex),
		viper.GetString(config.MeiliMasterKey),
	)
	if err != nil {
		panic(err)
	}
	logger.InfoLogger.Info().Msg("Search engine initialized successfully")

	// Init routing
	routers.InitRouting(server, db, cache, logger, jwtUtil, middleW, search)

	// Server startup
	if err := server.Run(":8080"); err != nil {
		panic(fmt.Errorf("server run error: %s", err))
	}
}
