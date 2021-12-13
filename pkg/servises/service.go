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
	GetUser(login, password string) (uint, bool, error)
	GenerateToken(login, password string) (string, string, error)
	GenerateTokenByID(userId uint, isMentor bool) (string, string, error)
	SaveProfilePicture(file multipart.File, filename string) (string, error)
	ParseToken(token string) (uint, bool, error)
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
	UpdateProfilePicture(filename string, userId uint) error
	GetUserEducation(userId uint) (string, error)
	CreateUserEducation(education forms.UserEducationInput, userId uint) error
	CreateUserWorkExperience(workExperience forms.UserWorkExperience, userId uint) error
	GetUserWorkExperience(userId uint) (string, error)
	SetUserEmail(email string, userId uint) error
	UpdateMentorSpecialization(specialization string, userId uint) error
	AddUserOtherInfo(data string, userId uint) error
	GetUserOtherInfo(userId uint) (string, error)
}

type Catalog interface {
	CreateCatalog(catalog forms.CatalogInput) (uint, error)
	GetCatalog() string
	GetMainCatalog() string
	GetCatalogChild() (string, error)
}

type Class interface {
	CreateUserClass(class forms.ClassesInput, userId uint) (uint, error)
	CreateTheoreticClass(class forms.TheoreticClassInput, userId uint) error
	CreatePracticClass(class forms.PracticClassInput, userId uint) error
	CreateKeyClass(class forms.KeyClass, userId uint) error
	GetUserClasses(userId uint) (string, error)
}

type Service struct {
	Authorization
	UserData
	Catalog
	Class
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UserData:      NewUserDataService(repos.UserData),
		Catalog:       NewCatalogService(repos.Catalog),
		Class:         NewClassesService(repos.Classes),
	}
}
