package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"runtime"
)

func (h *Handler) GetStatus(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
	})
	return
}

func (h *Handler) GetUserData(c *gin.Context) {
	const pathToProfilePicture = "/user/profile-picture/"
	userId, _ := c.Get(userCtx)
	if userId, ok := userId.(uint); ok {
		user, err := h.services.GetUserData(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные пользователя"})
			return
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"first_name":       user.FirstName,
				"second_name":      user.SecondName,
				"patronymic":       user.Patronymic,
				"date_of_birthday": user.DateOfBirthday,
				"description":      user.Description,
				"email":            user.Email,
				"phone":            user.Phone,
				"is_mentor":        user.IsMentor,
				"is_verify_email":  user.IsVerifyEmail,
				"is_verify_phone":  user.IsVerifyPhone,
				"profile_picture":  pathToProfilePicture + user.ProfilePicture,
				"time":             user.Time,
			})
	}
}

func (h *Handler) GetUserProfilePicture(c *gin.Context) {
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.Join(filepath.Dir(b), "../..")
	path := Root + "/media/user/profile_picture/" + c.Param("filename")
	c.FileAttachment(path, "profile_picture")
}

func (h *Handler) GetUserCommunications(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	communications, err := h.services.GetUserCommunications(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные пользователя"})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"communications": communications,
		},
	)
}

func (h *Handler) GetMessengers(c *gin.Context) {
	messengers, err := h.services.GetMessengers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список мессенджеров"})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"messengers": messengers,
		},
	)
}

func (h *Handler) CreateUserCommunication(c *gin.Context) {
	var input forms.UserCommunicationInput
	userId, _ := c.Get(userCtx)
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма добавления способа связи"})
		return
	}
	err := h.services.CreateUserCommunication(input, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания способа коммуникации"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *Handler) UpdateBaseProfileData(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	var input forms.UpdateBaseProfileData
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма обновления пользовательских данных"})
		return
	}
	err := h.services.UpdateBaseProfileData(input, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления данных пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
