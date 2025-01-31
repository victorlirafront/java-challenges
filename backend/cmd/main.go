package main

import (
	"blog-api/database"
	"blog-api/middleware"
	"blog-api/routes"
	"blog-api/utils"
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
	router.Use(middleware.DatabaseMiddleware(db))
	router.Use(middleware.CORSMiddleware())

	router.GET("/posts", func(context *gin.Context) {
		page, perPage := utils.GetPaginationParams(context)

		posts, err := routes.GetPaginatedPosts(db, page, perPage)
		if err != nil {
			context.JSON(500, gin.H{
				"error": fmt.Sprintf("Erro ao buscar os posts: %v", err),
			})
			return
		}

		context.JSON(200, posts)
	})

	router.GET("/search", routes.SearchPostsHandler(db))

	router.DELETE("/delete/:id", routes.DeletePostHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
