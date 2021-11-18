package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"mime/multipart"
)

type Authorization interface {
	CreateUser(user forms.SignUpUserForm) (uint, error)
	CreateMentorUser(user forms.SignUpMentorForm, profilePicturePath string) (uint, error)
	UpgradeUserToMentor(userId uint, formData forms.SignUpUserToMentorForm) error
	GetUser(login, password string) (uint, error)
	GenerateToken(login, password string) (string, string, error)
	GenerateTokenByID(userId uint) (string, string, error)
	SaveProfilePicture(file multipart.File, filename string) (string, error)
	ParseToken(token string) (uint, error)
	ParseRefreshToken(token string) (uint, error)
	SendVerifyEmail(userId uint) error
	VerifyEmail(userId uint) error
}

type UserData interface {
	GetUserData(userId uint) (models.User, error)
	GetUserCommunications(userId uint) ([]models.Communication, error)
	GetMessengers() ([]models.Messenger, error)
	CreateUserCommunication(input forms.UserCommunicationInput, userId uint) error
	UpdateBaseProfileData(input forms.UpdateBaseProfileData, userId uint) error
}

type Catalog interface {
	CreateCatalog(catalog forms.CatalogInput) (uint, error)
	GetCatalog() string
	GetMainCatalog() string
}

type Service struct {
	Authorization
	UserData
	Catalog
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UserData:      NewUserDataService(repos.UserData),
		Catalog:       NewCatalogService(repos.Catalog),
	}
}
