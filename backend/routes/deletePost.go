package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeletePostHandler(c *gin.Context) {
	// Recupera o db do contexto da requisição
	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Banco de dados não disponível"})
		return
	}

	// Faz um tipo de assert para garantir que seja o tipo correto (*sql.DB)
	sqlDb, ok := db.(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Banco de dados inválido"})
		return
	}

	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	query := "DELETE FROM posts WHERE id = ?"
	result, err := sqlDb.Exec(query, postID)
	if err != nil {
		log.Printf("Erro ao deletar post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar post"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Erro ao verificar linhas afetadas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post não encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Post com ID %d deletado com sucesso", postID),
	})
}
