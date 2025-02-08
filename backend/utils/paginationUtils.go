package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPaginationParams(c *gin.Context) (int, int) {
	page := 1
	perPage := 100000

	pageQuery := c.DefaultQuery("page", "1")
	if pageQuery != "" {
		page, _ = strconv.Atoi(pageQuery)
	}

	pagesLimit := c.DefaultQuery("limit", "100000")
	if pagesLimit != "" {
		perPage, _ = strconv.Atoi(pagesLimit)
	}

	return page, perPage
}
