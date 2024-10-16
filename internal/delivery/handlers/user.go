package handlers

import (
	"fmt"
	"net/http"
	"yurii-lib/internal/models/dto"
	"yurii-lib/internal/service"
	"yurii-lib/internal/validators"

	"github.com/gin-gonic/gin"
)

type userHandlers struct {
	service service.UserService
}

func InitUserHandlers(service service.UserService) UserHandlers {
	return userHandlers{
		service: service,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Creates a new user with provided login and password
// @Tags user
// @Accept  json
// @Produce  json
// @Param  user body dto.UserCreate true "User information"
// @Success 201 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 418 {object} map[string]string
// @Router /user [post]
func (u userHandlers) Create(c *gin.Context) {
	var user dto.UserCreate

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	valid := validators.ValidatePassword(user.Password)
	if !valid {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "insecure password"})
		return
	}

	exists, err := u.service.CheckByLogin(c.Request.Context(), user.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if exists {
		c.AbortWithStatusJSON(
			http.StatusConflict,
			gin.H{"message": fmt.Sprintf("user with '%s' login already exists", user.Login)},
		)
		return
	}

	id, err := u.service.Create(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusTeapot, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetPassword godoc
// @Summary Get user password by ID
// @Description Retrieves the password for a user given their ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param  id body int true "User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 418 {object} map[string]string
// @Router /user [get]
func (u userHandlers) GetPassword(c *gin.Context) {
	body := struct {
		ID int `json:"id"`
	}{
		ID: 0,
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if body.ID < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is less than 1"})
		return
	}

	password, err := u.service.GetPassword(c.Request.Context(), body.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusTeapot, gin.H{"message": err.Error()})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, gin.H{"password": password})
}

// UpdatePassword godoc
// @Summary Update user password
// @Description Updates the password for the user with the given ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param  id body int true "User ID"
// @Param  password body string true "New password"
// @Success 200
// @Failure 400 {object} map[string]string
// @Router /user [put]
func (u userHandlers) UpdatePassword(c *gin.Context) {
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

	if err := u.service.UpdatePassword(c.Request.Context(), body.ID, body.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes the user with the provided ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param  id body int true "User ID"
// @Success 200
// @Failure 400 {object} map[string]string
// @Router /user [delete]
func (u userHandlers) Delete(c *gin.Context) {
	body := struct {
		ID int `json:"id"`
	}{
		ID: 0,
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if body.ID < 1 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "id is less than 1"})
		return
	}

	if err := u.service.Delete(c.Request.Context(), body.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}
