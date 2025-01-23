package main

import (
	"blog-api/database"
	"blog-api/middleware"
	"blog-api/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	router.GET("/posts", func(c *gin.Context) {
		posts, err := routes.GetPosts(db)
		if err != nil {
			c.JSON(500, gin.H{
				"error": fmt.Sprintf("Erro ao buscar os posts: %v", err),
			})
			return
		}

		c.JSON(200, posts)
	})

	fmt.Println("Servidor rodando em http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
