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
