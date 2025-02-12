package routes

import (
	"blog-api/models"
	"blog-api/utils"
	"database/sql"
	"fmt"
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
	db := c.MustGet("db").(*sql.DB)

	username := c.PostForm("username")
	password := c.PostForm("password")

	if len(username) < 8 || len(password) < 8 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": "Username and password must be at least 8 characters long",
		})
		return
	}

	var user models.User
	result := db.QueryRow("SELECT id, username, hashed_password, role FROM users WHERE username = ?", username)
	err := result.Scan(&user.ID, &user.Username, &user.HashedPassword, &user.Role)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while accessing database",
		})
		return
	}

	if !checkPasswordHash(password, user.HashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	_, err = db.Exec("UPDATE users SET session_token = ?, csrf_token = ? WHERE id = ?", sessionToken, csrfToken, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user tokens",
		})
		return
	}

	accessToken, err := utils.GenerateAdminToken(user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("erro ao gerar token: %v", err)})
		return
	}

	cookieExpireDuration := 24 * time.Hour
	cookieSecure := false
	cookieHttpOnly := true
	expiration := time.Now().Add(cookieExpireDuration)

	c.SetCookie("session_token", sessionToken, int(time.Until(expiration).Seconds()), "/", "", cookieSecure, cookieHttpOnly)
	c.SetCookie("csrf_token", csrfToken, int(time.Until(expiration).Seconds()), "/", "", cookieSecure, false)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user.Username,
		"token":   accessToken,
		"role":    user.Role,
	})
}
