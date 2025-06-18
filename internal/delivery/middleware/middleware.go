package middleware

import (
	"eliborate/pkg/config"
	"eliborate/pkg/logging"
	"eliborate/pkg/utils"
)

type Middleware struct {
	jwtUtil utils.JWT
	logger  *logging.Log
	corsCfg *config.CorsConfig
}

func InitMiddleware(util utils.JWT, logger *logging.Log, corsCfg *config.CorsConfig) Middleware {
	return Middleware{
		jwtUtil: util,
		logger:  logger,
		corsCfg: corsCfg,
	}
}
