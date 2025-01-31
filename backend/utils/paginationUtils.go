package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	page := 1
	perPage := 100000

	// Obtém o valor do parâmetro "page" da query string
	pageQuery := c.DefaultQuery("page", "1") // "1" é o valor padrão, caso "page" não esteja presente
	if pageQuery != "" {
		page, _ = strconv.Atoi(pageQuery) // Converte para inteiro, ignorando erro
	}

	// Obtém o valor do parâmetro "limit" da query string
	pagesLimit := c.DefaultQuery("limit", "100000") // "100000" é o valor padrão, caso "limit" não esteja presente
	if pagesLimit != "" {
		// Converte "pagesLimit" para inteiro, ignorando erro
		perPage, _ = strconv.Atoi(pagesLimit)
	}

	return page, perPage
}
