package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetChatsList(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	chats, err := h.services.GetOpenChats(userId.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить список чатов"})
		return
	}
	c.JSON(http.StatusOK, chats)
}

func (h *Handler) GetChatMessages(c *gin.Context) {
	receiverID := c.Param("userID")
	userId, _ := c.Get(userCtx)
	chatData, messagesList, err := h.services.GetChatData(userId.(uint), receiverID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить сообщения"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"chat": chatData, "messages": messagesList})
}
