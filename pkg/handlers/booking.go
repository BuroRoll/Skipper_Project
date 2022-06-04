package handlers

import (
	"Skipper/pkg/models/forms"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetClassData(c *gin.Context) {
	classId := c.Query("classId")
	mentorId := c.Query("mentorId")
	tc, pc, kc, err := h.services.GetClassById(classId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Занятие не найдено"})
		return
	}
	mentor, _ := strconv.ParseUint(mentorId, 10, 64)
	mentorCommuntications, err := h.services.GetUserCommunications(uint(mentor))
	c.JSON(http.StatusOK, gin.H{
		"theoretic_class":       tc,
		"practic_class":         pc,
		"key_class":             kc,
		"mentor_communications": mentorCommuntications,
	})
}

func (h *Handler) BookClass(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	var bookingInput forms.BookingClassInput
	if err := c.BindJSON(&bookingInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма записи на занятие"})
		return
	}
	userCommunications, _ := h.services.GetUserCommunications(userId.(uint))
	err := h.services.CheckBookingCommunications(userCommunications, bookingInput.Communication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "У пользователя нет способа связи с ментором"})
		return
	}
	err = h.services.BookingClass(bookingInput, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка записи на занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) GetBookingsToMe(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	status := c.Query("status")
	Bookings, err := h.services.GetBookingsToMe(userId.(uint), status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения списка занятий"})
		return
	}
	c.JSON(http.StatusOK, Bookings)
}

func (h *Handler) ChangeStatusBookingClass(c *gin.Context) {
	bookingId := c.Query("booking_id")
	newStatus := c.Query("new_status")
	userId, _ := c.Get(userCtx)
	booking_id, _ := strconv.ParseUint(bookingId, 10, 64)
	status := h.services.GetBookingStatus(uint(booking_id))
	bookingUsersData := h.services.GetBookingUsersById(bookingId)
	if status == "canceled" && userId != bookingUsersData.MentiDataId {
		bookingChangeStatusNotification, receiverId := h.services.CreateBookingStatusChangeNotification(bookingUsersData, userId.(uint), newStatus, status, "offer to change status")
		SendClassNotification(bookingChangeStatusNotification, strconv.Itoa(int(receiverId)))
		c.JSON(http.StatusOK, gin.H{"status": "Уведомление отправлено менти"})
		return
	}
	bookingChangeStatusNotification, receiverId := h.services.CreateBookingStatusChangeNotification(bookingUsersData, userId.(uint), newStatus, status, "status change")
	err := h.services.ChangeStatusBookingClass(newStatus, bookingId)
	SendClassNotification(bookingChangeStatusNotification, strconv.Itoa(int(receiverId)))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка смены статуса занятия"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) GetMyBookings(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	status := c.Query("status")
	bookings, err := h.services.GetMyBookings(userId.(uint), status)
	fmt.Println(bookings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения списка занятий"})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

func (h *Handler) GetBookingTimes(c *gin.Context) {
	bookingClassId := c.Param("booking_class_id")
	classMaskTime, err := h.services.GetClassTimeMask(bookingClassId)
	classTime, err := h.services.GetClassTime(bookingClassId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения времени занятия"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"class_time_mask": classMaskTime,
		"class_time":      classTime,
	})
}

func (h *Handler) ChangeBookingTimes(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	var newBookingTimeInput forms.ChangeBookingTimeInput
	if err := c.BindJSON(&newBookingTimeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма изменения времени занятия"})
		return
	}
	err, userData := h.services.ChangeBookingTime(newBookingTimeInput, userId.(uint))
	notificationData := h.services.CreateClassTimeChangeNotification(userData, newBookingTimeInput.ClassId, newBookingTimeInput.Receiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось изменить время занятия"})
		return
	}
	SendClassNotification(notificationData, strconv.Itoa(int(newBookingTimeInput.Receiver)))
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
