package routes

import (
	"blog-api/models"
	"blog-api/utils"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte([]byte(password)))
	return err == nil
}

func Login(c *gin.Context) {
	// Obtém o db do contexto
	db := c.MustGet("db").(*sql.DB)

	// Obtém os parâmetros do corpo da requisição
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Verifica se o nome de usuário e a senha têm pelo menos 8 caracteres
	if len(username) < 8 || len(password) < 8 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": "Username and password must be at least 8 characters long",
		})
		return
	}

	// Recupera o usuário do banco de dados
	var user models.User
	result := db.QueryRow("SELECT id, username, hashed_password FROM users WHERE username = ?", username)
	err := result.Scan(&user.ID, &user.Username, &user.HashedPassword)
	if err == sql.ErrNoRows {
		// Se o usuário não for encontrado no banco, retorna erro
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	} else if err != nil {
		// Se houve algum outro erro ao consultar o banco
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while accessing database",
		})
		return
	}

	// Verifica se a senha fornecida corresponde ao hash armazenado
	if !checkPasswordHash(password, user.HashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Gerar o token de sessão e CSRF
	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	// Atualizar o usuário com os tokens gerados
	_, err = db.Exec("UPDATE users SET session_token = ?, csrf_token = ? WHERE id = ?", sessionToken, csrfToken, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user tokens",
		})
		return
	}

	// Configurações do cookie
	cookieExpireDuration := 24 * time.Hour // 24 horas
	cookieSecure := false                  // Defina como true se usar HTTPS
	cookieHttpOnly := true                 // Impede o acesso via JavaScript

	// Calcula a data de expiração dos cookies
	expiration := time.Now().Add(cookieExpireDuration)

	// Define o cookie de sessão (HttpOnly para maior segurança)
	c.SetCookie("session_token", sessionToken, int(time.Until(expiration).Seconds()), "/", "", cookieSecure, cookieHttpOnly)

	// Define o cookie CSRF (não HttpOnly, pois pode ser acessado por JavaScript)
	c.SetCookie("csrf_token", csrfToken, int(time.Until(expiration).Seconds()), "/", "", cookieSecure, false)

	// Resposta de sucesso de login
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user.Username,
	})
}
