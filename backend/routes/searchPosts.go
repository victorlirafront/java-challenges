package routes

import (
	"blog-api/models"
	"blog-api/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SearchPostsHandler(db *sql.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		page, perPage := utils.GetPaginationParams(context)
		query := context.DefaultQuery("query", "")
		category := context.DefaultQuery("category", "all")
		offset := (page - 1) * perPage

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

		if category != "all" && category != "" {
			sqlQuery += " AND category = ?"
			params = append(params, category)
		}

		sqlQuery += " LIMIT ? OFFSET ?"
		params = append(params, perPage, offset)

		rows, err := db.Query(sqlQuery, params...)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao executar a query: %v", err)})
			return
		}
		defer rows.Close()

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
				context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao escanear os resultados: %v", err)})
				return
			}
			posts = append(posts, post)
		}

		if err := rows.Err(); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao iterar os resultados: %v", err)})
			return
		}

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
			context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao contar os posts: %v", err)})
			return
		}

		totalPages := (totalCount + perPage - 1) / perPage

		context.JSON(http.StatusOK, gin.H{
			"data":       posts,
			"page":       page,
			"perPage":    perPage,
			"total":      totalCount,
			"totalPages": totalPages,
		})
	}
}
