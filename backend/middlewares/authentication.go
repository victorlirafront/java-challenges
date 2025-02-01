package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Função para carregar o arquivo .env
func Authenticate(c *gin.Context) {
	// Pega o token do cabeçalho Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação não fornecido"})
		c.Abort()
		return
	}

	// Verifica se o token começa com "Bearer "
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação inválido"})
		c.Abort()
		return
	}

	// Extraí o token (remover o "Bearer " do início)
	tokenString := authHeader[7:]

	// Carrega a chave secreta do .env
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Chave secreta JWT não configurada"})
		c.Abort()
		return
	}

	// Verifica e valida o token
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Aqui usamos a chave secreta do .env para verificar o token
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// Checa se a chave "user_id" existe e é uma string válida no claims
	userID, ok := (*claims)["user_id"].(string)
	fmt.Print("new", userID)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ID do usuário não encontrado no token"})
		c.Abort()
		return
	}

	// Apenas armazena o userID no contexto
	c.Set("userID", userID)
	c.Next()
}
