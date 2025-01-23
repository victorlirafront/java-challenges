package routes

import (
	"blog-api/models"
	"blog-api/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Função que retorna os parâmetros de paginação
func SearchPostsHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtendo os parâmetros de página e limite da URL
		page, perPage := utils.GetPaginationParams(c.Request)
		query := c.DefaultQuery("query", "")
		category := c.DefaultQuery("category", "all")

		// Calculando o offset com base na página e no número de itens por página
		offset := (page - 1) * perPage

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
		params = append(params, perPage, offset)

		// Executando a consulta para os posts
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

		// Verificando erros ao iterar os resultados
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao iterar os resultados: %v", err)})
			return
		}

		// Contando o total de posts para a mesma consulta (sem LIMIT)
		countQuery := `
			SELECT COUNT(*) 
			FROM posts 
			WHERE title LIKE ?`
		var paramsCount []interface{}
		paramsCount = append(paramsCount, fmt.Sprintf("%%%s%%", query))

		if category != "all" && category != "" {
			countQuery += " AND category = ?"
			paramsCount = append(paramsCount, category)
		}

		var totalCount int
		err = db.QueryRow(countQuery, paramsCount...).Scan(&totalCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao contar os posts: %v", err)})
			return
		}

		// Calculando o total de páginas
		totalPages := (totalCount + perPage - 1) / perPage // Total de páginas com base no número total de posts e no limite

		// Retornando os resultados com a contagem de total de páginas
		c.JSON(http.StatusOK, gin.H{
			"data":       posts,
			"page":       page,
			"perPage":    perPage,
			"total":      totalCount,
			"totalPages": totalPages,
		})
	}
}
