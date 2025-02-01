package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Função para autenticação padrão (somente leitura) com permissões baseadas em role
func AuthenticateRead(c *gin.Context) {
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
		return []byte(secretKey), nil
	})

	fmt.Print("regular token ", token)

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
		c.Abort()
		return
	}

	// Checa se a chave "user_id" existe e é uma string válida no claims
	userID, ok := (*claims)["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ID do usuário não encontrado no token"})
		c.Abort()
		return
	}

	// Verifica se o ID do usuário é o ID do administrador ou do usuário regular
	adminUserID := os.Getenv("ADMIN_USER_ID")
	regularUserID := os.Getenv("REGULAR_USER_ID")

	if userID == adminUserID {
		// Se for o admin, armazena isso no contexto
		c.Set("role", "admin")
	} else if userID == regularUserID {
		// Se for o usuário regular, armazena isso no contexto
		c.Set("role", "regular")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autorizado"})
		c.Abort()
		return
	}

	// Armazena o userID no contexto
	c.Set("userID", userID)
	c.Next()
}
