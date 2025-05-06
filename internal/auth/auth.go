package auth

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
)

type AuthService struct {
	JWTSecret string
}

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func NewAuthService(JWTSecret string) *AuthService {
	return &AuthService{
		JWTSecret: JWTSecret,
	}
}

func (a *AuthService) GenerateToken(role string) (string, error) {
	claims := &Claims{
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(a.JWTSecret))
	if err != nil {
		return "", fmt.Errorf("ошибка при генерации токена: %w", err)
	}
	return signedToken, nil
}
