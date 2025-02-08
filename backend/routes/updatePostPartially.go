package routes

import (
	"blog-api/database"
	"blog-api/models"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func updatePostPartially(db *sql.DB, postID string, postUpdate *models.Post) error {
	query := "UPDATE posts SET "
	params := []interface{}{}
	setClauses := []string{}

	if postUpdate.Title != "" {
		setClauses = append(setClauses, "title = ?")
		params = append(params, postUpdate.Title)
	}
	if postUpdate.Content != "" {
		setClauses = append(setClauses, "content = ?")
		params = append(params, postUpdate.Content)
	}
	if postUpdate.Date != "" {
		setClauses = append(setClauses, "date = ?")
		params = append(params, postUpdate.Date)
	}
	if postUpdate.Category != "" {
		setClauses = append(setClauses, "category = ?")
		params = append(params, postUpdate.Category)
	}
	if postUpdate.MetaTagTitle != "" {
		setClauses = append(setClauses, "meta_tag_title = ?")
		params = append(params, postUpdate.MetaTagTitle)
	}
	if postUpdate.MetaTagDescription != "" {
		setClauses = append(setClauses, "meta_tag_description = ?")
		params = append(params, postUpdate.MetaTagDescription)
	}
	if postUpdate.PostImage != "" {
		setClauses = append(setClauses, "post_image = ?")
		params = append(params, postUpdate.PostImage)
	}
	if postUpdate.PostBackground != "" {
		setClauses = append(setClauses, "post_background = ?")
		params = append(params, postUpdate.PostBackground)
	}
	if postUpdate.Author != "" {
		setClauses = append(setClauses, "author = ?")
		params = append(params, postUpdate.Author)
	}
	if postUpdate.Keywords != "" {
		setClauses = append(setClauses, "keywords = ?")
		params = append(params, postUpdate.Keywords)
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("nenhum campo para atualizar")
	}

	query += joinClauses(setClauses) + " WHERE id = ?"
	params = append(params, postID)

	_, err := db.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("erro ao atualizar o post: %v", err)
	}

	return nil
}

func joinClauses(clauses []string) string {
	return stringJoin(clauses, ", ")
}

func stringJoin(slice []string, separator string) string {
	result := ""
	for i, str := range slice {
		if i > 0 {
			result += separator
		}
		result += str
	}
	return result
}
func CallUpdatePost(c *gin.Context) {
	postID := c.Param("id")

	db, err := database.CreateDatabaseConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err),
		})
		return
	}
	defer db.Close()

	var postUpdate models.Post
	if err := c.ShouldBindJSON(&postUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos ou incompletos no corpo da requisição.",
		})
		return
	}

	err = updatePostPartially(db, postID, &postUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Erro ao atualizar o post: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post atualizado com sucesso.",
	})
}
