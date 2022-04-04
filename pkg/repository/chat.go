package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"gorm.io/gorm"
	"strconv"
)

type ChatPostgres struct {
	db *gorm.DB
}

func NewChatPostgres(db *gorm.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

func (c ChatPostgres) CreateMessage(input forms.MessageInput) (models.Message, error) {
	message := models.Message{
		SenderID:   input.SenderID,
		ReceiverID: input.ReceiverID,
		ChatID:     input.ChatID,
		Body:       input.Message,
		IsRead:     false,
	}
	result := c.db.Create(&message)
	if result.Error != nil {
		return models.Message{}, result.Error
	}
	return message, nil
}

func (c ChatPostgres) GetOpenChats(userId uint) ([]models.Chat, error) {
	var chats []models.Chat
	c.db.
		Where("sender_id = ?", userId).
		Or("receiver_id = ?", userId).
		Find(&chats)
	return chats, nil
}

func (c ChatPostgres) GetChatData(userId string, receiverID string) (models.Chat, []models.Message, error) {
	var chat models.Chat
	if err := c.db.Where(
		"sender_id = ? AND receiver_id = ?", userId, receiverID).Or(
		"sender_id = ? AND receiver_id = ?", receiverID, userId).
		Preload("Sender", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, first_name, second_name")
		}).
		Preload("Receiver", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, first_name, second_name")
		}).
		First(&chat).Error; err != nil {
		chat.SenderID = userId
		chat.ReceiverID = receiverID

		var sender models.User
		c.db.First(&sender, userId)
		chat.Sender = sender

		var receiver models.User
		c.db.First(&receiver, receiverID)
		chat.Receiver = receiver

		c.db.Create(&chat)
	}

	var messages []models.Message
	chatId := strconv.FormatUint(uint64(chat.ID), 10)
	c.db.Debug().Where("chat_id = ?", chatId).Find(&messages)
	return chat, messages, nil
}
