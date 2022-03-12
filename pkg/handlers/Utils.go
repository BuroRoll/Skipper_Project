package handlers

import (
	"Skipper/pkg/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 10
	page := 1
	sort := "created_at asc"
	var search []string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		case "search":
			search = strings.Split(queryValue, ",")
			break
		}
	}
	return models.Pagination{
		Limit:  limit,
		Page:   page,
		Sort:   sort,
		Search: search,
	}

}
