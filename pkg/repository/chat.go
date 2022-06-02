package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"fmt"
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
	var chat models.Chat
	c.db.First(&chat, input.ChatID)
	chat.LastMessage = message
	c.db.Save(&chat)
	chat.LastMessageId = strconv.Itoa(int(chat.LastMessage.ID))
	c.db.Save(&chat)
	fmt.Println(chat.LastMessage)
	return message, nil
}

type Chats struct {
	gorm.Model
	Sender              models.User `gorm:"foreignkey:SenderID"`
	SenderID            string
	Receiver            models.User `gorm:"foreignkey:ReceiverID"`
	ReceiverID          string
	LastMessage         models.Message `gorm:"foreignKey:ChatID;references:ID"`
	CountUnreadMessages uint           `json:"count_unread_messages"`
}

func (c ChatPostgres) GetOpenChats(userId uint) ([]Chats, error) {
	var chats []Chats
	c.db.
		Preload("Sender", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, first_name, second_name, profile_picture")
		}).
		Preload("Receiver", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, first_name, second_name, profile_picture")
		}).
		Preload("LastMessage").
		Select("*").
		Joins("INNER JOIN (SELECT created_at AS last_message_time, id AS message_id FROM messages) AS message_data ON message_data.message_id::varchar(255) = chats.last_message_id").
		Joins("LEFT JOIN (SELECT chat_id, COUNT(*) as count_unread_messages FROM messages WHERE (receiver_id in (?) AND is_read IS false) GROUP BY chat_id) AS d ON chat_id = chats.id", strconv.Itoa(int(userId))).
		Where("sender_id = ?", userId).
		Or("receiver_id = ?", userId).
		Order("last_message_time DESC").
		Find(&chats)
	return chats, nil
}

func (c ChatPostgres) GetChatData(userId string, receiverID string) (models.Chat, []models.Message, error) {
	var chat models.Chat
	if err := c.db.Where(
		"sender_id = ? AND receiver_id = ?", userId, receiverID).Or(
		"sender_id = ? AND receiver_id = ?", receiverID, userId).
		Preload("Sender", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, first_name, second_name, profile_picture")
		}).
		Preload("Receiver", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id, first_name, second_name, profile_picture")
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
	c.db.
		Where("chat_id = ?", chatId).
		Order("created_at").
		Find(&messages)
	return chat, messages, nil
}

func (c ChatPostgres) ReadMessages(chatId string, userId string) error {
	var message models.Message
	c.db.
		Model(&message).
		Where("chat_id in (?) and receiver_id in (?)", chatId, userId).
		UpdateColumn("is_read", true)
	return nil
}
