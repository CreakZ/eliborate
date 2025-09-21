package handlers

import (
	"eliborate/internal/convertors"
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
	var userCreate dto.UserCreate

	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	userCreateDomain := convertors.DtoUserCreateToDomain(userCreate)

	id, err := u.service.Create(c.Request.Context(), userCreateDomain)
	if err != nil {
		pqErr, ok := err.(*pq.Error)
		if !ok {
			c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		} else if pqErr.Code == "23505" {
			c.JSON(
				http.StatusConflict,
				responses.NewMessageResponse(
					fmt.Sprintf("user with '%s' login already exists", userCreateDomain.Login),
				),
			)
		}
		return
	}

	c.JSON(http.StatusCreated, responses.NewBookCreateResponse(id))
}

// UpdatePassword godoc
// @Summary Update user password
// @Description Updates the password for the user with the given ID
// @Tags user
// @Accept json
// @Produce json
// @Param password body dto.PasswordUpdate true "New password"
// @Security BearerAuth
// @Success 200 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 401 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /users [patch]
func (u userHandlers) UpdatePassword(c *gin.Context) {
	id, err := GetIdFromKeys(c.Keys)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewMessageResponse(err.Error()))
		return
	}

	update := dto.PasswordUpdate{}

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	if err := u.service.UpdatePassword(c.Request.Context(), id, update.Password); err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
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
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.NewMessageResponse(err.Error()))
		return
	}

	if err := u.service.Delete(c.Request.Context(), id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
}
