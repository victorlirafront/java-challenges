package routes

import (
	"blog-api/middlewares"
	"blog-api/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Protected(c *gin.Context) {
	// Chama a função Authorize para verificar a autenticação e CSRF
	if err := middlewares.SessionAuthMiddleware(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Obtém o parâmetro "username" do formulário ou query string, dependendo da requisição
	username := c.DefaultPostForm("username", "")
	if username == "" {
		username = c.DefaultQuery("username", "") // Usado se for uma requisição GET
	}

	// Verifica se o parâmetro username está presente
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username is required",
		})
		return
	}

	// Recupera o session_token e csrf_token dos cookies
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Session token not found",
		})
		return
	}

	csrfToken, err := c.Cookie("csrf_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "CSRF token not found",
		})
		return
	}

	// Acessa o banco de dados para verificar o usuário com base no sessionToken e csrfToken
	db := c.MustGet("db").(*sql.DB)

	var user models.User
	err = db.QueryRow("SELECT id, username, session_token, csrf_token FROM users WHERE session_token = ? AND csrf_token = ?", sessionToken, csrfToken).Scan(&user.ID, &user.Username, &user.SessionToken, &user.CSRFToken)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid session or CSRF token",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while accessing database",
			})
		}
		return
	}

	// Retorna uma mensagem de sucesso se a validação for bem-sucedida
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("CSRF validation successful! Welcome, %s", username),
	})
}
