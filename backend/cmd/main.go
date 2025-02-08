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
	db, err := database.CreateDatabaseConnection()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	router := gin.Default()
	router.Use(middlewares.DatabaseMiddleware(db))
	router.Use(middlewares.CORSMiddleware())

	router.POST("/register", routes.Register)
	router.POST("/login", routes.Login)
	router.POST("/logout", routes.Logout)
	router.POST("/protected", routes.Protected)

	router.GET("/posts", func(c *gin.Context) {

		page, perPage := utils.GetPaginationParams(c)
		posts, err := routes.GetPaginatedPosts(db, page, perPage)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Erro ao buscar os posts: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, posts)
	})

	router.GET("/search", middlewares.AuthenticateRead, middlewares.RoleMiddleware("regular", "admin"), routes.SearchPostsHandler(db))

	router.POST("/posts", middlewares.AuthenticateAdmin, func(c *gin.Context) {
		routes.CreatePost(db, c)
	})

	router.GET("/posts/:id", middlewares.AuthenticateRead, middlewares.RoleMiddleware("regular", "admin"), func(c *gin.Context) {
		postID := c.Param("id")

		post, err := routes.GetPostByID(db, postID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Post não encontrado: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, post)
	})

	router.PATCH("/posts/:id", middlewares.AuthenticateAdmin, func(c *gin.Context) {
		routes.CallUpdatePost(c)
	})

	userAdmin := os.Getenv("REGULAR_USER_ID")
	if userAdmin == "" {
		log.Fatal("REGULAR_USER_ID não está definido no ambiente")
	}
	tokenAdmin, err := utils.GenerateJWT(userAdmin)
	if err != nil {
		log.Fatalf("Erro ao gerar token: %v", err)
	}
	fmt.Println("Token gerado:", tokenAdmin)

	router.DELETE("/delete/:id", middlewares.AuthenticateAdmin, routes.DeletePostHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")

	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
