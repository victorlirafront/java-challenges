package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthenticateAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação não fornecido"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação inválido"})
			c.Abort()
			return
		}

		tokenString := authHeader[7:]

		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Chave secreta JWT não configurada"})
			c.Abort()
			return
		}

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		userID, ok := (*claims)["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "ID do usuário não encontrado no token"})
			c.Abort()
			return
		}

		role, ok := (*claims)["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role não encontrado no token"})
			c.Abort()
			return
		}

		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Acesso restrito: somente administradores podem realizar esta ação"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("role", role)
		c.Next()
	}

}
