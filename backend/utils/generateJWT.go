package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateJWT(userID string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("chave secreta não configurada")
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAdminToken() (string, error) {
	userAdmin := os.Getenv("REGULAR_USER_ID")
	if userAdmin == "" {
		return "", fmt.Errorf("REGULAR_USER_ID não está definido no ambiente")
	}

	tokenAdmin, err := generateJWT(userAdmin)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token: %v", err)
	}

	return tokenAdmin, nil
}
