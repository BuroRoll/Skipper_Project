package service

import (
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
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
