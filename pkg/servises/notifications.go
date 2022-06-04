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

type TimeChangeData struct {
	ClassId    uint
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
}

func (n NotificationsService) CreateClassTimeChangeNotification(user models.User, classId uint, receiver uint) string {
	data := TimeChangeData{
		ClassId:    classId,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
	}
	jsonData, _ := json.Marshal(data)
	notification := models.ClassNotification{
		Receiver: receiver,
		Type:     "time change",
		IsRead:   false,
		Data:     string(jsonData),
	}
	notification = n.repo.CreateClassNotification(notification)
	jsonNotification, _ := json.Marshal(notification)
	return string(jsonNotification)
}

type StatusChangeData struct {
	ClassId    uint   `json:"class_id"`
	ClassName  string `json:"class_name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	NewStatus  string `json:"new_status"`
	OldStatus  string `json:"old_status"`
}

func (n NotificationsService) CreateBookingStatusChangeNotification(bookingUsersData repository.BookingUsers, userId uint, newStatus string, oldStatus string, notificationType string) (string, uint) {
	data := StatusChangeData{
		ClassId:   bookingUsersData.Id,
		ClassName: bookingUsersData.ClassDataName,
		NewStatus: newStatus,
		OldStatus: oldStatus,
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
		Type:     notificationType,
		IsRead:   false,
		Data:     string(jsonData),
		Receiver: receiverId,
	}
	notification = n.repo.CreateClassNotification(notification)
	jsonNotification, _ := json.Marshal(notification)
	return string(jsonNotification), receiverId
}
