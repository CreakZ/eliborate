package handlers

import (
	"eliborate/internal/delivery/responses"
	"eliborate/internal/models/dto"
	"eliborate/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
// @Accept json
// @Produce json
// @Param user body dto.UserCreate true "User information"
// @Success 201 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 404 {object} responses.MessageResponse
// @Failure 409 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /users [post]
func (u userHandlers) Create(c *gin.Context) {
	var user dto.UserCreate

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	id, err := u.service.Create(c.Request.Context(), user)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		}
		if pqErr.Code == "23505" {
			c.AbortWithStatusJSON(http.StatusConflict, responses.NewMessageResponse(fmt.Sprintf("user with '%s' login already exists", user.Login)))
		}
		return
	}

	c.JSON(http.StatusCreated, responses.NewMessageResponse(fmt.Sprintf("user with id %d created successfully", id)))
}

// UpdatePassword godoc
// @Summary Update user password
// @Description Updates the password for the user with the given ID
// @Tags user
// @Accept json
// @Produce json
// @Param id body int true "User ID"
// @Param password body string true "New password"
// @Security BearerAuth
// @Success 200
// @Failure 400 {object} responses.MessageResponse
// @Failure 401 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /users [patch]
func (u userHandlers) UpdatePassword(c *gin.Context) {
	id, err := GetIdFromKeys(c.Keys)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewMessageResponseFromErr(err))
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

	if err := u.service.UpdatePassword(c.Request.Context(), id, body.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	c.Status(http.StatusOK)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes the user with the provided ID
// @Tags user
// @Accept json
// @Produce json
// @Param id body int true "User ID"
// @Security BearerAuth
// @Success 200
// @Failure 400 {object} responses.MessageResponse
// @Failure 401 {object} responses.MessageResponse
// @Router /users [delete]
func (u userHandlers) Delete(c *gin.Context) {
	id, err := GetIdFromKeys(c.Keys)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewMessageResponseFromErr(err))
		return
	}

	if err := u.service.Delete(c.Request.Context(), id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	c.Status(http.StatusOK)
}
