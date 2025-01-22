package routes

import (
	"database/sql"
	"fmt"
)

// Estrutura para os posts
type Post struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Date     string `json:"date"`
	Category string `json:"category"`
}

// Função para buscar os posts do banco de dados e retornar para o cliente
func GetPosts(db *sql.DB) ([]Post, error) {
	// Consultar todos os posts
	rows, err := db.Query("SELECT id, title, content, date, category FROM posts")
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar posts: %v", err)
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
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
