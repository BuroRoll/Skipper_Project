package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"errors"
	"gorm.io/gorm"
)

type UserDataPostgres struct {
	db *gorm.DB
}

func NewUserDataPostgres(db *gorm.DB) *UserDataPostgres {
	return &UserDataPostgres{db: db}
}

func (u UserDataPostgres) GetUserById(userId uint) (models.User, error) {
	var user models.User
	result := u.db.First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, gorm.ErrRecordNotFound
	}
	return user, nil
}

func (u UserDataPostgres) GetUserCommunications(userId uint) ([]models.Communication, error) {
	var communications []models.Communication
	result := u.db.Where("parent_id=?", userId).
		Preload("Messenger", func(db *gorm.DB) *gorm.DB {
			return db.Select("Id, Name")
		}).Find(&communications)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return communications, nil
}

func (u UserDataPostgres) GetMessengers() ([]models.Messenger, error) {
	var messengers []models.Messenger
	result := u.db.Raw("SELECT id, name FROM messengers").Find(&messengers)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return messengers, nil
}

func (u UserDataPostgres) CreateUserCommunication(input forms.UserCommunicationInput, userId uint) (uint, error) {
	var communication models.Communication
	var messenger models.Messenger
	u.db.First(&messenger, input.MessengerId)
	communication.Login = input.Login
	communication.ParentId = userId
	communication.Messenger = append(communication.Messenger, &messenger)
	u.db.Omit("Messenger.*").Create(&communication)
	return communication.ID, nil
}

func (u UserDataPostgres) UpdateBaseProfileData(input forms.UpdateBaseProfileData, userId uint) error {
	var user models.User
	result := u.db.First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	user.FirstName = input.FirstName
	user.SecondName = input.SecondName
	user.Patronymic = input.Patronymic
	user.DateOfBirthday = input.DateOfBirthday
	user.Time = input.Time
	user.Description = input.Description
	u.db.Save(&user)
	return nil
}

func (u UserDataPostgres) UpdateProfilePicture(filename string, userId uint) error {
	var user models.User
	result := u.db.First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	user.ProfilePicture = filename
	u.db.Save(&user)
	return nil
}

func (u UserDataPostgres) GetUserEducation(userId uint) ([]models.Education, error) {
	var education []models.Education
	result := u.db.Where("parent_id=?", userId).Find(&education)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return education, nil
}

func (u UserDataPostgres) CreateUserEducation(education forms.UserEducationInput, userId uint) (uint, error) {
	userEducation := models.Education{
		ParentId:    userId,
		Institution: education.Institution,
		StartYear:   education.StartYear,
		EndYear:     education.EndYear,
		Degree:      education.Degree,
	}
	result := u.db.Create(&userEducation)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return 0, gorm.ErrRegistered
	}
	return userEducation.ID, nil
}

func (u UserDataPostgres) CreateUserWorkExperience(workExperience forms.UserWorkExperience, userId uint) (uint, error) {
	userWorkExperience := models.WorkExperience{
		ParentId:     userId,
		Organization: workExperience.Organization,
		StartYear:    workExperience.StartYear,
		EndYear:      workExperience.EndYear,
	}
	result := u.db.Create(&userWorkExperience)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return 0, gorm.ErrRegistered
	}
	return userWorkExperience.ID, nil
}

func (u UserDataPostgres) GetUserWorkExperience(userId uint) ([]models.WorkExperience, error) {
	var userWorkExperience []models.WorkExperience
	result := u.db.Where("parent_id=?", userId).Find(&userWorkExperience)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return userWorkExperience, nil
}

func (u UserDataPostgres) SetUserEmail(email string, userId uint) error {
	var user models.User
	result := u.db.First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	user.Email = email
	u.db.Save(&user)
	return nil
}

func (u UserDataPostgres) UpdateMentorSpecialization(specialization string, userId uint) error {
	var user models.User
	result := u.db.First(&user, userId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	user.Specialization = specialization
	u.db.Save(&user)
	return nil
}

func (u UserDataPostgres) AddUserOtherInfo(data string, userId uint) (uint, error) {
	otherInfo := models.OtherInformation{
		ParentId: userId,
		Data:     data,
	}
	result := u.db.Create(&otherInfo)
	if errors.Is(result.Error, gorm.ErrRegistered) {
		return 0, gorm.ErrRegistered
	}
	return otherInfo.ID, nil
}

func (u UserDataPostgres) GetUserOtherInfo(userId uint) ([]models.OtherInformation, error) {
	var userOtherInfo []models.OtherInformation
	result := u.db.Where("parent_id=?", userId).Find(&userOtherInfo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return userOtherInfo, nil
}

func (u UserDataPostgres) DeleteUserCommunication(communicationId string) error {
	result := u.db.Exec("DELETE FROM messenger_communication WHERE communication_id = (?)", communicationId)
	result = u.db.Exec("DELETE FROM communications WHERE id = (?)", communicationId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (u UserDataPostgres) DeleteUserEducation(educationId string) error {
	result := u.db.Exec("DELETE FROM educations WHERE id = (?)", educationId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (u UserDataPostgres) DeleteUserWorkExperience(workExperienceId string) error {
	result := u.db.Exec("DELETE FROM work_experiences WHERE id = (?)", workExperienceId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (u UserDataPostgres) DeleteUserOtherInfo(otherInfoId string) error {
	result := u.db.Exec("DELETE FROM other_informations WHERE id = ?", otherInfoId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return gorm.ErrRecordNotFound
	}
	return nil
}
