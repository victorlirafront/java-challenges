package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")

		// Verifica se o role está na lista de roles permitidos
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		// Se o role não for permitido, retorna erro de acesso negado
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Acesso negado. Permissão insuficiente.",
		})
		c.Abort()
	}
}
