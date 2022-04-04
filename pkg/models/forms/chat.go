package forms

type MessageInput struct {
	SenderID   string `json:"senderId"`
	ReceiverID string `json:"receiverId"`
	Message    string `json:"message"`
	ChatID     string `json:"chatId"`
}
