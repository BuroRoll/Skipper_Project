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
	GetUserByEmailOrPhone(login string) (models.User, error)
	VerifyEmail(userId uint) error
	ChangeUserPassword(user models.User, newPassword string) error
}

type UserData interface {
	GetUserById(userId uint) (models.User, error)
	GetUserCommunications(userId uint) ([]models.Communication, error)
	GetMessengers() ([]models.Messenger, error)
	CreateUserCommunication(input forms.UserCommunicationInput, userId uint) (uint, error)
	UpdateBaseProfileData(input forms.UpdateBaseProfileData, userId uint) error
	UpdateProfilePicture(filename string, userId uint) error
	GetUserEducation(userId uint) ([]models.Education, error)
	CreateUserEducation(education forms.UserEducationInput, userId uint) (uint, error)
	CreateUserWorkExperience(workExperience forms.UserWorkExperience, userId uint) (uint, error)
	GetUserWorkExperience(userId uint) ([]models.WorkExperience, error)
	SetUserEmail(email string, userId uint) error
	UpdateMentorSpecialization(specialization string, userId uint) error
	AddUserOtherInfo(data string, userId uint) (uint, error)
	GetUserOtherInfo(userId uint) ([]models.OtherInformation, error)
	DeleteUserCommunication(communicationId string) error
	DeleteUserEducation(educationId string) error
	DeleteUserWorkExperience(workExperienceId string) error
	DeleteUserOtherInfo(otherInfoId string) error
	GetUnreadMessagesCount(userId uint) unreadMessagesCounter
	GetMentorCountStudents(userId uint) uint
	GetMentorCountLessons(userId uint, time string, isComplete bool) float64
	GetMentiCountLessons(userId uint, time string, isComplete bool) float64
	ChangePassword(user models.User, newPassword string) error
	AddFavouriteMentor(user models.User, favUser models.User) error
	AddFavouriteMenti(user models.User, favUser models.User) error
	GetFavouriteMentors(userId uint) ([]models.User, error)
	GetFavouriteMentis(userId uint) ([]models.User, error)
	DeleteFavouriteMentor(user models.User, favUser models.User) error
	DeleteFavouriteMenti(user models.User, favUser models.User) error
	CheckFavouriteUser(user models.User, favUser models.User) bool
}

type Catalog interface {
	GetCatalog() []models.Catalog0
	GetMainCatalog() []MainCatalog
	GetCatalogChild() []models.Catalog3
	GetClasses(**models.Pagination) ([]models.User, error)
}

type Classes interface {
	CreateUserClasses(input models.Class) (uint, error)
	CreateTheoreticClass(input models.TheoreticClass) (uint, models.TheoreticClass, error)
	CreatePracticClass(input models.PracticClass) (uint, models.PracticClass, error)
	CreateKeyClass(input models.KeyClass) (uint, models.KeyClass, error)
	GetCatalogTags(catalogId uint) (models.Catalog3, error)
	GetUserClasses(userId uint) ([]models.Class, error)
	DeleteClass(classId string) error
	DeleteTheoreticClass(classId string) error
	DeletePracticClass(classId string) error
	DeleteKeyClass(classId string) error
	UpdateClass(classData models.Class, classId uint) error
	UpdateTheoreticClass(classData models.TheoreticClass, classId uint) (models.TheoreticClass, error)
	UpdatePracticClass(classData models.PracticClass, classId uint) (models.PracticClass, error)
	UpdateKeyClass(classData models.KeyClass, classId uint) (models.KeyClass, error)
	GetClassById(classId string) (models.Class, error)
	CalcAverageClassPrice(classId uint)
}
type Booking interface {
	BookingClass(data forms.BookingClassInput, mentiId uint) error
	GetBookingsToMe(mentorId uint, status string) ([]UserBooking, error)
	GetMyBookings(mentiId uint, status string) ([]UserBooking, error)
	ChangeStatusBookingClass(newStatus string, bookingClassId string) error
	GetMessengerByCommunication(id uint) uint
	GetClassTimeMask(classId string) (BookingTimeMask, error)
	GetClassTime(classId string) ([]ClassTime, error)
	ChangeBookingTime(classId uint, time []string) error
	GetReceiverName(userId uint) models.User
	GetBookingUsersById(bookingId string) BookingUsers
	GetBookingStatus(bookingId uint) Status
	GetBookingById(bookingId uint) models.UserClass
	ChangeBookingCommunication(bookingId uint, communicationId uint) error
	SetBookingUnsuccess(bookingId uint) error
}

type Chat interface {
	CreateMessage(input forms.MessageInput) (models.Message, error)
	GetOpenChats(userId uint) ([]Chats, error)
	GetChatData(userId string, receiverID string) (models.Chat, []models.Message, error)
	ReadMessages(chatId string, userId string) error
}

type Comments interface {
	CreateComment(comment forms.CommentInput) error
	CreateLessonComment(LessonComment models.LessonComment) error
	GetComments(userId uint) ([]CommentData, error)
	CalcRating(userId uint)
}

type Notifications interface {
	GetAllClassNotifications(userId string) []models.ClassNotification
	CreateClassNotification(notification models.ClassNotification) models.ClassNotification
	ReadNotification(notificationId uint) error
	DeleteNotification(notificationId uint) error
}

type Repository struct {
	Authorization
	UserData
	Catalog
	Classes
	Booking
	Chat
	Comments
	Notifications
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserData:      NewUserDataPostgres(db),
		Catalog:       NewCatalogPostgres(db),
		Classes:       NewClassesPostgres(db),
		Booking:       NewBookingPostgres(db),
		Chat:          NewChatPostgres(db),
		Comments:      NewCommentsPostgres(db),
		Notifications: NewNotificationsPostgres(db),
	}
}
