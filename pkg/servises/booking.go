package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"encoding/json"
	"errors"
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

func (b BookingService) CheckBookingCommunications(userCommunications []models.Communication, communicationId uint) error {
	messengerId := b.repo.GetMessengerByCommunication(communicationId)
	for _, i := range userCommunications {
		for _, j := range i.Messenger {
			if j.ID == messengerId {
				return nil
			}
		}
	}
	return errors.New("user have not communication")
}

func (b BookingService) GetClassTimeMask(classId string) (string, error) {
	classTimeMask, err := b.repo.GetClassTimeMask(classId)
	if err != nil {
		return "", err
	}
	return classTimeMask.Time, nil
}

func (b BookingService) GetClassTime(classId string) (string, error) {
	classTime, err := b.repo.GetClassTime(classId)
	if err != nil {
		return "", err
	}
	var classTimes []string
	for _, j := range classTime {
		classTimes = append(classTimes, j.Time)
	}
	jsonClassTime, _ := json.Marshal(classTimes)
	return string(jsonClassTime), nil
}

func (b BookingService) ChangeBookingTime(newBookingTime forms.ChangeBookingTimeInput, userId uint) (error, models.User) {
	err := b.repo.ChangeBookingTime(newBookingTime.ClassId, newBookingTime.Time)
	userData := b.repo.GetReceiverName(userId)
	if err != nil {
		return err, userData
	}
	return nil, userData
}
