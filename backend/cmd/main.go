package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/routes"
	"blog-api/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Gerar o token JWT de um usuário administrador (fake)
	userId := os.Getenv("ADMIN_USER_ID")
	token, err := utils.GenerateJWT(userId)
	if err != nil {
		fmt.Println("Erro ao gerar token:", err)
	}
	fmt.Println("Token gerado:", token)

	// Configuração do roteador
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.DatabaseMiddleware(db))
	router.Use(middlewares.CORSMiddleware())

	// Definir rotas
	// Rota para obter posts paginados
	router.GET("/posts", func(context *gin.Context) {
		page, perPage := utils.GetPaginationParams(context)

		posts, err := routes.GetPaginatedPosts(db, page, perPage)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Erro ao buscar os posts: %v", err),
			})
			return
		}

		context.JSON(http.StatusOK, posts)
	})

	// Rota de busca de posts
	router.GET("/search", routes.SearchPostsHandler(db))

	// Rota para deletar um post (requer autenticação)
	router.DELETE("/delete/:id", middlewares.Authenticate, routes.DeletePostHandler)

	// Iniciar servidor
	fmt.Println("Servidor rodando em http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
