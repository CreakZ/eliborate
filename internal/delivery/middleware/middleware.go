package middleware

import (
	"eliborate/pkg/config"
	"eliborate/pkg/utils"
)

type Middleware struct {
	jwtUtil utils.JWT
	corsCfg *config.CorsConfig
}

func InitMiddleware(util utils.JWT, corsCfg *config.CorsConfig) Middleware {
	return Middleware{
		jwtUtil: util,
		corsCfg: corsCfg,
	}
}
