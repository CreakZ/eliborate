package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const userID = "user_id"

func (m Middleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")

		if !strings.Contains(auth, "Bearer") {
			m.logger.InfoLogger.Info().Msg("no authentication token provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no authentication token provided"})
			return
		}

		token := strings.Split(auth, " ")[1]

		claim, valid, err := m.jwtUtil.Authorize(token)
		if err != nil {
			m.logger.InfoLogger.Info().Msg(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		if !valid {
			m.logger.InfoLogger.Info().Msg("jwt is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "jwt is not valid"})
			return
		}

		c.Set(userID, claim.ID)
	}
}
