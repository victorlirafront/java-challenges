package routes

import (
	"blog-api/middlewares"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	// Chama a função Authorize para verificar a autenticação e CSRF
	if err := middlewares.SessionAuthMiddleware(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Obtém o parâmetro "username" do formulário
	username := c.DefaultPostForm("username", "")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username is required",
		})
		return
	}

	// Acessa o banco de dados para atualizar o usuário, limpando os tokens
	db := c.MustGet("db").(*sql.DB)

	// Atualiza os valores dos tokens para nulos no banco de dados
	_, err := db.Exec("UPDATE users SET session_token = ?, csrf_token = ? WHERE username = ?", "", "", username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while logging out",
		})
		return
	}

	// Limpa os cookies de sessão e CSRF no cliente
	c.SetCookie("session_token", "", -1, "/", "", true, true) // Cookie com HttpOnly
	c.SetCookie("csrf_token", "", -1, "/", "", true, false)   // Cookie sem HttpOnly

	// Retorna resposta de sucesso
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
