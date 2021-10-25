package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetStatus(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	fmt.Println("UserId:", userId)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
	})
	return
}
