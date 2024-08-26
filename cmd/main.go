package main

import (
	"fmt"
	"yurii-lib/internal/delivery/middleware"
	"yurii-lib/internal/delivery/routers"
	"yurii-lib/pkg/config"
	"yurii-lib/pkg/database"
	"yurii-lib/pkg/log"
	"yurii-lib/pkg/utils/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Инициализация логгера
	logger, infoFile, errorFile := log.InitLogger()
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
	svc := database.InitS3Client()
	logger.InfoLogger.Info().Msg("S3 initialized successfully")

	// Инициализация JWT
	jwtUtil := jwt.InitJWTUtil()
	logger.InfoLogger.Info().Msg("JWT initialized successfully")

	// Инициализация middleware
	middlewarrior := middleware.InitMiddleware(jwtUtil, logger)
	logger.InfoLogger.Info().Msg("Middleware initialized successfully")

	// Инициализация маршрутизаторов
	routers.InitRouting(server, db, svc, logger, middlewarrior)

	// Запуск сервера
	if err := server.Run(":8080"); err != nil {
		panic(fmt.Errorf("server run error: %s", err))
	}
}
