package middleware

import (
	"yurii-lib/pkg/log"
	"yurii-lib/pkg/utils/jwt"
)

type Middleware struct {
	jwtUtil jwt.JWT
	logger  *log.Log
}

func InitMiddleware(util jwt.JWT, logger *log.Log) Middleware {
	return Middleware{
		jwtUtil: util,
		logger:  logger,
	}
}
