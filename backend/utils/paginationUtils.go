package utils

import (
	"net/http"
	"strconv"
)

func GetPaginationParams(r *http.Request) (int, int) {
	page := 1
	perPage := 100000

	pageQuery := r.URL.Query().Get("page")
	if pageQuery != "" {
		page, _ = strconv.Atoi(pageQuery)
	}

	perPageQuery := r.URL.Query().Get("per_page")
	if perPageQuery != "" {
		perPage, _ = strconv.Atoi(perPageQuery)
	}

	return page, perPage
}
