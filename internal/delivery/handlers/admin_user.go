package handlers

import (
	"eliborate/internal/constants"
	"eliborate/internal/delivery/responses"
	"eliborate/internal/models/dto"
	"eliborate/internal/service"
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
// @Param password body string true "New admin user password"
// @Security BearerAuth
// @Success 200 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 401 {object} responses.MessageResponse
// @Failure 403 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /admins [patch]
func (a adminUserHandlers) UpdatePassword(c *gin.Context) {
	id, role, err := GetIdAndRoleFromKeys(c.Keys)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.NewMessageResponse(err.Error()))
		return
	}

	if role != constants.RoleAdmin {
		c.JSON(http.StatusForbidden, responses.NewMessageResponse("insufficient role"))
		return
	}

	update := dto.PasswordUpdate{}

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	if err := a.service.UpdatePassword(c.Request.Context(), id, update.Password); err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
}
