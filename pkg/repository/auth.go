package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"errors"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetUser(login, password string) (uint, error) {
	var user models.User
	result := r.db.Where("email=? AND password=? OR phone=? AND password=?", login, password, login, password).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user.ID, gorm.ErrRecordNotFound
	}
	return user.ID, nil
}

func (r *AuthPostgres) CreateUser(user_register forms.SignUpUserForm) (uint, error) {
	var user models.User
	user = models.User{
		Phone:      user_register.Phone,
		FirstName:  user_register.FirstName,
		SecondName: user_register.SecondName,
		Password:   user_register.Password,
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return user.ID, result.Error
	}
	return user.ID, nil
}

func (r *AuthPostgres) CreateMentor(mentor_register forms.SignUpMentorForm, profilePicturePath string) (uint, error) {
	var user models.User
	user = models.User{
		Phone:          mentor_register.Phone,
		Password:       mentor_register.Password,
		FirstName:      mentor_register.FirstName,
		SecondName:     mentor_register.SecondName,
		Description:    mentor_register.Description,
		Specialization: mentor_register.Specialization,
		Time:           mentor_register.Time,
		IsMentor:       true,
		ProfilePicture: profilePicturePath,
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return user.ID, result.Error
	}
	return user.ID, nil
}

func (r *AuthPostgres) UpgradeUserToMentor(userId uint, registerData forms.SignUpUserToMentorForm) error {
	var user models.User
	result := r.db.Where("id=?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	user.Specialization = registerData.Specialization
	user.Description = registerData.Description
	user.Time = registerData.Time
	user.IsMentor = true
	r.db.Save(user)
	return nil
}
