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
	ResetPassword(login string) error
	SetNewPassword(userId uint, newPassword string) error
}

type UserData interface {
	GetUserData(userId uint) (models.User, error)
	GetUserCommunications(userId uint) ([]models.Communication, error)
	GetMessengers() ([]models.Messenger, error)
	CreateUserCommunication(input forms.UserCommunicationInput, userId uint) (uint, error)
	UpdateBaseProfileData(input forms.UpdateBaseProfileData, userId uint) error
	UpdateProfilePicture(filename string, userId uint) error
	GetUserEducation(userId uint) (string, error)
	CreateUserEducation(education forms.UserEducationInput, userId uint) (uint, error)
	CreateUserWorkExperience(workExperience forms.UserWorkExperience, userId uint) (uint, error)
	GetUserWorkExperience(userId uint) (string, error)
	SetUserEmail(email string, userId uint) error
	UpdateMentorSpecialization(specialization string, userId uint) error
	AddUserOtherInfo(data string, userId uint) (uint, error)
	GetUserOtherInfo(userId uint) (string, error)
	DeleteUserCommunication(communicationId string) error
	DeleteUserEducation(educationId string) error
	DeleteUserWorkExperience(workExperienceId string) error
	DeleteUserOtherInfo(otherInfoId string) error
	GetUnreadMessagesCount(userId uint) uint
	GetUserStatistic(userId uint, userStatus string) (models.Statistic, error)
	ChangePassword(userId uint, oldPassword string, newPassword string) error
	AddUserToFavourite(userId uint, UserToFavourite uint, status string) error
	GetFavourites(userId uint, status string) ([]FavouriteUsers, error)
	DeleteFavourite(userId uint, favUserId uint, status string) error
	IsFavouriteUser(userId uint, favUser uint) bool
}

type Catalog interface {
	GetCatalog() string
	GetMainCatalog() string
	GetCatalogChild() (string, error)
	GetClasses(*models.Pagination) (string, error)
	GetClassesWithFavourite(userId uint, pagination *models.Pagination) (string, error)
}

type Class interface {
	CreateUserClass(class forms.ClassesInput, userId uint) (uint, error)
	CreateTheoreticClass(class forms.TheoreticClassInput, userId uint) (uint, error)
	CreatePracticClass(class forms.PracticClassInput, userId uint) (uint, error)
	CreateKeyClass(class forms.KeyClass, userId uint) (uint, error)
	GetUserClasses(userId uint) (string, error)
	DeleteClass(classId string) error
	DeleteTheoreticClass(classId string) error
	DeletePracticClass(classId string) error
	DeleteKeyClass(classId string) error
	UpdateClass(input forms.UpdateClassesInput) error
	UpdateTheoreticClass(input forms.UpdateSubclassInput) error
	UpdatePracticClass(input forms.UpdateSubclassInput) error
	UpdateKeyClass(input forms.UpdateKeyClassInput) error
	GetClassById(classId string) (string, string, string, error)
}

type Booking interface {
	BookingClass(bookingClassData forms.BookingClassInput, mentiId uint) error
	GetBookingsToMe(mentorId uint, status string) (string, error)
	GetMyBookings(mentorId uint, status string) (string, error)
	ChangeStatusBookingClass(newStatus string, bookingClassId string) error
	CheckBookingCommunications(userCommunications []models.Communication, communicationId uint) error
	GetClassTimeMask(classId string) (string, error)
	GetClassTime(classId string) (string, error)
	ChangeBookingTime(newBookingTime forms.ChangeBookingTimeInput, userId uint) (error, models.User)
	GetBookingUsersById(bookingId string) repository.BookingUsers
	GetBookingStatus(bookingId uint) string
	ChangeBookingCommunication(newBookingCommunicationInput forms.ChangeBookingCommunication) error
	SetBookingUnsuccess(bookingId uint) error
}

type Chat interface {
	CreateMessage(messageInput forms.MessageInput) (string, error)
	GetOpenChats(userId uint) (string, error)
	GetChatData(userId uint, receiverID string) (string, string, error)
	ReadMessages(chatId string, userId uint) error
}

type Comments interface {
	CreateComment(comment forms.CommentInput) error
	GetComments(userId uint) (string, error)
	CreateLessonComment(lessonComment forms.CommentInput) error
}

type Notifications interface {
	GetAllClassNotifications(userId string) string
	CreateClassTimeChangeNotification(user models.User, classId uint, receiver uint) string
	CreateBookingStatusChangeNotification(bookingUsersData repository.BookingUsers, userId uint, newStatus string, oldStatus string, notificationType string) (string, uint)
	CreateChangeBookingCommunicationNotification(
		senderId uint,
		bookingUsers repository.BookingUsers,
		classId uint, newCommunicationId uint,
		mentorCommunications []models.Communication) (string, uint)
	ReadNotification(notificationId uint) error
	DeleteNotification(notificationId uint) error
}

type Reports interface {
	MakeReport(userId uint, reportForm forms.ReportUser) error
}

type Service struct {
	Authorization
	UserData
	Catalog
	Class
	Booking
	Chat
	Comments
	Notifications
	Reports
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UserData:      NewUserDataService(repos.UserData),
		Catalog:       NewCatalogService(repos.Catalog),
		Class:         NewClassesService(repos.Classes),
		Booking:       NewBookingService(repos.Booking),
		Chat:          NewChatService(repos.Chat),
		Comments:      NewCommentsService(repos.Comments),
		Notifications: NewNotificationsService(repos.Notifications),
		Reports:       NewReportsService(repos.Reports),
	}
}
