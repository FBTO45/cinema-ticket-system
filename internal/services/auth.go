package services

import (
	"cinema-ticket-system/internal/models"
	"cinema-ticket-system/internal/repositories"
	"cinema-ticket-system/pkg/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
	jwtUtil  *utils.JWTUtil
}

func NewAuthService(userRepo *repositories.UserRepository, jwtUtil *utils.JWTUtil) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (as *AuthService) Login(username, password string) (string, time.Time, *models.User, error) {
	user, err := as.userRepo.GetByUsername(username)
	if err != nil {
		return "", time.Time{}, nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", time.Time{}, nil, err
	}

	token, expiresAt, err := as.jwtUtil.GenerateToken(user.UserID)
	if err != nil {
		return "", time.Time{}, nil, err
	}

	return token, expiresAt, user, nil
}