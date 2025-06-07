package handlers

import (
	"eliborate/internal/models/dto"
	"eliborate/internal/service"
	utils "eliborate/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type publicHandlers struct {
	publicService service.PublicService
	jwt           utils.JWT
}

func InitPublicHandlers(publicService service.PublicService, jwt utils.JWT) PublicHandlers {
	return publicHandlers{
		publicService: publicService,
		jwt:           jwt,
	}
}

// @Summary Login an admin user
// @Description Logs in an admin user and returns an access token if the login credentials are valid.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.AdminUserCreate true "Login of the admin user"
// @Success 201 {object} map[string]string "Signed JWT"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /public/admin [post]
func (p publicHandlers) LoginAdminUser(c *gin.Context) {
	userData := dto.Credentials{}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := p.publicService.GetAdminUserByLogin(c.Request.Context(), userData.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "wrong password provided"})
		return
	}

	accessToken := p.jwt.CreateToken(user.ID, true)
	c.JSON(http.StatusCreated, gin.H{"jwt": accessToken})
}

// @Summary Logs in a regular user
// @Description Logs in a regular user and returns an access token if the login credentials are valid.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.UserCreate true "Login of the user"
// @Success 201 {object} map[string]string "Access token"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Not Found"
// @Router /public/user [post]
func (p publicHandlers) LoginUser(c *gin.Context) {
	userData := dto.Credentials{}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := p.publicService.GetUserByLogin(c.Request.Context(), userData.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "wrong password provided"})
		return
	}

	accessToken := p.jwt.CreateToken(user.ID, false)
	c.JSON(http.StatusCreated, gin.H{"jwt": accessToken})
}
