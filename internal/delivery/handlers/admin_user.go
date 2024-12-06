package handlers

import (
	"eliborate/internal/constants"
	"eliborate/internal/delivery/responses"
	"eliborate/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type adminUserHandlers struct {
	service service.AdminUserService
}

func InitAdminUserHandlers(service service.AdminUserService) AdminUserHandlers {
	return adminUserHandlers{
		service: service,
	}
}

// UpdatePassword godoc
// @Summary Update admin user password
// @Description Update admin user password according to his ID
// @Tags admin auth
// @Accept json
// @Produce json
// @Param id body int true "Admin user ID"
// @Param password body string true "New admin user password"
// @Success 200
// @Failure 400 {object} responses.MessageResponse
// @Failure 401 {object} responses.MessageResponse
// @Failure 403 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /admins [patch]
func (a adminUserHandlers) UpdatePassword(c *gin.Context) {
	id, role, err := GetIdAndRoleFromKeys(c.Keys)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewMessageResponseFromErr(err))
		return
	}

	if role != constants.RoleAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, responses.NewMessageResponse(fmt.Sprintf("insufficient role: '%s'", role)))
		return
	}

	body := struct {
		Password string `json:"password"`
	}{
		Password: "",
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	if err := a.service.UpdatePassword(c.Request.Context(), id, body.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	c.Status(http.StatusOK)
}
