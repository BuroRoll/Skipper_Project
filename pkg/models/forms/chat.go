package forms

type MessageInput struct {
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Message    string `json:"message"`
	ChatID     string `json:"chatId"`
}

type ReadChatInput struct {
	ChatId uint `json:"chatId"`
	UserId uint `json:"user_id"`
}
