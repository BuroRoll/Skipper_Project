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

func (r *AuthPostgres) GetUser(login, password string) (uint, bool, error) {
	var user models.User
	result := r.db.Where("email=? AND password=? OR phone=? AND password=?", login, password, login, password).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user.ID, false, gorm.ErrRecordNotFound
	}
	return user.ID, user.IsMentor, nil
}

func (r *AuthPostgres) GetUserById(userId uint) (models.User, error) {
	var user models.User
	result := r.db.Where("id=?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (r *AuthPostgres) CreateUser(userRegister forms.SignUpUserForm) (uint, error) {
	var user models.User
	user = models.User{
		Phone:      userRegister.Phone,
		FirstName:  userRegister.FirstName,
		SecondName: userRegister.SecondName,
		Password:   userRegister.Password,
	}
	result := r.db.Create(&user)
	if result.Error != nil {
		return user.ID, result.Error
	}
	return user.ID, nil
}

func (r *AuthPostgres) CreateMentor(mentorRegister forms.SignUpMentorForm, profilePicturePath string) (uint, error) {
	var user models.User
	user = models.User{
		Phone:          mentorRegister.Phone,
		Password:       mentorRegister.Password,
		FirstName:      mentorRegister.FirstName,
		SecondName:     mentorRegister.SecondName,
		Description:    mentorRegister.Description,
		Specialization: mentorRegister.Specialization,
		Time:           mentorRegister.Time,
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

func (r *AuthPostgres) VerifyEmail(userId uint) error {
	var user models.User
	result := r.db.Where("id=?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	user.IsVerifyEmail = true
	r.db.Save(user)
	return nil
}

func (r *AuthPostgres) GetUserByEmailOrPhone(login string) (models.User, error) {
	var user models.User
	result := r.db.Where("email=? or phone = ?", login, login).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (r *AuthPostgres) ChangeUserPassword(user models.User, newPassword string) error {
	err := r.db.Model(&user).Update("password", newPassword)
	return err.Error
}
