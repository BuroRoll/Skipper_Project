package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"runtime"
)

func (h *Handler) GetStatus(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	fmt.Println("UserId:", userId)
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось получить данные пользователя"})
		}
		c.JSON(http.StatusOK, gin.H{
			"first_name":      user.FirstName,
			"second_name":     user.SecondName,
			"is_mentor":       user.IsMentor,
			"profile_picture": pathToProfilePicture + user.ProfilePicture,
		})
	}
}

func (h *Handler) getUserProfilePicture(c *gin.Context) {
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.Join(filepath.Dir(b), "../..")
	path := Root + "/media/user/profile_picture/" + c.Param("filename")
	c.FileAttachment(path, "profile_picture")
}
