package routes

import (
	"blog-api/models"
	"database/sql"
	"fmt"
)

func GetPaginatedPosts(db *sql.DB, page, perPage int) ([]models.Post, error) {
	offset := (page - 1) * perPage

	// Consulta SQL com paginação
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

	// Verifica se houve erro ao iterar sobre as linhas
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar sobre os posts: %v", err)
	}

	return posts, nil
}
