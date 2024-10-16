package middleware

import (
	"yurii-lib/pkg/lgr"
	"yurii-lib/pkg/utils/jwt"
)

type Middleware struct {
	jwtUtil jwt.JWT
	logger  *lgr.Log
}

func InitMiddleware(util jwt.JWT, logger *lgr.Log) Middleware {
	return Middleware{
		jwtUtil: util,
		logger:  logger,
	}
}
