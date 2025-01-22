package main

import (
	"blog-api/database" // Pacote que contém a função de conexão com o banco de dados
	"fmt"
	"log"
)

func main() {
	// Cria a conexão com o banco de dados
	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Busca os posts no banco de dados
	posts, err := database.GetPosts(db)
	if err != nil {
		log.Fatalf("Erro ao buscar os posts: %v", err)
	}

	// Exibe os posts na tela
	for _, post := range posts {
		fmt.Printf("ID: %s\n", post.ID)
		fmt.Printf("Title: %s\n", post.Title)
		fmt.Printf("Content: %s\n", post.Content)
		fmt.Printf("Date: %s\n", post.Date)
		fmt.Printf("Category: %s\n", post.Category)
		fmt.Printf("Author: %s\n", post.Author)
		fmt.Println("--------")
	}

	fmt.Println("Conexão estabelecida com sucesso!")
}
