package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"encoding/json"
)

type UserDataService struct {
	repo repository.UserData
}

func NewUserDataService(repo repository.UserData) *UserDataService {
	return &UserDataService{repo: repo}
}

func (u UserDataService) GetUserData(userId uint) (models.User, error) {
	return u.repo.GetUserById(userId)
}

func (u UserDataService) GetUserCommunications(userId uint) ([]models.Communication, error) {
	return u.repo.GetUserCommunications(userId)
}

func (u UserDataService) GetMessengers() ([]models.Messenger, error) {
	return u.repo.GetMessengers()
}

func (u UserDataService) CreateUserCommunication(input forms.UserCommunicationInput, userId uint) error {
	return u.repo.CreateUserCommunication(input, userId)
}

func (u UserDataService) UpdateBaseProfileData(input forms.UpdateBaseProfileData, userId uint) error {
	return u.repo.UpdateBaseProfileData(input, userId)
}

func (u UserDataService) UpdateProfilePicture(filename string, userId uint) error {
	return u.repo.UpdateProfilePicture(filename, userId)
}

func (u UserDataService) GetUserEducation(userId uint) (string, error) {
	userEducation, err := u.repo.GetUserEducation(userId)
	if err != nil {
		return "", err
	}
	jsonData, err := json.Marshal(userEducation)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (u UserDataService) CreateUserEducation(education forms.UserEducationInput, userId uint) error {
	return u.repo.CreateUserEducation(education, userId)
}

func (u UserDataService) CreateUserWorkExperience(workExperience forms.UserWorkExperience, userId uint) error {
	return u.repo.CreateUserWorkExperience(workExperience, userId)
}

func (u UserDataService) GetUserWorkExperience(userId uint) (string, error) {
	userEducation, err := u.repo.GetUserWorkExperience(userId)
	if err != nil {
		return "", err
	}
	jsonData, err := json.Marshal(userEducation)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
