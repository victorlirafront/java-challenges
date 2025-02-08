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
	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Banco de dados não disponível"})
		return
	}

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

	var existsCheck int
	err = sqlDb.QueryRow("SELECT COUNT(1) FROM posts WHERE id = ?", postID).Scan(&existsCheck)
	if err != nil || existsCheck == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post não encontrado"})
		return
	}

	query := "DELETE FROM posts WHERE id = ?"
	_, err = sqlDb.Exec(query, postID)
	if err != nil {
		log.Printf("Erro ao deletar post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Post com ID %d deletado com sucesso", postID),
	})
}
