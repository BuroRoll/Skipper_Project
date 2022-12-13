package handlers

import (
	"Skipper/pkg/models/forms"
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

// DeleteChat
// @Description  Удаление чата
// @Tags 		 chats
// @Accept       json
// @Produce      json
// @Param        chatID   path      int  true  "chatID"
// @Success      200  		{object} 	forms.SuccessResponse
// @Router       /api/chat/:chatID [delete]
func (h *Handler) DeleteChat(c *gin.Context) {
	chatIDForm := forms.DeleteChatForm{}
	if err := c.ShouldBindUri(&chatIDForm); err != nil {
		c.JSON(http.StatusBadRequest, forms.ErrorResponse{Error: "Неверная форма данных"})
		return
	}
	err := h.services.DeleteChat(chatIDForm.ChatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, forms.ErrorResponse{Error: "Не удалось удалить чат"})
		return
	}
	c.JSON(http.StatusOK, forms.SuccessResponse{Status: "Чат успешно удалён"})
}
