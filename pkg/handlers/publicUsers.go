package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetMentorData(c *gin.Context) {
	mentorId := parseId(c.Param("id"))
	user, err := h.services.GetUserData(mentorId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения данных о пользователе"})
		return
	}
	if !user.IsMentor {
		c.JSON(http.StatusBadRequest, gin.H{"error": "пользователь не является ментором"})
		return
	}
	workExperience, err := h.services.GetUserWorkExperience(mentorId)
	education, err := h.services.GetUserEducation(mentorId)
	classes, err := h.services.GetUserClasses(mentorId)
	communications, err := h.services.GetUserCommunications(mentorId)
	otherInfo, err := h.services.GetUserOtherInfo(mentorId)
	comments, err := h.services.GetComments(mentorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные пользователя"})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"first_name":      user.FirstName,
			"second_name":     user.SecondName,
			"patronymic":      user.Patronymic,
			"description":     user.Description,
			"specialization":  user.Specialization,
			"profile_picture": pathToProfilePicture + user.ProfilePicture,
			"time":            user.Time,
			"register_date":   user.CreatedAt,
			"work_experience": workExperience,
			"education":       education,
			"classes":         classes,
			"communications":  communications,
			"other_info":      otherInfo,
			"comments":        comments,
			"rating":          user.Rating,
		},
	)
}

func (h *Handler) GetMentiData(c *gin.Context) {
	mentiId := parseId(c.Param("id"))
	user, err := h.services.GetUserData(mentiId)
	comments, err := h.services.GetComments(mentiId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные пользователя"})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"first_name":      user.FirstName,
			"second_name":     user.SecondName,
			"patronymic":      user.Patronymic,
			"profile_picture": pathToProfilePicture + user.ProfilePicture,
			"register_date":   user.CreatedAt,
			"comments":        comments,
			"rating":          user.Rating,
		},
	)
}

func parseId(stringId string) uint {
	id, _ := strconv.ParseUint(stringId, 10, 64)
	userId := uint(id)
	return userId
}

func (h *Handler) GetUserStatistic(c *gin.Context) {
	userId := parseId(c.Param("id"))
	statistic, err := h.services.GetUserStatistic(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, statistic)
}
