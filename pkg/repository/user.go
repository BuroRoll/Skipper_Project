package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"errors"
	"gorm.io/gorm"
)

type UserDataPostgres struct {
	db *gorm.DB
}

func NewUserDataPostgres(db *gorm.DB) *UserDataPostgres {
	return &UserDataPostgres{db: db}
}

func (u UserDataPostgres) GetUserById(userId uint) (models.User, error) {
	var user models.User
	result := u.db.Where("id=?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (u UserDataPostgres) GetUserCommunications(userId uint) ([]models.Communication, error) {
	var communications []models.Communication
	result := u.db.Where("parent_id=?", userId).
		Preload("Messenger", func(db *gorm.DB) *gorm.DB {
			return db.Select("Id, Name")
		}).Find(&communications)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return communications, nil
}

func (u UserDataPostgres) GetMessengers() ([]models.Messenger, error) {
	var messengers []models.Messenger
	result := u.db.Raw("SELECT id, name FROM messengers").Find(&messengers)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return messengers, nil
}

func (u UserDataPostgres) CreateUserCommunication(input forms.UserCommunicationInput, userId uint) error {
	var communication models.Communication
	var messenger models.Messenger
	u.db.Find(&messenger, input.MessengerId)
	communication.Login = input.Login
	communication.ParentId = userId
	communication.Messenger = append(communication.Messenger, &messenger)
	u.db.Omit("Messenger.*").Create(&communication)
	return nil
}

func (u UserDataPostgres) UpdateBaseProfileData(input forms.UpdateBaseProfileData, userId uint) error {
	var user models.User
	result := u.db.Where("id=?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	user.FirstName = input.FirstName
	user.SecondName = input.SecondName
	user.Patronymic = input.Patronymic
	user.DateOfBirthday = input.DateOfBirthday
	user.Time = input.Time
	u.db.Save(user)
	return nil
}
