package routes

import (
	"blog-api/models"
	"blog-api/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPaginatedPosts(db *sql.DB, page, perPage int) ([]models.Post, error) {
	offset := (page - 1) * perPage

	query := fmt.Sprintf("SELECT id, title, content, date, category FROM posts LIMIT %d OFFSET %d", perPage, offset)

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar posts: %v", err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Date, &post.Category)
		if err != nil {
			return nil, fmt.Errorf("erro ao processar os posts: %v", err)
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre os posts: %v", err)
	}

	return posts, nil
}

func GetPostsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		page, perPage := utils.GetPaginationParams(c)
		posts, err := getPaginatedPosts(db, page, perPage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Erro ao buscar os posts: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, posts)
	}
}
