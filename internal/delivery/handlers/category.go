package handlers

import (
	"eliborate/internal/delivery/responses"
	"eliborate/internal/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type categoryHandlers struct {
	service service.CategoryService
}

func NewCategoryHandlers(service service.CategoryService) CategoryHandlers {
	return categoryHandlers{
		service: service,
	}
}

// @Summary Create new book category
// @Description Create new book category
// @Tags categories
// @Accept json
// @Produce json
// @Success 201 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /categories [post]
func (h categoryHandlers) Create(c *gin.Context) {
	body := struct {
		Name string `json:"name"`
	}{
		Name: "",
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.NewMessageResponseFromErr(err))
		return
	}

	if err := h.service.Create(c.Request.Context(), body.Name); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	c.JSON(http.StatusCreated, responses.NewMessageResponse(fmt.Sprintf("category '%s' created successfully", body.Name)))
}

// @Summary Get all book categories
// @Description Get all book categories
// @Tags categories
// @Produce json
// @Success 201 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /categories [get]
func (h categoryHandlers) GetAll(c *gin.Context) {
	categories, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, responses.NewMessageResponseFromErr(err))
		return
	}

	body := struct {
		Categories []string `json:"categories"`
	}{
		Categories: categories,
	}

	c.JSON(http.StatusOK, body)
}

// @Summary Update book category
// @Description Update book category (not done yet)
// @Tags categories
// @Produce json
// @Success 201 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Failure 503                            "handler is unavailable"
// @Router /categories [patch]
func (h categoryHandlers) Update(c *gin.Context) {
	c.AbortWithStatus(http.StatusServiceUnavailable)
}

// @Summary Delete book category
// @Description Delete book category (not done yet)
// @Tags categories
// @Produce json
// @Success 201 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Failure 503                            "handler is unavailable"
// @Router /categories [delete]
func (h categoryHandlers) Delete(c *gin.Context) {
	c.AbortWithStatus(http.StatusServiceUnavailable)
}
