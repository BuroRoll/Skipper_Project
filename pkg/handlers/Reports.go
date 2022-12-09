package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ReportUser
// @Description  Жалоба на пользователя
// @Tags 		 reports
// @Accept       json
// @Produce      json
// @Param 		 request 	body 		forms.ReportUser	true 	"query params"
// @Success      200  		{object} 	forms.SuccessResponse
// @Router       /api/reports/ [post]
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
