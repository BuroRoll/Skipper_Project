package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (uint, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
