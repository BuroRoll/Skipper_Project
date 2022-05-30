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

type StatusChangeData struct {
	ClassId    uint   `json:"class_id"`
	ClassName  string `json:"class_name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	NewStatus  string `json:"new_status"`
}

func (n NotificationsService) CreateBookingStatusChangeNotification(bookingUsersData repository.BookingUsers, userId uint, newStatus string) (string, uint) {
	data := StatusChangeData{
		ClassId:   bookingUsersData.Id,
		ClassName: bookingUsersData.ClassDataName,
		NewStatus: newStatus,
	}
	var receiverId uint
	if userId == bookingUsersData.MentorDataId {
		data.FirstName = bookingUsersData.MentorFirstName
		data.SecondName = bookingUsersData.MentorSecondName
		receiverId = bookingUsersData.MentiDataId
	} else {
		data.FirstName = bookingUsersData.MentiFirstName
		data.SecondName = bookingUsersData.MentiSecondName
		receiverId = bookingUsersData.MentorDataId
	}
	jsonData, _ := json.Marshal(data)
	notification := models.ClassNotification{
		Type:     "status change",
		IsRead:   false,
		Data:     string(jsonData),
		Receiver: receiverId,
	}
	n.repo.CreateNotification(notification)
	jsonNotification, _ := json.Marshal(notification)
	return string(jsonNotification), receiverId
}
