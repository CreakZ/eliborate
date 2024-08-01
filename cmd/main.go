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

	db := database.ConnectDB()

	// Инициализация JWT
	jwtUtil := jwt.InitJWTUtil()

	// Инициализация middleware
	middlewarrior := middleware.InitMiddleware(jwtUtil, logger)

	// Инициализация маршрутизаторов
	routers.InitRouting(server, db, logger, middlewarrior)

	// Запуска сервера
	if err := server.Run(":8080"); err != nil {
		panic(fmt.Errorf("server run error: %s", err))
	}
}
