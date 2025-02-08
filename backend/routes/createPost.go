package routes

import (
	"blog-api/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(db *sql.DB, c *gin.Context) {
	var newPost models.Post

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos ou incompletos no corpo da requisição.",
		})
		return
	}

	query := `INSERT INTO posts (title, content, date, category, meta_tag_title, meta_tag_description, post_image, post_background, author, keywords) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query, newPost.Title, newPost.Content, newPost.Date, newPost.Category, newPost.MetaTagTitle, newPost.MetaTagDescription,
		newPost.PostImage, newPost.PostBackground, newPost.Author, newPost.Keywords)

	if err != nil {
		log.Printf("Erro ao inserir post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Erro ao criar o post: %v", err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post criado com sucesso.",
	})
}
