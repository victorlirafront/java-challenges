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
	// Criação de banco de dados (caso seja necessário)
	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Configuração do roteador Gin
	router := gin.Default()

	// Definir Middlewares (se necessário)
	router.Use(middlewares.DatabaseMiddleware(db))
	router.Use(middlewares.CORSMiddleware())

	// Definir rotas
	router.POST("/register", routes.Register)   // Rota de registro
	router.POST("/login", routes.Login)         // Rota de login
	router.POST("/logout", routes.Logout)       // Rota de logout
	router.POST("/protected", routes.Protected) // Rota de acesso protegido

	// Configuração para posts
	router.GET("/posts", func(c *gin.Context) {
		// Pega o parâmetro de paginação
		page, perPage := utils.GetPaginationParams(c)

		// Tenta buscar os posts com paginação
		posts, err := routes.GetPaginatedPosts(db, page, perPage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Erro ao buscar os posts: %v", err),
			})
			return
		}

		// Retorna os posts encontrados
		c.JSON(http.StatusOK, posts)
	})

	// Rota de busca de posts
	router.GET("/search", middlewares.AuthenticateRead, middlewares.RoleMiddleware("regular", "admin"), routes.SearchPostsHandler(db))

	// Definir rota para criar um post
	router.POST("/posts", middlewares.AuthenticateAdmin, func(c *gin.Context) {
		routes.CreatePost(db, c) // Chama a função de criar post
	})

	// Rota para obter um post específico pelo ID
	router.GET("/posts/:id", middlewares.AuthenticateRead, middlewares.RoleMiddleware("regular", "admin"), func(c *gin.Context) {
		postID := c.Param("id")

		// Chama a função para buscar o post pelo ID
		post, err := routes.GetPostByID(db, postID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Post não encontrado: %v", err),
			})
			return
		}

		// Retorna o post encontrado
		c.JSON(http.StatusOK, post)
	})

	// Rota para atualizar um post parcialmente pelo ID
	router.PATCH("/posts/:id", middlewares.AuthenticateAdmin, func(c *gin.Context) {
		routes.CallUpdatePost(c)
	})

	// Gerar o token JWT de um usuário administrador (fake)
	userAdmin := os.Getenv("REGULAR_USER_ID")
	if userAdmin == "" {
		log.Fatal("REGULAR_USER_ID não está definido no ambiente")
	}
	tokenAdmin, err := utils.GenerateJWT(userAdmin)
	if err != nil {
		log.Fatalf("Erro ao gerar token: %v", err)
	}
	fmt.Println("Token gerado:", tokenAdmin)

	// Rota para deletar um post (requer autenticação)
	router.DELETE("/delete/:id", middlewares.AuthenticateAdmin, routes.DeletePostHandler)

	// Iniciar servidor
	fmt.Println("Servidor rodando em http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
