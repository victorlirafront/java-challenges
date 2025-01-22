package main

import (
	"blog-api/database" // Importando o pacote
	"fmt"
)

func main() {
	db := database.CreateDatabaseConnection()
	defer db.Close()

	fmt.Println("Conexão estabelecida com sucesso!")
}
