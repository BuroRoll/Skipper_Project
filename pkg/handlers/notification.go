package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ClassDataNotification struct {
	Data   string
	UserId uint
}

func (h *Handler) SendClassNotification(c *gin.Context) {
	var data ClassDataNotification
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	SendClassNotification(data.Data, strconv.Itoa(int(data.UserId)))
}

func (h *Handler) GetAllClassNotifications(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	//userId := c.Query("user_id")
	classNotifications := h.services.GetAllClassNotifications(userId.(string))
	c.JSON(http.StatusOK, classNotifications)
}
