package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTUtil struct {
	secretKey string
	expiresIn time.Duration
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewJWTUtil(secretKey string, expiresIn time.Duration) *JWTUtil {
	return &JWTUtil{
		secretKey: secretKey,
		expiresIn: expiresIn,
	}
}

func (ju *JWTUtil) GenerateToken(userID uint) (string, time.Time, error) {
	expirationTime := time.Now().Add(ju.expiresIn)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(ju.secretKey))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

func (ju *JWTUtil) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(ju.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}