package routes

import (
	"blog-api/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreatePost cria um novo post no banco de dados
func CreatePost(db *sql.DB, c *gin.Context) {
	// Estrutura para armazenar os dados do novo post
	var newPost models.Post

	// Fazendo o bind dos dados JSON da requisição para a struct
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos ou incompletos no corpo da requisição.",
		})
		return
	}

	// Prepara a consulta SQL para inserir o novo post
	query := `INSERT INTO posts (title, content, date, category, meta_tag_title, meta_tag_description, post_image, post_background, author, keywords) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Executa a consulta no banco de dados
	_, err := db.Exec(query, newPost.Title, newPost.Content, newPost.Date, newPost.Category, newPost.MetaTagTitle, newPost.MetaTagDescription,
		newPost.PostImage, newPost.PostBackground, newPost.Author, newPost.Keywords)

	if err != nil {
		log.Printf("Erro ao inserir post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Erro ao criar o post: %v", err),
		})
		return
	}

	// Retorna uma resposta de sucesso
	c.JSON(http.StatusCreated, gin.H{
		"message": "Post criado com sucesso.",
	})
}
