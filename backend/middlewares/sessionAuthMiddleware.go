package middlewares

import (
	"blog-api/models"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware(c *gin.Context) error {
	// Obtém os tokens dos cookies
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		return fmt.Errorf("session token not found")
	}

	csrfToken, err := c.Cookie("csrf_token")
	if err != nil {
		return fmt.Errorf("csrf token not found")
	}

	// Acessa o banco de dados para verificar os tokens
	db := c.MustGet("db").(*sql.DB)

	var user models.User
	err = db.QueryRow("SELECT id, session_token, csrf_token FROM users WHERE session_token = ? AND csrf_token = ?", sessionToken, csrfToken).Scan(&user.ID, &user.SessionToken, &user.CSRFToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid session or csrf token")
		}
		return fmt.Errorf("error while accessing database")
	}

	return nil
}
