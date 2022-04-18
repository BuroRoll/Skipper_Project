package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Body       string
	ChatID     string
	SenderID   string
	ReceiverID string
	IsRead     bool
}

type Chat struct {
	gorm.Model
	Sender      User `gorm:"foreignkey:SenderID"`
	SenderID    string
	Receiver    User `gorm:"foreignkey:ReceiverID"`
	ReceiverID  string
	LastMessage Message
}
