package service

import (
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"encoding/json"
)

type BookingService struct {
	repo repository.Booking
}

func NewBookingService(repo repository.Booking) *BookingService {
	return &BookingService{repo: repo}
}

func (b BookingService) BookingClass(classForm forms.BookingClassInput, mentiId uint) error {
	return b.repo.BookingClass(classForm, mentiId)
}

func (b BookingService) GetBookingsToMe(mentorId uint, status string) (string, error) {
	considerationsList, err := b.repo.GetBookingsToMe(mentorId, status)

	if err != nil {
		return "", err
	}
	jsonConsiderationList, _ := json.Marshal(considerationsList)
	return string(jsonConsiderationList), nil
}

func (b BookingService) GetMyBookings(mentiId uint, status string) (string, error) {
	considerationsList, err := b.repo.GetMyBookings(mentiId, status)

	if err != nil {
		return "", err
	}
	jsonConsiderationList, _ := json.Marshal(considerationsList)
	return string(jsonConsiderationList), nil
}

func (b BookingService) ChangeStatusBookingClass(newStatus string, bookingClassId string) error {
	err := b.repo.ChangeStatusBookingClass(newStatus, bookingClassId)
	if err != nil {
		return err
	}
	return nil
}
