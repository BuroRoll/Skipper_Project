package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateComment(c *gin.Context) {
	var comment forms.CommentInput
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	if len(comment.Text) == 0 {

	} else {
		err := h.services.CreateComment(comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить комментарий"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
