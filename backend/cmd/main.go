package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/models"
	"blog-api/routes"
	"blog-api/utils"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var users = map[string]models.Login{}
var AuthError = errors.New("Unauthorized")

func hasPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte([]byte(password)))
	return err == nil
}

func removePadding(token string) string {
	token = strings.ReplaceAll(token, "%3D", "=")
	return token
}

func Authorize(c *gin.Context) error {
	username := c.DefaultPostForm("username", "") // Para POST, ou c.DefaultQuery("username", "") para GET
	if username == "" {
		// Adiciona log para verificar se o username está vazio
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username is missing"})
		return AuthError
	}

	user, ok := users[username]
	if !ok {
		// Adiciona log para verificar se o usuário não foi encontrado
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return AuthError
	}

	// Obtém o cookie de sessão
	sessionToken, err := c.Cookie("session_token")
	if err != nil || sessionToken == "" || sessionToken != user.SessionToken {
		// Adiciona log para verificar se o cookie de sessão não está presente ou é inválido
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
		return AuthError
	}

	// Obtém o token CSRF dos cabeçalhos
	csrfToken := c.GetHeader("X-CSRF-Token")

	// Remover o padding dos tokens
	csrfToken = removePadding(csrfToken)
	userToken := removePadding(user.CSRFToken)

	fmt.Println("csrfToken:", csrfToken)
	fmt.Println("userToken:", userToken)

	// Comparar os tokens CSRF (com e sem padding)
	if csrfToken != userToken || csrfToken == "" {
		// Adiciona log para verificar se o token CSRF está ausente ou não corresponde
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing CSRF token"})
		return AuthError
	}

	// Se passou todas as verificações, retorna nil (autorizado)
	return nil
}

func register(c *gin.Context) {
	// Acessando o db do contexto, mas agora como *sql.DB
	db := c.MustGet("db").(*sql.DB)

	username := c.PostForm("username")
	password := c.PostForm("password")

	// Verificar se o nome de usuário e a senha têm pelo menos 8 caracteres
	if len(username) < 8 || len(password) < 8 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": "Username and password must be at least 8 characters long",
		})
		return
	}

	// Verifica se o nome de usuário já existe no banco de dados
	var user models.User
	result := db.QueryRow("SELECT * FROM users WHERE username = ?", username)
	err := result.Scan(&user.ID, &user.Username, &user.HashedPassword)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "User already exists",
		})
		return
	}

	// Gerar a senha criptografada
	hashedPassword, err := hasPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Criar um novo usuário no banco de dados usando Query ou Exec (sem GORM)
	_, err = db.Exec("INSERT INTO users (username, hashed_password) VALUES (?, ?)", username, hashedPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func login(c *gin.Context) {
	// Obtém os parâmetros do corpo da requisição
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println("username", username)
	fmt.Println("password", password)

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
	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	// Armazena os tokens no usuário
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user

	// Configurações do cookie
	cookieExpireDuration := 24 * time.Hour // 24 horas
	cookieSecure := false                  // Defina como true se usar HTTPS
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
	// Chama a função Authorize para verificar a autenticação e CSRF
	if err := Authorize(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Limpa os cookies de sessão e CSRF
	c.SetCookie("session_token", "", -1, "/", "", true, true) // Cookie com HttpOnly
	c.SetCookie("csrf_token", "", -1, "/", "", true, false)   // Cookie sem HttpOnly

	// Limpa os tokens do usuário no banco de dados
	username := c.DefaultPostForm("username", "")
	user, _ := users[username]
	user.SessionToken = ""
	user.CSRFToken = ""
	users[username] = user

	// Retorna resposta de sucesso
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func protected(c *gin.Context) {
	// Chama a função Authorize para verificar a autenticação e CSRF
	if err := Authorize(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Obtém o parâmetro "username" do formulário ou query string, dependendo da requisição
	username := c.DefaultPostForm("username", "")
	if username == "" {
		username = c.DefaultQuery("username", "") // Usado se for uma requisição GET
	}

	// Retorna uma mensagem de sucesso
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("CSRF validation successful! Welcome, %s", username),
	})
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
	router.POST("/register", register)   // Rota de registro
	router.POST("/login", login)         // Rota de login
	router.POST("/logout", logout)       // Rota de logout
	router.POST("/protected", protected) // Rota de acesso protegido

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
