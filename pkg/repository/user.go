package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"errors"
	"gorm.io/gorm"
	"strconv"
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

type unreadMessagesCounter struct {
	ReceiverId string `json:"receiver_id"`
	Count      uint   `json:"count"`
}

func (u UserDataPostgres) GetUnreadMessagesCount(userId uint) unreadMessagesCounter {
	var unreadMessagesCount unreadMessagesCounter
	u.db.
		Table("messages").
		Select("receiver_id, count(is_read)").
		Where("is_read = false AND receiver_id = ?", strconv.Itoa(int(userId))).
		Group("receiver_id").
		First(&unreadMessagesCount)
	return unreadMessagesCount
}

type completedStudents struct {
	UserId                 uint   `json:"user_id"`
	Status                 string `json:"status"`
	CountCompletedStudents uint   `json:"count_completed_students"`
}

func (u UserDataPostgres) GetMentorCountStudents(userId uint) uint {
	var stat completedStudents
	u.db.Select("user_id, status, COUNT(DISTINCT menti_id) AS count_completed_students").
		Table("user_classes").
		Where("user_id = ? and status = ?", userId, "completed").
		Group("user_id, status").
		Find(&stat)
	return stat.CountCompletedStudents
}

type completedLessons struct {
	UserId       uint    `json:"user_id"`
	CountLessons float64 `json:"count_lessons"`
}

func (u UserDataPostgres) GetMentorCountLessons(userId uint, time string, isComplete bool) float64 {
	var stat completedLessons
	queryBuider := u.db.Select("count(bt.is_end) as count_lessons, count(is_success) as isok, count(bt.booking_class_id) as bcid, count(uc.id) as ids, uc.user_id, count(bt.time) as ctime").
		Table("booking_times bt").
		Joins("inner join user_classes uc on uc.id = bt.booking_class_id").
		Group("uc.user_id")
	switch time {
	case "full":
		queryBuider.Where("user_id = ? and is_end = true", userId)
	case "last_month":
		queryBuider.Where("user_id = ? and is_end = true and is_success = ? and TO_DATE(left(bt.time, -2),'YYYY/MM/DD') > CURRENT_DATE - INTERVAL '30 days'", userId, isComplete)
	case "last_three_month":
		queryBuider.Where("user_id = ? and is_end = true and is_success = ? and TO_DATE(left(bt.time, -2),'YYYY/MM/DD') > CURRENT_DATE - INTERVAL '90 days'", userId, isComplete)
	default:
		queryBuider.Where("user_id = ? and is_end = true and is_success = ?", userId, isComplete)
	}
	queryBuider.Find(&stat)
	return stat.CountLessons
}

func (u UserDataPostgres) GetMentiCountLessons(userId uint, time string, isComplete bool) float64 {
	var stat completedLessons
	queryBuider := u.db.Select("count(bt.is_end) as count_lessons, count(is_success) as isok, count(bt.booking_class_id) as bcid, count(uc.id) as ids, uc.menti_id as user_id, count(bt.time) as ctime").
		Table("booking_times bt").
		Joins("inner join user_classes uc on uc.id = bt.booking_class_id").
		Group("uc.menti_id")
	switch time {
	case "full":
		queryBuider.Where("uc.menti_id = ? and is_end = true", userId)
	case "last_month":
		queryBuider.Where("uc.menti_id = ? and is_end = true and is_success = ? and TO_DATE(left(bt.time, -2),'YYYY/MM/DD') > CURRENT_DATE - INTERVAL '30 days'", userId, isComplete)
	case "last_three_month":
		queryBuider.Where("uc.menti_id = ? and is_end = true and is_success = ? and TO_DATE(left(bt.time, -2),'YYYY/MM/DD') > CURRENT_DATE - INTERVAL '90 days'", userId, isComplete)
	default:
		queryBuider.Where("uc.menti_id = ? and is_end = true and is_success = ?", userId, isComplete)
	}
	queryBuider.Find(&stat)
	return stat.CountLessons
}

func (u UserDataPostgres) ChangePassword(user models.User, newPassword string) error {
	err := u.db.Model(&user).Update("password", newPassword)
	return err.Error
}

func (u UserDataPostgres) GetFavouriteMentors(userId uint) ([]models.User, error) {
	var user models.User
	result := u.db.Preload("FavouriteMentors").First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return user.FavouriteMentors, nil
}

func (u UserDataPostgres) GetFavouriteMentis(userId uint) ([]models.User, error) {
	var user models.User
	result := u.db.Preload("FavouriteMentis").First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return user.FavouriteMentis, nil
}

func (u UserDataPostgres) AddFavouriteMentor(user models.User, favUser models.User) error {
	u.db.Model(&user).Association("FavouriteMentors").Append(&favUser)
	return nil
}

func (u UserDataPostgres) AddFavouriteMenti(user models.User, favUser models.User) error {
	u.db.Model(&user).Association("FavouriteMentis").Append(&favUser)
	return nil
}

func (u UserDataPostgres) DeleteFavouriteMentor(user models.User, favUser models.User) error {
	u.db.Model(&user).Association("FavouriteMentors").Delete(&favUser)
	return nil
}

func (u UserDataPostgres) DeleteFavouriteMenti(user models.User, favUser models.User) error {
	u.db.Model(&user).Association("FavouriteMentis").Delete(&favUser)
	return nil
}
func (u UserDataPostgres) CheckFavouriteUser(user models.User, favUser models.User) bool {
	var user1 models.User
	var user2 models.User
	ids := []uint{favUser.ID}
	u.db.Model(&user).Where("id IN ?", ids).Association("FavouriteMentors").Find(&user1)
	u.db.Model(&user).Where("id IN ?", ids).Association("FavouriteMentis").Find(&user2)
	return user1.ID != 0 || user2.ID != 0
}
