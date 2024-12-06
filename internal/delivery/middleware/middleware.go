package middleware

import (
	"eliborate/pkg/logging"
	"eliborate/pkg/utils"
)

type Middleware struct {
	jwtUtil utils.JWT
	logger  *logging.Log
}

func InitMiddleware(util utils.JWT, logger *logging.Log) Middleware {
	return Middleware{
		jwtUtil: util,
		logger:  logger,
	}
}
