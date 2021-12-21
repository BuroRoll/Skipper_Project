package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user forms.SignUpUserForm) (uint, error)
	CreateMentor(user forms.SignUpMentorForm, profilePicturePath string) (uint, error)
	UpgradeUserToMentor(userId uint, form forms.SignUpUserToMentorForm) error
	GetUser(email, password string) (uint, bool, error)
	GetUserById(userId uint) (models.User, error)
	VerifyEmail(userId uint) error
}

type UserData interface {
	GetUserById(userId uint) (models.User, error)
	GetUserCommunications(userId uint) ([]models.Communication, error)
	GetMessengers() ([]models.Messenger, error)
	CreateUserCommunication(input forms.UserCommunicationInput, userId uint) error
	UpdateBaseProfileData(input forms.UpdateBaseProfileData, userId uint) error
	UpdateProfilePicture(filename string, userId uint) error
	GetUserEducation(userId uint) ([]models.Education, error)
	CreateUserEducation(education forms.UserEducationInput, userId uint) error
	CreateUserWorkExperience(workExperience forms.UserWorkExperience, userId uint) error
	GetUserWorkExperience(userId uint) ([]models.WorkExperience, error)
	SetUserEmail(email string, userId uint) error
	UpdateMentorSpecialization(specialization string, userId uint) error
	AddUserOtherInfo(data string, userId uint) error
	GetUserOtherInfo(userId uint) ([]models.OtherInformation, error)
}

type Catalog interface {
	CreateMainCatalog(name string) (uint, error)
	CreateChildCatalog(name string, parentId *uint) (uint, error)
	GetCatalog() string
	GetMainCatalog() string
	GetCatalogChild() []models.Catalog3
	GetClasses() ([]models.User, error)
}

type Classes interface {
	CreateUserClasses(input models.Class) (uint, error)
	CreateTheoreticClass(input models.TheoreticClass) (uint, error)
	CreatePracticClass(input models.PracticClass) (uint, error)
	CreateKeyClass(input models.KeyClass) (uint, error)
	GetCatalogTags(catalogId uint) (models.Catalog3, error)
	GetUserClasses(userId uint) ([]models.Class, error)
	DeleteClass(classId string) error
	DeleteTheoreticClass(classId string) error
	DeletePracticClass(classId string) error
	DeleteKeyClass(classId string) error
	UpdateClass(classData models.Class, classId uint) error
	UpdateTheoreticClass(classData models.TheoreticClass, classId uint) error
	UpdatePracticClass(classData models.PracticClass, classId uint) error
	UpdateKeyClass(classData models.KeyClass, classId uint) error
}

type Repository struct {
	Authorization
	UserData
	Catalog
	Classes
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserData:      NewUserDataPostgres(db),
		Catalog:       NewCatalogPostgres(db),
		Classes:       NewClassesPostgres(db),
	}
}
