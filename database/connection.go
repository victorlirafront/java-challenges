package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importa o driver MySQL
	"github.com/joho/godotenv"
)

// Estrutura que representa um post
type Post struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	Content            string `json:"content"`
	Date               string `json:"date"`
	Category           string `json:"category"`
	MetaTagTitle       string `json:"meta_tag_title"`
	MetaTagDescription string `json:"meta_tag_description"`
	PostImage          string `json:"post_image"`
	PostBackground     string `json:"post_background"`
	Author             string `json:"author"`
	Keywords           string `json:"keywords"`
}

func CreateDatabaseConnection() (*sql.DB, error) {
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir a conexão: %v", err)
	}

	// Testar a conexão
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Conexão bem-sucedida ao banco de dados MySQL!")
	return db, nil
}

func GetPosts(db *sql.DB) ([]Post, error) {
	// Consulta SQL para buscar todos os posts
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, fmt.Errorf("Erro ao executar a consulta: %v", err)
	}
	defer rows.Close()

	var posts []Post

	// Processa os resultados da consulta
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Date, &post.Category, &post.MetaTagTitle, &post.MetaTagDescription, &post.PostImage, &post.PostBackground, &post.Author, &post.Keywords); err != nil {
			return nil, fmt.Errorf("Erro ao ler os dados da linha: %v", err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Erro ao iterar sobre as linhas: %v", err)
	}

	return posts, nil
}
