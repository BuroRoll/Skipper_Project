package handlers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetStatus(c *gin.Context) {
	session := sessions.Default(c)
	sessionId := session.Get("user_id")

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": sessionId,
	})
}
