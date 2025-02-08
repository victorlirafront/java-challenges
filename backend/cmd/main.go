package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/models"
	"blog-api/routes"
	"blog-api/utils"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]models.Login{}

func hasPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte([]byte(password)))
	return err == nil
}

func generateToken(length int) string {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate token %v", err)
	}

	return base64.URLEncoding.EncodeToString(bytes)
}

func register(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Invalid method",
		})
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	if len(username) < 8 || len(password) < 8 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": "Invalid username / password",
		})
		return
	}

	if _, ok := users[username]; ok {
		c.JSON(http.StatusConflict, gin.H{
			"error": "User already exists",
		})
		return
	}

	hashedPassword, err := hasPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	fmt.Print(hashedPassword)

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func login(c *gin.Context) {
	// Obtém os parâmetros do corpo da requisição
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Verifica se o usuário existe no mapa
	user, ok := users[username]

	fmt.Print("users--aqui", users)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Verifica se a senha é válida
	if !checkPasswordHash(password, user.HashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Gerar o token de sessão e CSRF
	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	// Armazena os tokens no usuário
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user

	// Configurações do cookie
	cookieExpireDuration := 24 * time.Hour // 24 horas
	cookieSecure := true                   // Defina como true se usar HTTPS
	cookieHttpOnly := true                 // Impede o acesso via JavaScript

	// Calcula a data de expiração dos cookies
	expiration := time.Now().Add(cookieExpireDuration)

	// Define o cookie de sessão (HttpOnly para maior segurança)
	c.SetCookie("session_token", sessionToken, int(time.Until(expiration).Seconds()), "/", "", cookieSecure, cookieHttpOnly)

	// Define o cookie CSRF (não HttpOnly, pois pode ser acessado por JavaScript)
	c.SetCookie("csrf_token", csrfToken, int(time.Until(expiration).Seconds()), "/", "", cookieSecure, false)

	// Resposta de sucesso de login
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    username, // Ou qualquer outra informação relevante
	})
}

func logout(c *gin.Context) {
	// Adicione a lógica de logout aqui
}

func protected(c *gin.Context) {
	// Adicione a lógica de acesso protegido aqui
}

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
	router.POST("/register", register)  // Rota de registro
	router.POST("/login", login)        // Rota de login
	router.POST("/logout", logout)      // Rota de logout
	router.GET("/protected", protected) // Rota de acesso protegido

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
	tokenAdmin, err := utils.GenerateJWT(userAdmin)
	if err != nil {
		fmt.Println("Erro ao gerar token:", err)
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
