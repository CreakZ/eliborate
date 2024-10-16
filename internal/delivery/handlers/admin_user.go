package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"yurii-lib/internal/service"
	"yurii-lib/internal/validators"

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

/*
func (a adminUserHandlers) Create(c *gin.Context) {

}
*/

// GetPassword godoc
// @Summary Get admin user password
// @Description Returns admin user according to his ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param  id body int true "admin ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404
// @Failure 500 {object} map[string]string
// @Router /admin [get]
func (a adminUserHandlers) GetPassword(c *gin.Context) {
	body := struct {
		ID int `json:"id"`
	}{
		ID: 0,
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	password, err := a.service.GetPassword(c.Request.Context(), body.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"password": password})
}

// UpdatePassword godoc
// @Summary Update admin user password
// @Description Update admin user password according to his ID
// @Tags admin
// @Accept  json
// @Produce  json
// @Param  id body int true "ID администратора"
// @Param  password body string true "Новый пароль"
// @Success 200
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /admin [put]
func (a adminUserHandlers) UpdatePassword(c *gin.Context) {
	body := struct {
		ID       int    `json:"id"`
		Password string `json:"password"`
	}{
		ID:       0,
		Password: "",
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	valid := validators.ValidatePassword(body.Password)
	if !valid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "password is invalid"})
		return
	}

	if err := a.service.UpdatePassword(c.Request.Context(), body.ID, body.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

/*
func (a adminUserHandlers) Delete(c *gin.Context) {

}
*/
