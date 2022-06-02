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

func (u UserDataService) CreateUserCommunication(input forms.UserCommunicationInput, userId uint) (uint, error) {
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

func (u UserDataService) CreateUserEducation(education forms.UserEducationInput, userId uint) (uint, error) {
	return u.repo.CreateUserEducation(education, userId)
}

func (u UserDataService) CreateUserWorkExperience(workExperience forms.UserWorkExperience, userId uint) (uint, error) {
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

func (u UserDataService) SetUserEmail(email string, userId uint) error {
	return u.repo.SetUserEmail(email, userId)
}

func (u UserDataService) UpdateMentorSpecialization(specialization string, userId uint) error {
	return u.repo.UpdateMentorSpecialization(specialization, userId)
}

func (u UserDataService) AddUserOtherInfo(data string, userId uint) (uint, error) {
	return u.repo.AddUserOtherInfo(data, userId)
}

func (u UserDataService) GetUserOtherInfo(userId uint) (string, error) {
	userOtherInfo, err := u.repo.GetUserOtherInfo(userId)
	if err != nil {
		return "", err
	}
	jsonData, err := json.Marshal(userOtherInfo)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (u UserDataService) DeleteUserCommunication(communicationId string) error {
	return u.repo.DeleteUserCommunication(communicationId)
}

func (u UserDataService) DeleteUserEducation(educationId string) error {
	return u.repo.DeleteUserEducation(educationId)
}

func (u UserDataService) DeleteUserWorkExperience(workExperienceId string) error {
	return u.repo.DeleteUserWorkExperience(workExperienceId)
}

func (u UserDataService) DeleteUserOtherInfo(otherInfoId string) error {
	return u.repo.DeleteUserOtherInfo(otherInfoId)
}

func (u UserDataService) GetUnreadMessagesCount(userId uint) uint {
	unreadMessagesCount := u.repo.GetUnreadMessagesCount(userId)
	return unreadMessagesCount.Count
}
