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
	n.db.Debug().
		Where("receiver = ?", userId).
		Order("created_at desc").
		Find(&classNotifications)
	return classNotifications
}

func (n NotificationsPostgres) CreateClassNotification(notification models.ClassNotification) models.ClassNotification {
	n.db.Save(&notification)
	return notification
}
