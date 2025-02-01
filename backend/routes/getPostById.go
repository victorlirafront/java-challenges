package routes

import (
	"blog-api/models"
	"database/sql"
	"fmt"
)

// GetPostByID busca um post específico pelo ID
func GetPostByID(db *sql.DB, postID string) (*models.Post, error) {
	// Consulta SQL para buscar um post pelo ID
	query := "SELECT id, title, content, date, category, meta_tag_title, meta_tag_description, post_image, post_background, author, keywords FROM posts WHERE id = ?"

	// Executa a consulta
	row := db.QueryRow(query, postID)

	// Cria um modelo para armazenar os dados do post
	var post models.Post

	// Faz o scan dos dados do post
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
		// Se o erro for "no rows", significa que o post não foi encontrado
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post com ID %s não encontrado", postID)
		}
		return nil, fmt.Errorf("erro ao buscar post: %v", err)
	}

	// Retorna o post encontrado
	return &post, nil
}
