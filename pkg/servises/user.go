package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"encoding/json"
	"errors"
	"math"
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

func (u UserDataService) GetUserStatistic(userId uint, userStatus string) (models.Statistic, error) {
	var stat models.Statistic
	if userStatus == "mentor" {
		stat = u.GetMentorStatistic(userId)
	} else if userStatus == "menti" {
		stat = u.GetMentiStatistic(userId)
	}
	return stat, nil
}

func (u UserDataService) GetMentiStatistic(userId uint) models.Statistic {
	countLessons := u.repo.GetMentiCountLessons(userId, "full", true)

	countCompletedLessons := u.repo.GetMentiCountLessons(userId, "", true)
	countLastMonthCompletedLessons := u.repo.GetMentiCountLessons(userId, "last_month", true)
	countLastThreeMonthCompletedLessons := u.repo.GetMentiCountLessons(userId, "last_three_month", true)

	countLastMonthUnclompletedLessons := u.repo.GetMentiCountLessons(userId, "last_month", false)
	countLastThreeMonthUnclompletedLessons := u.repo.GetMentiCountLessons(userId, "last_three_month", false)
	countUncomplitedLessons := u.repo.GetMentiCountLessons(userId, "", false)

	var fullAttendance = 0.0
	var lastMonthAttendance = 0.0
	var lastThreeMonthAttendance = 0.0
	if countLessons != 0 {
		fullAttendance = math.Round(countCompletedLessons / countLessons * 100)
	}
	if sum := countLastMonthCompletedLessons + countLastMonthUnclompletedLessons; sum != 0 {
		lastMonthAttendance = math.Round(countLastMonthCompletedLessons / (sum) * 100)
	}
	if sum := countLastThreeMonthCompletedLessons + countLastThreeMonthUnclompletedLessons; sum != 0 {
		lastThreeMonthAttendance = math.Round(countLastThreeMonthCompletedLessons / (sum) * 100)
	}
	stat := models.Statistic{
		LessonsCount:                      countCompletedLessons,
		StudentsCount:                     0,
		LastMonthLessonsCount:             countLastMonthCompletedLessons,
		LastThreeMonthsLessonsCount:       countLastThreeMonthCompletedLessons,
		LastMonthUnclompletedLessons:      countLastMonthUnclompletedLessons,
		LastThreeMonthUnclompletedLessons: countLastThreeMonthUnclompletedLessons,
		UncomplitedLessons:                countUncomplitedLessons,
		FullAttendance:                    fullAttendance,
		LastMonthAttendance:               lastMonthAttendance,
		LastThreeMonthAttendance:          lastThreeMonthAttendance,
	}
	return stat
}

func (u UserDataService) GetMentorStatistic(userId uint) models.Statistic {
	countCompletedStudents := u.repo.GetMentorCountStudents(userId)
	countLessons := u.repo.GetMentorCountLessons(userId, "full", true)

	countCompletedLessons := u.repo.GetMentorCountLessons(userId, "", true)
	countLastMonthCompletedLessons := u.repo.GetMentorCountLessons(userId, "last_month", true)
	countLastThreeMonthCompletedLessons := u.repo.GetMentorCountLessons(userId, "last_three_month", true)

	countLastMonthUnclompletedLessons := u.repo.GetMentorCountLessons(userId, "last_month", false)
	countLastThreeMonthUnclompletedLessons := u.repo.GetMentorCountLessons(userId, "last_three_month", false)
	countUncomplitedLessons := u.repo.GetMentorCountLessons(userId, "", false)

	var fullAttendance = 0.0
	var lastMonthAttendance = 0.0
	var lastThreeMonthAttendance = 0.0
	if countLessons != 0 {
		fullAttendance = math.Round(countCompletedLessons / countLessons * 100)
	}
	if sum := countLastMonthCompletedLessons + countLastMonthUnclompletedLessons; sum != 0 {
		lastMonthAttendance = math.Round(countLastMonthCompletedLessons / (sum) * 100)
	}
	if sum := countLastThreeMonthCompletedLessons + countLastThreeMonthUnclompletedLessons; sum != 0 {
		lastThreeMonthAttendance = math.Round(countLastThreeMonthCompletedLessons / (sum) * 100)
	}
	stat := models.Statistic{
		LessonsCount:                      countCompletedLessons,
		StudentsCount:                     countCompletedStudents,
		LastMonthLessonsCount:             countLastMonthCompletedLessons,
		LastThreeMonthsLessonsCount:       countLastThreeMonthCompletedLessons,
		LastMonthUnclompletedLessons:      countLastMonthUnclompletedLessons,
		LastThreeMonthUnclompletedLessons: countLastThreeMonthUnclompletedLessons,
		UncomplitedLessons:                countUncomplitedLessons,
		FullAttendance:                    fullAttendance,
		LastMonthAttendance:               lastMonthAttendance,
		LastThreeMonthAttendance:          lastThreeMonthAttendance,
	}
	return stat
}

func (u UserDataService) ChangePassword(userId uint, oldPassword string, newPassword string) error {
	hashPassword := generatePasswordHash(oldPassword)
	user, _ := u.repo.GetUserById(userId)
	if user.Password != hashPassword {
		return errors.New("Неверный старый пароль ")
	}
	newHashPassword := generatePasswordHash(newPassword)
	err := u.repo.ChangePassword(user, newHashPassword)
	return err
}

func (u UserDataService) AddUserToFavourite(userId uint, UserToFavourite uint, status string) error {
	user, _ := u.repo.GetUserById(userId)
	favUser, _ := u.repo.GetUserById(UserToFavourite)
	if status == "mentor" {
		err := u.repo.AddFavouriteMentor(user, favUser)
		return err
	} else if status == "menti" {
		err := u.repo.AddFavouriteMenti(user, favUser)
		return err
	} else {
		return errors.New("Такого типа пользователей не существует ")
	}
}

type FavouriteUsers struct {
	Id             uint   `json:"id"`
	FirstName      string `json:"FirstName"`
	SecondName     string `json:"SecondName"`
	Description    string `json:"description"`
	Specialization string `json:"specialization"`
	ProfilePicture string `json:"profilePicture"`
}

func (u UserDataService) GetFavourites(userId uint, status string) ([]FavouriteUsers, error) {
	var err error
	var users []models.User
	if status == "mentor" {
		users, err = u.repo.GetFavouriteMentors(userId)
	} else if status == "menti" {
		users, err = u.repo.GetFavouriteMentis(userId)
	} else {
		return nil, errors.New("Такого типа пользователей не существует ")
	}
	if err != nil {
		return nil, err
	}
	jsonUsers, _ := json.Marshal(users)
	var fUsers []FavouriteUsers
	err = json.Unmarshal(jsonUsers, &fUsers)
	return fUsers, nil
}
