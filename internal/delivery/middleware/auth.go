package middleware

import (
	"eliborate/internal/constants"
	"eliborate/internal/delivery/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m Middleware) BearerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewMessageResponse("no jwt provided"))
			return
		}

		claim, valid, err := m.jwtUtil.Authorize(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
			return
		}

		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewMessageResponse("jwt is not valid"))
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
