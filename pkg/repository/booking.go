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

	booking.Time = data.Time

	booking.Communication = data.Communication

	b.db.Save(&booking)
	return nil
}
