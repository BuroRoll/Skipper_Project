package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateUserClass(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	var class forms.ClassesInput
	if err := c.BindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	classId, err := h.services.CreateUserClass(class, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания занятия"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "class_is": classId})
}

func (h *Handler) CreateTheoreticClass(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	var theoreticClass forms.TheoreticClassInput
	if err := c.BindJSON(&theoreticClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	classId, err := h.services.CreateTheoreticClass(theoreticClass, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания занятия"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "class_id": classId})
}

func (h *Handler) CreatePracticClass(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	var practicClass forms.PracticClassInput
	if err := c.BindJSON(&practicClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	classId, err := h.services.CreatePracticClass(practicClass, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания занятия"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "class_id": classId})
}

func (h *Handler) CreateKeyClass(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	var keyClass forms.KeyClass
	if err := c.BindJSON(&keyClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	classId, err := h.services.CreateKeyClass(keyClass, userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания занятия"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "class_id": classId})
}

func (h *Handler) GetUserClasses(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	classes, err := h.services.GetUserClasses(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения занятий пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"classes": classes})
}

func (h *Handler) DeleteClass(c *gin.Context) {
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	classId := c.Param("id")
	err := h.services.DeleteClass(classId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) DeleteTheoreticClass(c *gin.Context) {
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	classId := c.Param("id")
	err := h.services.DeleteTheoreticClass(classId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) DeletePracticClass(c *gin.Context) {
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	classId := c.Param("id")
	err := h.services.DeletePracticClass(classId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) DeleteKeyClass(c *gin.Context) {
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	classId := c.Param("id")
	err := h.services.DeleteKeyClass(classId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) UpdateClass(c *gin.Context) {
	var class forms.UpdateClassesInput
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	if err := c.BindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	err := h.services.UpdateClass(class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) UpdateTheoreticClass(c *gin.Context) {
	var theoreticClass forms.UpdateSubclassInput
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	if err := c.BindJSON(&theoreticClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	err := h.services.UpdateTheoreticClass(theoreticClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) UpdatePracticClass(c *gin.Context) {
	var theoreticClass forms.UpdateSubclassInput
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	if err := c.BindJSON(&theoreticClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	err := h.services.UpdatePracticClass(theoreticClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) UpdateKeyClass(c *gin.Context) {
	var keyClass forms.UpdateKeyClassInput
	isMentor, _ := c.Get(isMentorCtx)
	if !isMentor.(bool) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь не является ментором"})
		return
	}
	if err := c.BindJSON(&keyClass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма пользовательских данных"})
		return
	}
	err := h.services.UpdateKeyClass(keyClass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить занятие"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
