package repository

import (
	"Skipper/pkg/models"
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
	n.db.Where("receiver = ?", userId).Find(&classNotifications)
	return classNotifications
}
