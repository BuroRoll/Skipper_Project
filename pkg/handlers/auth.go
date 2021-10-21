package handlers

import (
	"Skipper/pkg/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input body")
		return
	}
	_, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error generate token": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) logout(c *gin.Context) {
	//session := sessions.Default(c)
	//session.Clear()
	//if err := session.Save(); err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{"message": "Success logout"})
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, "Successfully logged out")
}
