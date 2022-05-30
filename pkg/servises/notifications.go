package service

import (
	"Skipper/pkg/models"
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

func (n NotificationsService) CreateClassTimeChangeNotification(user models.User, classId uint, receiver uint) string {
	notification := n.repo.CreateClassTimeChangeNotification(classId, user.FirstName, user.SecondName, receiver)
	jsonNotification, _ := json.Marshal(notification)
	return string(jsonNotification)
}
