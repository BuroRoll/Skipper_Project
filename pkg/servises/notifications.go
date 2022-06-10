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
	ChatUserId uint   `json:"chat_user_id"`
}

func (n NotificationsService) CreateClassTimeChangeNotification(user models.User, classId uint, receiver uint) string {
	data := TimeChangeData{
		ClassId:    classId,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		ChatUserId: user.ID,
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
	ChatUserId uint   `json:"chat_user_id"`
	ClassId    uint   `json:"class_id"`
	ClassName  string `json:"class_name"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	NewStatus  string `json:"new_status"`
	OldStatus  string `json:"old_status"`
	IsMentor   bool   `json:"is_mentor"`
	MentorId   uint   `json:"mentor_id"`
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
		data.ChatUserId = bookingUsersData.MentorDataId
		receiverId = bookingUsersData.MentiDataId
	} else {
		data.FirstName = bookingUsersData.MentiFirstName
		data.SecondName = bookingUsersData.MentiSecondName
		data.ChatUserId = bookingUsersData.MentiDataId
		receiverId = bookingUsersData.MentorDataId
	}
	data.MentorId = bookingUsersData.MentorDataId
	if receiverId == bookingUsersData.MentorDataId {
		data.IsMentor = true
	} else {
		data.IsMentor = false
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

type CommunicationData struct {
	ChatUserId           uint                   `json:"chat_user_id"`
	ClassId              uint                   `json:"class_id"`
	ClassName            string                 `json:"class_name"`
	FirstName            string                 `json:"first_name"`
	SecondName           string                 `json:"second_name"`
	NewCommunicationId   uint                   `json:"new_communication_id"`
	MentorCommunications []models.Communication `json:"mentor_communications"`
}

func (n NotificationsService) CreateChangeBookingCommunicationNotification(senderId uint, bookingUsers repository.BookingUsers, classId uint, newCommunicationId uint, mentorCommunications []models.Communication) (string, uint) {
	var receiverId uint
	data := CommunicationData{
		ClassId:              classId,
		ClassName:            bookingUsers.ClassDataName,
		NewCommunicationId:   newCommunicationId,
		MentorCommunications: mentorCommunications,
	}
	if senderId == bookingUsers.MentorDataId {
		data.FirstName = bookingUsers.MentorFirstName
		data.SecondName = bookingUsers.MentorSecondName
		data.ChatUserId = bookingUsers.MentorDataId
		receiverId = bookingUsers.MentiDataId
	} else {
		data.FirstName = bookingUsers.MentiFirstName
		data.SecondName = bookingUsers.MentiSecondName
		data.ChatUserId = bookingUsers.MentiDataId
		receiverId = bookingUsers.MentorDataId
	}
	jsonData, _ := json.Marshal(data)
	notification := models.ClassNotification{
		Type:     "communication change",
		IsRead:   false,
		Data:     string(jsonData),
		Receiver: receiverId,
	}
	notification = n.repo.CreateClassNotification(notification)
	jsonNotification, _ := json.Marshal(notification)
	return string(jsonNotification), receiverId
}

func (n NotificationsService) ReadNotification(notificationId uint) error {
	err := n.repo.ReadNotification(notificationId)
	return err
}

func (n NotificationsService) DeleteNotification(notificationId uint) error {
	err := n.repo.DeleteNotification(notificationId)
	return err
}
