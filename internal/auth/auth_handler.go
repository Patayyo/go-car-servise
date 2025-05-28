package auth

import (
	"car-service/internal/dto"
	"car-service/internal/model"
	"car-service/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Reg(r *gin.Engine) {
	r.POST("/signup", h.SignUp)
	r.POST("/login", h.Login)
}

// @Summary Регистрация пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.SignUpInput true "Пользователь"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /signup [post]
func (h *AuthHandler) SignUp(c *gin.Context) {
	var input dto.SignUpInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Errorf("invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.service.SignUp(&model.User{
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		logrus.Errorf("failed to register user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	logrus.Info("user registered successfully", input.Email)
	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

// @Summary Авторизация пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param user body dto.LoginInput true "Пользователь"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Errorf("invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	token, err := h.service.Login(input.Email, input.Password)
	if err != nil {
		logrus.Errorf("login failed: %v", err)
		if err.Error() == "invalid email or password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login user"})
		}
		return
	}

	logrus.Infof("user logged successfully: %s", input.Email)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
