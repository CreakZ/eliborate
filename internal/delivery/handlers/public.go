package handlers

import (
	"net/http"
	"yurii-lib/internal/constants"
	"yurii-lib/internal/service"
	utils "yurii-lib/pkg/utils/compare"
	"yurii-lib/pkg/utils/jwt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	token       = "token"
	refreshUUID = "refresh_uuid"
)

type publicHandlers struct {
	publicService service.PublicService
	jwt           jwt.JWT
}

func InitPublicHandlers(publicService service.PublicService, jwt jwt.JWT) PublicHandlers {
	return publicHandlers{
		publicService: publicService,
		jwt:           jwt,
	}
}

// @Summary Logs in an admin user
// @Description Logs in an admin user and returns an access token if the login credentials are valid.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body string true "Login of the admin user"
// @Param password body string true "Password of the admin user"
// @Success 201 {object} map[string]string "Access token"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /public/admin [post]
func (p publicHandlers) LoginAdminUser(c *gin.Context) {
	var userData = struct {
		Login    string `json:"login" validate:"required,containsany"`
		Password string `json:"password" validate:"required,containsany"`
	}{
		Login:    "",
		Password: "",
	}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, password, err := p.publicService.GetByLogin(c.Request.Context(), constants.TypeAdminUsers, userData.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	equal := utils.ComparePasswords([]byte(userData.Password), []byte(password))
	if !equal {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "wrong password"})
		return
	}

	accessToken := p.jwt.CreateToken(id, true)

	c.JSON(http.StatusCreated, gin.H{token: accessToken})
}

// @Summary Logs in a regular user
// @Description Logs in a regular user and returns an access token if the login credentials are valid.
// @Tags auth
// @Accept json
// @Produce json
// @Param login body string true "Login of the user"
// @Param password body string true "Password of the user"
// @Success 201 {object} map[string]string "Access token"
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Not Found"
// @Router /login/user [post]
func (p publicHandlers) LoginUser(c *gin.Context) {
	var userData = struct {
		Login    string `json:"login" validate:"required,containsany"`
		Password string `json:"password" validate:"required,containsany"`
	}{
		Login:    "",
		Password: "",
	}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(userData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, password, err := p.publicService.GetByLogin(c.Request.Context(), constants.TypeUsers, userData.Login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	equal := utils.ComparePasswords([]byte(userData.Password), []byte(password))
	if !equal {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "wrong password"})
		return
	}

	accessToken := p.jwt.CreateToken(id, false)

	c.JSON(http.StatusCreated, gin.H{token: accessToken})
}
