package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func CreateDatabaseConnection() (*sql.DB, error) {
	// Carregar o arquivo .env apenas no ambiente local
	err := godotenv.Load()
	if err != nil {
		log.Printf("Aviso: Não foi possível carregar o arquivo .env: %v", err)
	}

	// Ler as variáveis de ambiente
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	// Verificar se todas as variáveis de ambiente estão configuradas
	if host == "" || port == "" || username == "" || password == "" || database == "" {
		return nil, fmt.Errorf("uma ou mais variáveis de ambiente não estão configuradas corretamente")
	}

	// Construir string de conexão (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a conexão: %v", err)
	}

	// Testar a conexão com o banco de dados
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Conexão bem-sucedida ao banco de dados MySQL!")
	return db, nil
}
