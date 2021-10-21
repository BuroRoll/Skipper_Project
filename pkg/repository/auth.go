package repository

import (
	"Skipper/pkg/models"
	"errors"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetUser(email, password string) (uint, error) {
	var user models.User
	result := r.db.Where("email=? AND password=?", email, password).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user.ID, gorm.ErrRecordNotFound
	}
	return user.ID, nil
}

func (r *AuthPostgres) CreateUser(user models.User) (uint, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return user.ID, result.Error
	}
	return user.ID, nil
}
