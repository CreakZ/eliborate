package handlers

import (
	"context"
	"eliborate/internal/convertors"
	"eliborate/internal/delivery/responses"
	"eliborate/internal/models/dto"
	"eliborate/internal/service"
	"net/http"
	"strconv"

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
// @Security BearerAuth
// @Success 201 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /categories [post]
func (h categoryHandlers) Create(c *gin.Context) {
	categoryCreate := dto.CategoryCreateUpdate{}

	if err := c.ShouldBindJSON(&categoryCreate); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	if err := h.service.Create(c.Request.Context(), categoryCreate.Name); err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, responses.NewSuccessMessageResponse())
}

// @Summary Get all book categories
// @Description Get all book categories
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Success 200 {object} responses.Categories
// @Failure 500 {object} responses.MessageResponse
// @Router /categories [get]
func (h categoryHandlers) GetAll(c *gin.Context) {
	categories, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	categoriesResp := responses.Categories{
		Categories: convertors.DomainCategoriesToDto(categories),
	}

	c.JSON(http.StatusOK, categoriesResp)
}

// @Summary Update book category
// @Description Update book category (not done yet)
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category id"
// @Param book body dto.CategoryCreateUpdate true "New category name"
// @Success 200 {object} responses.MessageResponse
// @Failure 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /categories/{id} [patch]
func (h categoryHandlers) Update(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("no 'id' provided"))
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	categoryUpdate := dto.CategoryCreateUpdate{}

	if err = c.ShouldBindJSON(&categoryUpdate); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	err = h.service.Update(context.Background(), id, categoryUpdate.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
}

// @Summary Delete book category
// @Description Delete book category (not done yet)
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category id"
// @Success 200 {object} responses.MessageResponse
// @Success 400 {object} responses.MessageResponse
// @Failure 500 {object} responses.MessageResponse
// @Router /categories/{id} [delete]
func (h categoryHandlers) Delete(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("no 'id' provided"))
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	err = h.service.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, responses.NewSuccessMessageResponse())
}
