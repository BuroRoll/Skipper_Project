package repository

import (
	"Skipper/pkg/models"
	"errors"
	"gorm.io/gorm"
)

type UserDataPostgres struct {
	db *gorm.DB
}

func (u UserDataPostgres) GetUserById(userId uint) (models.User, error) {
	var user models.User
	result := u.db.Where("id=?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func NewUserDataPostgres(db *gorm.DB) *UserDataPostgres {
	return &UserDataPostgres{db: db}
}
