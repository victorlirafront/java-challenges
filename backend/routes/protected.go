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
	if err := middlewares.SessionAuthMiddleware(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
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

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("CSRF validation successful! Welcome, %s", user.Username),
	})
}
