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
	classNotifications := h.services.GetAllClassNotifications(strconv.Itoa(int(userId.(uint))))
	c.JSON(http.StatusOK, classNotifications)
}

func (h *Handler) ReadNotification(c *gin.Context) {
	notificationIdQuery := c.Query("notification_id")
	notificationId, _ := strconv.ParseUint(notificationIdQuery, 10, 64)
	err := h.services.ReadNotification(uint(notificationId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось пометить уведомление как прочитанное"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) DeleteNotification(c *gin.Context) {
	notificationIdQuery := c.Query("notification_id")
	notificationId, _ := strconv.ParseUint(notificationIdQuery, 10, 64)
	err := h.services.DeleteNotification(uint(notificationId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить уведомление"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
