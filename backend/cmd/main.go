package main

import (
	"blog-api/database"
	"blog-api/middlewares"
	"blog-api/models"
	"blog-api/routes"
	"blog-api/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hasPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte([]byte(password)))
	return err == nil
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
	// Obtém o db do contexto
	db := c.MustGet("db").(*sql.DB)

	// Obtém os parâmetros do corpo da requisição
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Verifica se o nome de usuário e a senha têm pelo menos 8 caracteres
	if len(username) < 8 || len(password) < 8 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"error": "Username and password must be at least 8 characters long",
		})
		return
	}

	// Recupera o usuário do banco de dados
	var user models.User
	result := db.QueryRow("SELECT id, username, hashed_password FROM users WHERE username = ?", username)
	err := result.Scan(&user.ID, &user.Username, &user.HashedPassword)
	if err == sql.ErrNoRows {
		// Se o usuário não for encontrado no banco, retorna erro
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	} else if err != nil {
		// Se houve algum outro erro ao consultar o banco
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while accessing database",
		})
		return
	}

	// Verifica se a senha fornecida corresponde ao hash armazenado
	if !checkPasswordHash(password, user.HashedPassword) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	// Gerar o token de sessão e CSRF
	sessionToken := utils.GenerateToken(32)
	csrfToken := utils.GenerateToken(32)

	// Atualizar o usuário com os tokens gerados
	_, err = db.Exec("UPDATE users SET session_token = ?, csrf_token = ? WHERE id = ?", sessionToken, csrfToken, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user tokens",
		})
		return
	}

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
		"user":    user.Username,
	})
}

func logout(c *gin.Context) {
	// Chama a função Authorize para verificar a autenticação e CSRF
	if err := middlewares.SessionAuthMiddleware(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	// Obtém o parâmetro "username" do formulário
	username := c.DefaultPostForm("username", "")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username is required",
		})
		return
	}

	// Acessa o banco de dados para atualizar o usuário, limpando os tokens
	db := c.MustGet("db").(*sql.DB)

	// Atualiza os valores dos tokens para nulos no banco de dados
	_, err := db.Exec("UPDATE users SET session_token = ?, csrf_token = ? WHERE username = ?", "", "", username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while logging out",
		})
		return
	}

	// Limpa os cookies de sessão e CSRF no cliente
	c.SetCookie("session_token", "", -1, "/", "", true, true) // Cookie com HttpOnly
	c.SetCookie("csrf_token", "", -1, "/", "", true, false)   // Cookie sem HttpOnly

	// Retorna resposta de sucesso
	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

func protected(c *gin.Context) {
	// Chama a função Authorize para verificar a autenticação e CSRF
	if err := middlewares.SessionAuthMiddleware(c); err != nil {
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

	// Verifica se o parâmetro username está presente
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username is required",
		})
		return
	}

	// Recupera o session_token e csrf_token dos cookies
	sessionToken, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Session token not found",
		})
		return
	}

	csrfToken, err := c.Cookie("csrf_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "CSRF token not found",
		})
		return
	}

	// Acessa o banco de dados para verificar o usuário com base no sessionToken e csrfToken
	db := c.MustGet("db").(*sql.DB)

	var user models.User
	err = db.QueryRow("SELECT id, username, session_token, csrf_token FROM users WHERE session_token = ? AND csrf_token = ?", sessionToken, csrfToken).Scan(&user.ID, &user.Username, &user.SessionToken, &user.CSRFToken)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid session or CSRF token",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while accessing database",
			})
		}
		return
	}

	// Retorna uma mensagem de sucesso se a validação for bem-sucedida
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
