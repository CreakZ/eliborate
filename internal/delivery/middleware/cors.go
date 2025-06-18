package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	AccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	AccessControlAllowMethods = "Access-Control-Allow-Methods"
	AccessControlAllowHeaders = "Access-Control-Allow-Headers"
)

func (m Middleware) CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(AccessControlAllowOrigin, "http://localhost:5173")
		c.Writer.Header().Set(AccessControlAllowMethods, "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set(AccessControlAllowHeaders, "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
