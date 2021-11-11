package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	user_id, _ := c.Get(userCtx)
	if userId, ok := user_id.(uint); ok {
		user, err := h.services.GetUserData(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"first_name":      user.FirstName,
			"second_name":     user.SecondName,
			"is_mentor":       user.IsMentor,
			"profile_picture": user.ProfilePicture,
		})
	}
}

func (h *Handler) getUserProfilePicture(c *gin.Context) {
	path := "media/user/profile_picture/" + c.Param("filename")
	fmt.Println(path)
	c.FileAttachment(path, "profile_picture")
}
