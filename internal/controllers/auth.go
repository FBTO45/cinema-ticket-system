package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"cinema-ticket-system/internal/models"
	"cinema-ticket-system/internal/repositories"
	"cinema-ticket-system/pkg/utils"
)

type AuthController struct {
	userRepo *repositories.UserRepository
	jwtUtil  *utils.JWTUtil
}

func NewAuthController(userRepo *repositories.UserRepository, jwtUtil *utils.JWTUtil) *AuthController {
	return &AuthController{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      *models.User `json:"user"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.userRepo.GetByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, expiresAt, err := ac.jwtUtil.GenerateToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	response := LoginResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		User:      user,
	}

	c.JSON(http.StatusOK, response)
}