package routes

import (
	"blog-api/models"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SearchPostsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("query", "")
		category := c.DefaultQuery("category", "all")
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "10")

		// Convertendo `page` e `limit` para inteiros
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = 10
		}

		offset := (page - 1) * limit

		// Construindo a consulta SQL (aplicando filtro apenas no title)
		sqlQuery := `
			SELECT 
				id, 
				title, 
				content, 
				date, 
				category, 
				meta_tag_title, 
				meta_tag_description, 
				post_image, 
				post_background, 
				author, 
				keywords 
			FROM posts 
			WHERE title LIKE ?`
		params := []interface{}{fmt.Sprintf("%%%s%%", query)}

		// Adicionando filtro por categoria, se necessário
		if category != "all" && category != "" {
			sqlQuery += " AND category = ?"
			params = append(params, category)
		}

		sqlQuery += " LIMIT ? OFFSET ?"
		params = append(params, limit, offset)

		// Exibindo a consulta para depuração
		fmt.Printf("Query executada: %s\n", sqlQuery)
		fmt.Printf("Parâmetros: %v\n", params)

		// Executando a consulta
		rows, err := db.Query(sqlQuery, params...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao executar a query: %v", err)})
			return
		}
		defer rows.Close()

		// Processando os resultados
		var posts []models.Post
		for rows.Next() {
			var post models.Post
			if err := rows.Scan(
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
			); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao escanear os resultados: %v", err)})
				return
			}
			posts = append(posts, post)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao iterar os resultados: %v", err)})
			return
		}

		// Retornando os resultados
		c.JSON(http.StatusOK, gin.H{
			"data":  posts,
			"page":  page,
			"limit": limit,
			"total": len(posts),
		})
	}
}
