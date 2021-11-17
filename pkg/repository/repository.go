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
	GetUser(email, password string) (uint, error)
}

type UserData interface {
	GetUserById(userId uint) (models.User, error)
}

type Catalog interface {
	CreateMainCatalog(name string) (uint, error)
	CreateChildCatalog(name string, parentId *uint) (uint, error)
	GetCatalog() string
	GetMainCatalog() string
}

type Repository struct {
	Authorization
	UserData
	Catalog
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserData:      NewUserDataPostgres(db),
		Catalog:       NewCatalogPostgres(db),
	}
}
