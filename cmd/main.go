// @title API Documentation
// @version 1.0
// @description This is a sample API documentation for your project
// @termsOfService http://example.com/terms/

// @contact.name Maxim
// @contact.email shejustwannagethigh@yandex.ru

// @host localhost:8080
// @BasePath /

package main

import (
	"fmt"
	"yurii-lib/internal/delivery/middleware"
	"yurii-lib/internal/delivery/routers"
	"yurii-lib/pkg/config"
	"yurii-lib/pkg/database"
	"yurii-lib/pkg/lgr"
	"yurii-lib/pkg/utils/jwt"

	_ "yurii-lib/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	server := gin.Default()

	server.StaticFile("/swagger", "./docs/swagger.json")

	// Инициализация логгера
	logger, infoFile, errorFile := lgr.InitLogger()
	defer infoFile.Close()
	defer errorFile.Close()
	logger.InfoLogger.Info().Msg("Logger initialized successfully")

	// Инициализация конфига
	config.InitConfig()
	logger.InfoLogger.Info().Msg("Config initialized successfully")

	// Инициализация подключения к БД
	db := database.ConnectDB()
	logger.InfoLogger.Info().Msg("Database initialized successfully")

	// Инициализация клиента S3-хранилища
	// svc := database.InitS3Client()
	// logger.InfoLogger.Info().Msg("S3 initialized successfully")

	// Инициализация клиента Redis
	cache := database.InitCache()

	// Инициализация JWT
	jwtUtil := jwt.InitJWTUtil()
	logger.InfoLogger.Info().Msg("JWT initialized successfully")

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Инициализация middleware
	middlewarrior := middleware.InitMiddleware(jwtUtil, logger)
	logger.InfoLogger.Info().Msg("Middleware initialized successfully")

	// Инициализация маршрутизаторов
	routers.InitRouting(server, db, cache, nil, logger, jwtUtil, middlewarrior)

	// Запуск сервера
	if err := server.Run(":8080"); err != nil {
		panic(fmt.Errorf("server run error: %s", err))
	}
}
