package routes

import (
	"blog-api/middlewares"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	if err := middlewares.SessionAuthMiddleware(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	username := c.DefaultPostForm("username", "")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username is required",
		})
		return
	}

	db := c.MustGet("db").(*sql.DB)

	_, err := db.Exec("UPDATE users SET session_token = ?, csrf_token = ? WHERE username = ?", "", "", username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while logging out",
		})
		return
	}

	c.SetCookie("session_token", "", -1, "/", "", true, true)
	c.SetCookie("csrf_token", "", -1, "/", "", true, false)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}
