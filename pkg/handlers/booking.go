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
	err := h.services.ChangeStatusBookingClass(newStatus, bookingId)
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
