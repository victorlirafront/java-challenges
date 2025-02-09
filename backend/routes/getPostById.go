package routes

import (
	"blog-api/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPostByID(db *sql.DB, postID string) (*models.Post, error) {
	query := "SELECT id, title, content, date, category, meta_tag_title, meta_tag_description, post_image, post_background, author, keywords FROM posts WHERE id = ?"
	row := db.QueryRow(query, postID)

	var post models.Post

	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Date,
		&post.Category,
		&post.MetaTagTitle,
		&post.MetaTagDescription,
		&post.PostImage,
		&post.PostBackground,
		&post.Author,
		&post.Keywords,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post com ID %s não encontrado", postID)
		}
		return nil, fmt.Errorf("erro ao buscar post: %v", err)
	}

	return &post, nil
}

func GetPostByIDHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("id")
		post, err := getPostByID(db, postID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Post não encontrado: %v", err),
			})
			return
		}
		c.JSON(http.StatusOK, post)
	}
}
