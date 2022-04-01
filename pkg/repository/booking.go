package repository

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"gorm.io/gorm"
)

type BookingPostgres struct {
	db *gorm.DB
}

func NewBookingPostgres(db *gorm.DB) *BookingPostgres {
	return &BookingPostgres{db: db}
}

func (b BookingPostgres) BookingClass(data forms.BookingClassInput, mentiId uint) error {
	var booking models.UserClass
	var class models.Class
	b.db.First(&class, data.ClassId)

	booking.Class = class
	booking.UserID = data.MentorId
	booking.ClassID = data.ClassId
	booking.ClassType = data.ClassType
	booking.MentiId = mentiId

	booking.Duration15 = data.Duration15
	booking.Price15 = data.Price15

	booking.Duration30_1 = data.Duration30_1
	booking.Price30_1 = data.Price30_1
	booking.Duration30_3 = data.Duration30_3
	booking.Price30_3 = data.Price30_3
	booking.Duration30_5 = data.Duration30_5
	booking.Price30_5 = data.Price30_5

	booking.Duration60_1 = data.Duration60_1
	booking.Price60_1 = data.Price60_1
	booking.Duration60_3 = data.Duration60_3
	booking.Price60_3 = data.Price60_3
	booking.Duration60_5 = data.Duration60_5
	booking.Price60_5 = data.Price60_5

	booking.Duration90_1 = data.Duration90_1
	booking.Price90_1 = data.Price90_1
	booking.Duration90_3 = data.Duration90_3
	booking.Price90_3 = data.Price90_3
	booking.Duration90_5 = data.Duration90_5
	booking.Price90_5 = data.Price90_5

	booking.PriceFullTime = data.PriceFullTime
	booking.FullTime = data.FullTime

	booking.Communication = data.Communication

	for _, t := range data.Time {
		booking.Time = append(booking.Time, models.BookingTime{
			Time: t,
		})
	}
	b.db.Save(&booking)
	return nil
}

type UserBooking struct {
	ID        uint
	UserID    uint
	ClassID   uint
	ClassType string
	Status    string
	MentiId   uint
	Class     models.Class

	Duration15   bool
	Price15      uint
	Duration30_1 bool
	Price30_1    uint
	Duration30_3 bool
	Price30_3    uint
	Duration30_5 bool
	Price30_5    uint

	Duration60_1 bool
	Price60_1    uint
	Duration60_3 bool
	Price60_3    uint
	Duration60_5 bool
	Price60_5    uint

	Duration90_1 bool
	Price90_1    uint
	Duration90_3 bool
	Price90_3    uint
	Duration90_5 bool
	Price90_5    uint

	FullTime      bool
	PriceFullTime uint

	Time []models.BookingTime `gorm:"foreignKey:BookingClassID;references:ID"`
	//Tags []models.Catalog3

	Communication       uint
	First_name          string `json:"menti_first_name"`
	Second_name         string `json:"menti_second_name"`
	User_time           string `json:"user_time"`
	Communication_login string `json:"communication_login"`
	Messenger_name      string `json:"messenger_name"`
}

func (b BookingPostgres) GetBookingsToMe(mentorId uint, status string) ([]UserBooking, error) {
	var bookings []UserBooking

	b.db.
		Unscoped().
		Table("user_classes").
		Preload("Class.Tags").
		Preload("Time").
		Select("*").
		Where("user_id=? AND status = ? AND new_data.messenger_id = messenger_data.messenger_id", mentorId, status).
		Joins("LEFT JOIN (select id as user_data_id, first_name, second_name, time as user_time from Users) AS menti_data ON user_classes.menti_id = menti_data.user_data_id").
		Joins("LEFT JOIN (SELECT messenger_id, communication_id as cmc_id FROM messenger_communication) AS messager_communications ON messager_communications.cmc_id = communication").
		//Joins("LEFT JOIN (SELECT id AS messenger_id, name AS messenger_name FROM messengers) AS messenger_data ON messenger_id = communication").
		Joins("LEFT JOIN (SELECT id AS messenger_id, name AS messenger_name FROM messengers) AS messenger_data ON messager_communications.messenger_id = messenger_data.messenger_id").
		Joins("LEFT JOIN (SELECT parent_id, login AS communication_login, id AS communication_id FROM communications) AS communication_data ON parent_id = menti_data.user_data_id").
		Joins("LEFT JOIN (SELECT communication_id AS n_id, messenger_id FROM messenger_communication) AS new_data ON new_data.n_id = communication_id").
		Find(&bookings)
	return bookings, nil
}

func (b BookingPostgres) GetMyBookings(mentiId uint, status string) ([]UserBooking, error) {
	var bookings []UserBooking
	b.db.
		Unscoped().
		Table("user_classes").
		Preload("Class.Tags").
		Preload("Time").
		Select("*").
		Where("menti_id=? AND status = ?", mentiId, status).
		Joins("LEFT JOIN (select id as user_data_id, first_name, second_name, time as user_time from Users) AS mentor_data ON user_classes.user_id = mentor_data.user_data_id").
		//Joins("LEFT JOIN (SELECT id AS messenger_id, name AS messenger_name FROM messengers) AS messenger_data ON messenger_id = communication").
		Joins("LEFT JOIN (SELECT messenger_id, communication_id as cmc_id FROM messenger_communication) AS messager_communications ON messager_communications.cmc_id = communication").
		Joins("LEFT JOIN (SELECT id AS messenger_id, name AS messenger_name FROM messengers) AS messenger_data ON messager_communications.messenger_id = messenger_data.messenger_id").
		Find(&bookings)
	return bookings, nil
}

func (b BookingPostgres) ChangeStatusBookingClass(newStatus string, bookingClassId string) error {
	err := b.db.Model(&models.UserClass{}).Where("id = ?", bookingClassId).Update("status", newStatus)
	if err != nil {
		return err.Error
	}
	return nil
}

type messenger_communication struct {
	MessengerId     uint `json:"messenger_id"`
	CommunicationId uint `json:"communication_id"`
}

func (b BookingPostgres) GetMessengerByCommunication(id uint) uint {
	var data messenger_communication
	b.db.
		Table("messenger_communication").
		Where("communication_id = ?", id).
		Find(&data)
	return data.MessengerId
}
