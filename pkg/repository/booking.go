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
	booking.Price90_5 = data.Price90_1

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

	Time []models.BookingTime `gorm:"foreignKey:BookingClassID;references:ID"`

	Communication uint
	First_name    string `json:"menti_first_name"`
	Second_name   string `json:"menti_second_name"`
	User_time     string `json:"user_time"`
}

func (b BookingPostgres) GetBookingsToMe(mentorId uint, status string) ([]UserBooking, error) {
	var bookings []UserBooking

	b.db.
		Unscoped().
		Preload("Time").
		Table("user_classes").
		Select("*").
		Where("user_id=? AND status = ?", mentorId, status).
		Joins("LEFT JOIN (select id as user_data_id, first_name, second_name, time as user_time from Users) AS menti_data ON user_classes.menti_id = menti_data.user_data_id").
		//"LEFT JOIN (select id as communication_id, ").

		//Joins("left join (select id, login from Communications) AS menti_communications ON user_classes.menti_id = menti_communications.id").
		Find(&bookings)
	return bookings, nil
}
