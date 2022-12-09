package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) ReportUser(c *gin.Context) {
	userId := c.GetUint(userCtx)
	var report forms.ReportUser
	if err := c.BindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	err := h.services.MakeReport(userId, report)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка отправки жалобы"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
