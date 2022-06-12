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
	downPrice := 0
	highPrice := 999_999_999
	downRating := 0
	highRating := 5
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
		case "down_price":
			downPrice, _ = strconv.Atoi(queryValue)
			break
		case "high_price":
			highPrice, _ = strconv.Atoi(queryValue)
			break
		case "down_rating":
			downRating, _ = strconv.Atoi(queryValue)
			break
		case "high_rating":
			highRating, _ = strconv.Atoi(queryValue)
			break
		}
	}
	return models.Pagination{
		Limit:      limit,
		Page:       page,
		Sort:       sort,
		Search:     search,
		DownPrice:  downPrice,
		HighPrice:  highPrice,
		DownRating: downRating,
		HighRating: highRating,
	}

}
