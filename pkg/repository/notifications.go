package repository

import (
	"Skipper/pkg/models"
	"encoding/json"
	"gorm.io/gorm"
)

type NotificationsPostgres struct {
	db *gorm.DB
}

func NewNotificationsPostgres(db *gorm.DB) *NotificationsPostgres {
	return &NotificationsPostgres{db: db}
}

func (n NotificationsPostgres) GetAllClassNotifications(userId string) []models.ClassNotification {
	var classNotifications []models.ClassNotification
	n.db.Debug().
		Where("receiver = ?", userId).Order("created_at desc").Find(&classNotifications)
	return classNotifications
}

func (n NotificationsPostgres) CreateClassTimeChangeNotification(classId uint, userFirstName string, userSecondName string, receiver uint) models.ClassNotification {
	data := jsonData(userFirstName, userSecondName, classId)
	notification := models.ClassNotification{
		Receiver: receiver,
		Type:     "time change",
		IsRead:   false,
		Data:     data,
	}
	n.db.Create(&notification)
	return notification
}

func (n NotificationsPostgres) CreateNotification(notification models.ClassNotification) {
	n.db.Save(&notification)
}

func jsonData(userFirstName string, userSecondName string, classId uint) string {
	type Data struct {
		ClassId        uint
		UserFirstName  string
		UserSecondName string
	}
	data := Data{
		ClassId:        classId,
		UserFirstName:  userFirstName,
		UserSecondName: userSecondName,
	}
	jsonInfo, _ := json.Marshal(data)
	return string(jsonInfo)
}
