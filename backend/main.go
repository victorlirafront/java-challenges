package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/routes"
	"database/sql"
	"fmt"
	"log"

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
	router.GET("/posts/:id", middlewares.AuthenticateRead, routes.GetPostByIDHandler(db))
	router.POST("/posts", middlewares.AuthenticateAdmin(), routes.CreatePostHandler(db))
	router.PATCH("/posts/:id", middlewares.AuthenticateAdmin(), routes.CallUpdatePost)
	router.DELETE("/delete/:id", middlewares.AuthenticateAdmin(), routes.DeletePostHandler)

	// Search route with authentication
	router.GET("/search", middlewares.AuthenticateRead, middlewares.RoleMiddleware("regular", "admin"), routes.SearchPostsHandler(db))
}

func main() {
	db, err := database.CreateDatabaseConnection()

	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	defer db.Close()

	router := createRouter(db)

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
