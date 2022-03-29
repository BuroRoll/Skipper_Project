package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"runtime"
)

const pathToProfilePicture = "/public-api/user/profile-picture/"

func (h *Handler) GetStatus(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": userId,
	})
	return
}

func (h *Handler) GetUserData(c *gin.Context) {
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
				"id":               user.ID,
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
				"specialization":   user.Specialization,
				"communications":   user.Communications,
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
	id, err := h.services.CreateUserCommunication(input, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания способа коммуникации"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "id": id})
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

func (h *Handler) UpdateProfilePicture(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	userId, _ := c.Get(userCtx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения файла"})
		return
	}
	filename := header.Filename
	_, err = h.services.Authorization.SaveProfilePicture(file, filename)
	err = h.services.UserData.UpdateProfilePicture(filename, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления данных пользователя"})
		return
	}
	profilePicture := pathToProfilePicture + filename
	c.JSON(http.StatusOK, gin.H{"ok": true, "profile_picture": profilePicture})
}

func (h *Handler) GetUserEducations(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	userEducation, err := h.services.GetUserEducation(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные об образовании пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"education": userEducation})
}

func (h *Handler) AddUserEducation(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	var input forms.UserEducationInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	id, err := h.services.CreateUserEducation(input, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить образование пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "id": id})
}

func (h *Handler) AddUserWorkExperience(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	var input forms.UserWorkExperience
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	id, err := h.services.CreateUserWorkExperience(input, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить опыт работы пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "id": id})
}

func (h *Handler) GetUserWorkExperience(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	userWorkExperience, err := h.services.GetUserWorkExperience(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные об опыте работы пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"work_experience": userWorkExperience})
}

func (h *Handler) UpdateSpecialization(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	var input forms.MentorSpecializationInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма данных"})
		return
	}
	err := h.services.UpdateMentorSpecialization(input.Specialization, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить данные"})
		return
	}
}

func (h *Handler) UserVerifyEmail(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	var input forms.UserEmailInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма данных"})
		return
	}
	err := h.services.SetUserEmail(input.Email, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить почту, попробуйте позже"})
		return
	}
	err = h.services.SendVerifyEmail(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось отправить письмо для подтверждения, попробуйте позже"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Подтверждение успешно отправлено"})
}

func (h *Handler) AddUserOtherInfo(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	var input forms.MentorOtherInfo
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	id, err := h.services.AddUserOtherInfo(input.Data, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось добавить данные"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "id": id})
}

func (h *Handler) GetUserOtherInfo(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	otherInfo, err := h.services.GetUserOtherInfo(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить данные пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"otherInfo": otherInfo,
	})
}

func (h *Handler) DeleteUserCommunication(c *gin.Context) {
	communicationId := c.Param("id")
	err := h.services.DeleteUserCommunication(communicationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить данные"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": "ok",
	})
}

func (h *Handler) DeleteUserEducation(c *gin.Context) {
	educationId := c.Param("id")
	err := h.services.DeleteUserEducation(educationId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить данные"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": "ok",
	})
}
func (h *Handler) DeleteUserWorkExperience(c *gin.Context) {
	workExperienceId := c.Param("id")
	err := h.services.DeleteUserWorkExperience(workExperienceId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить данные"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": "ok",
	})
}
func (h *Handler) DeleteUserOtherInfo(c *gin.Context) {
	otherInfoId := c.Param("id")
	err := h.services.DeleteUserOtherInfo(otherInfoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить данные"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"delete": "ok",
	})
}
