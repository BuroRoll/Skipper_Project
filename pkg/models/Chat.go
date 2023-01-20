package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Body       string
	ChatID     string `gorm:"OnDelete:CASCADE;"`
	SenderID   string
	ReceiverID string
	IsRead     bool
}

type Chat struct {
	gorm.Model
	Sender        User `gorm:"foreignKey:SenderID"`
	SenderID      string
	Receiver      User `gorm:"foreignKey:ReceiverID"`
	ReceiverID    string
	LastMessage   Message `gorm:"foreignKey:LastMessageId"`
	LastMessageId string
}
