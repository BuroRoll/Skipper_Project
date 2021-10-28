package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.SignUpUserForm) (uint, error)
	CreateMentorUser(user models.SignUpMentorForm) (uint, error)
	UpgradeUserToMentor(userId uint, formData models.SignUpUserToMentorForm) error
	GetUser(login, password string) (uint, error)
	GenerateToken(login, password string) (string, string, error)
	GenerateTokenByID(userId uint) (string, string, error)
	ParseToken(token string) (uint, error)
	ParseRefreshToken(token string) (uint, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
