package service

import (
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"encoding/json"
	"strconv"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (c ChatService) CreateMessage(messageInput forms.MessageInput) (string, error) {
	message, err := c.repo.CreateMessage(messageInput)
	if err != nil {
		return "", err
	}
	result, _ := json.Marshal(message)
	resultString := string(result)
	return resultString, nil
}

func (c ChatService) GetOpenChats(userId uint) (string, error) {
	chatsList, err := c.repo.GetOpenChats(userId)
	if err != nil {
		return "", err
	}
	result, _ := json.Marshal(chatsList)
	resultString := string(result)
	return resultString, nil
}

func (c ChatService) GetChatData(userId uint, receiverID string) (string, string, error) {
	user_id := strconv.FormatUint(uint64(userId), 10)
	chatData, messages, err := c.repo.GetChatData(user_id, receiverID)
	if err != nil {
		return "", "", err
	}
	chatInfo, _ := json.Marshal(chatData)
	messagesList, _ := json.Marshal(messages)
	return string(chatInfo), string(messagesList), nil
}

func (c ChatService) ReadMessages(chatId string, userId uint) error {
	user_id := strconv.FormatUint(uint64(userId), 10)
	return c.repo.ReadMessages(chatId, user_id)
}

func (c ChatService) DeleteChat(chatId uint) error {
	return c.repo.DeleteChat(chatId)
}
