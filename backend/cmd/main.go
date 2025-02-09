package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/routes"
	"blog-api/utils"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func createRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.DatabaseMiddleware(db))
	router.Use(middlewares.CORSMiddleware())
	registerRoutes(router, db)

	return router
}

func registerRoutes(router *gin.Engine, db *sql.DB) {
	// User authentication and registration
	router.POST("/register", routes.Register)
	router.POST("/login", routes.Login)
	router.POST("/logout", routes.Logout)
	router.POST("/protected", routes.Protected)

	// Posts routes
	router.GET("/posts", routes.GetPostsHandler(db))
	router.GET("/posts/:id", routes.GetPostByIDHandler(db))
	router.POST("/posts", middlewares.AuthenticateAdmin, routes.CreatePostHandler(db))
	router.PATCH("/posts/:id", middlewares.AuthenticateAdmin, routes.CallUpdatePost)
	router.DELETE("/delete/:id", middlewares.AuthenticateAdmin, routes.DeletePostHandler)

	// Search route with authentication
	router.GET("/search", middlewares.AuthenticateRead, middlewares.RoleMiddleware("regular", "admin"), routes.SearchPostsHandler(db))
}

func generateAdminToken() (string, error) {
	userAdmin := os.Getenv("REGULAR_USER_ID")
	if userAdmin == "" {
		return "", fmt.Errorf("REGULAR_USER_ID não está definido no ambiente")
	}

	tokenAdmin, err := utils.GenerateJWT(userAdmin)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token: %v", err)
	}

	return tokenAdmin, nil
}

func main() {
	db, err := database.CreateDatabaseConnection()

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	defer db.Close()

	tokenAdmin, err := generateAdminToken()

	if err != nil {
		log.Fatalf("Erro ao gerar token de admin: %v", err)
	}

	fmt.Println("Token gerado:", tokenAdmin)

	router := createRouter(db)

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
