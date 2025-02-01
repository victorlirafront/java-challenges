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

	// Recupera o ID do post a ser deletado
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Verifica se o post existe antes de tentar deletá-lo
	var existsCheck int
	err = sqlDb.QueryRow("SELECT COUNT(1) FROM posts WHERE id = ?", postID).Scan(&existsCheck)
	if err != nil || existsCheck == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post não encontrado"})
		return
	}

	// Deleta o post
	query := "DELETE FROM posts WHERE id = ?"
	_, err = sqlDb.Exec(query, postID)
	if err != nil {
		log.Printf("Erro ao deletar post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar post"})
		return
	}

	// Retorna a resposta de sucesso
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Post com ID %d deletado com sucesso", postID),
	})
}
