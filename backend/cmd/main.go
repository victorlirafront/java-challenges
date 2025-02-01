package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/routes"
	"blog-api/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Configuração do roteador
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.DatabaseMiddleware(db))
	router.Use(middlewares.CORSMiddleware())

	// Definir rotas
	// Rota para obter posts paginados
	router.GET("/posts", middlewares.AuthenticateRead, func(c *gin.Context) {
		// Recupera o role do usuário a partir do contexto
		role, _ := c.Get("role")

		fmt.Print("role ", role)

		// Verifica se o usuário tem permissão para acessar os posts
		if role == "regular" || role == "admin" {
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
		} else {
			// Se o role não for regular ou admin, retorna erro de acesso negado
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Acesso negado. Permissão insuficiente.",
			})
		}
	})

	// Rota de busca de posts
	router.GET("/search", middlewares.AuthenticateRead, routes.SearchPostsHandler(db))

	// Definir rota para criar um post
	router.POST("/posts", middlewares.AuthenticateAdmin, func(c *gin.Context) {
		routes.CreatePost(db, c) // Chama a função de criar post
	})

	// Rota para obter um post específico pelo ID
	router.GET("/posts/:id", func(c *gin.Context) {
		postID := c.Param("id")

		fmt.Print(postID)

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

	// userAdmin := os.Getenv("ADMIN_USER_ID") //pay attentino
	// tokenAdmin, err := utils.GenerateJWT(userAdmin)
	// if err != nil {
	// 	fmt.Println("Erro ao gerar token:", err)
	// }
	// fmt.Println("Token gerado:", tokenAdmin)

	// Rota para deletar um post (requer autenticação)
	router.DELETE("/delete/:id", middlewares.AuthenticateAdmin, routes.DeletePostHandler)

	// Iniciar servidor
	fmt.Println("Servidor rodando em http://localhost:8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
