package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/repository"
)

type UserDataService struct {
	repo repository.UserData
}

func (u UserDataService) GetUserData(userId uint) (models.User, error) {
	return u.repo.GetUserById(userId)
}

func NewUserDataService(repo repository.UserData) *UserDataService {
	return &UserDataService{repo: repo}
}
