package repository

import (
	"Skipper/pkg/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.SignUpUserForm) (uint, error)
	CreateMentor(user models.SignUpMentorForm) (uint, error)
	UpgradeUserToMentor(userId uint, form models.SignUpUserToMentorForm) error
	GetUser(email, password string) (uint, error)
	GetUserById(userId uint) (models.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
