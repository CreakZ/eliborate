package middleware

import (
	"eliborate/internal/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m Middleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			m.logger.InfoLogger.Info().Msg("no jwt provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "no jwt provided"})
			return
		}

		claim, valid, err := m.jwtUtil.Authorize(token)
		if err != nil {
			m.logger.InfoLogger.Info().Msg(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if !valid {
			m.logger.InfoLogger.Info().Msg("jwt is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "jwt is not valid"})
			return
		}

		userRole := constants.RoleClient
		if claim.IsAdmin {
			userRole = constants.RoleAdmin
		}

		c.Set(constants.KeyUserID, claim.ID)
		c.Set(constants.KeyRole, userRole)
	}
}
