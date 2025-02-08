package routes

import (
	"blog-api/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hasPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func Register(c *gin.Context) {
	// Acessando o db do contexto, mas agora como *sql.DB
	db := c.MustGet("db").(*sql.DB)

	username := c.PostForm("username")
	password := c.PostForm("password")

	// Verificar se o nome de usuário e a senha têm pelo menos 8 caracteres
	if len(username) < 8 || len(password) < 8 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": "Username and password must be at least 8 characters long",
		})
		return
	}

	// Verifica se o nome de usuário já existe no banco de dados
	var user models.User
	result := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
	err := result.Scan(&user.ID, &user.Username, &user.HashedPassword)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "User already exists",
		})
		return
	}

	// Gerar a senha criptografada
	hashedPassword, err := hasPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Criar um novo usuário no banco de dados usando Query ou Exec (sem GORM)
	_, err = db.Exec("INSERT INTO users (username, hashed_password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}
