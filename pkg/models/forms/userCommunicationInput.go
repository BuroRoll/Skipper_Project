package forms

type UserCommunicationInput struct {
	MessengerId uint   `json:"messenger_id" binding:"required"`
	Login       string `json:"login" binding:"required"`
}
