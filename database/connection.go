package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importa o driver MySQL
	"github.com/joho/godotenv"
)

// No Go, funções e variáveis são exportáveis quando seus nomes começam com uma letra maiúscula.
// Isso permite que outros pacotes acessem essas funções ou variáveis.
func CreateDatabaseConnection() *sql.DB {
	err := godotenv.Load("../config/.env")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão: %v", err)
	}

	// Testar a conexão
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Conexão bem-sucedida ao banco de dados MySQL!")
	return db
}
