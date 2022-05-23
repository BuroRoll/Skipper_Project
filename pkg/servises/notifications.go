package service

import (
	"Skipper/pkg/repository"
	"encoding/json"
)

type NotificationsService struct {
	repo repository.Notifications
}

func NewNotificationsService(repo repository.Notifications) *NotificationsService {
	return &NotificationsService{repo: repo}
}

func (n NotificationsService) GetAllClassNotifications(userId string) string {
	classesNotifications := n.repo.GetAllClassNotifications(userId)
	jsonClassesNotifications, _ := json.Marshal(classesNotifications)
	return string(jsonClassesNotifications)
}
