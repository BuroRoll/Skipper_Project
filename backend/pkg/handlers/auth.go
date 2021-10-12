package handlers

import (
	"Skipper/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(input)
		c.JSON(http.StatusBadRequest, "invalid input body")
		return
	}
	userId, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", userId)
	c.JSON(http.StatusOK, gin.H{"user_id": userId})
}

func (h *Handler) signIn(c *gin.Context) {
	var user models.User
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid input body")
	}

	userId, err := h.services.Authorization.GetUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", userId)

	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user", "userId": userId})
}

func (h *Handler) logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success logout"})
}
