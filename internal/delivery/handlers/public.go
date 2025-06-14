package handlers

import (
	"eliborate/internal/delivery/responses"
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
// @Param credentials body dto.Credentials true "Login of the admin user"
// @Header 201 {object} map[string]string "Signed JWT"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /auth/admin [post]
func (p publicHandlers) LoginAdminUser(c *gin.Context) {
	userData := dto.Credentials{}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse(err.Error()))
		return
	}

	user, err := p.publicService.GetAdminUserByLogin(c.Request.Context(), userData.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.NewMessageResponse(err.Error()))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewMessageResponse("wrong password provided"))
		return
	}

	jwt := p.jwt.CreateToken(user.ID, true)

	c.JSON(http.StatusCreated, responses.NewJwtResponse(jwt))
}

// @Summary Logs in a regular user
// @Description Logs in a regular user and returns an access token if the login credentials are valid.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body dto.Credentials true "Login of the user"
// @Success 201 {object} responses.JwtResponse "Access token"
// @Failure 400 {object} responses.MessageResponse "Bad Request"
// @Failure 404 {object} responses.MessageResponse "Not Found"
// @Router /auth/user [post]
func (p publicHandlers) LoginUser(c *gin.Context) {
	userData := dto.Credentials{}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := p.publicService.GetUserByLogin(c.Request.Context(), userData.Login)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong password provided"})
		return
	}

	jwt := p.jwt.CreateToken(user.ID, false)

	c.JSON(http.StatusCreated, responses.NewJwtResponse(jwt))
}
