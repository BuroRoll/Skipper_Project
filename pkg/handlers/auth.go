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
	token, refreshToken, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	c.JSON(http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, refreshToken, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error generate token": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) refreshToken(c *gin.Context) {
	type tokenReqBody struct {
		RefreshToken string `json:"refreshToken"`
	}
	tokenReq := tokenReqBody{}
	err := c.Bind(&tokenReq)
	fmt.Println("get this token: ", tokenReq.RefreshToken)
	userId, err := h.services.ParseRefreshToken(tokenReq.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error generate token": err.Error()})
		return
	}
	token, refreshToken, err := h.services.Authorization.GenerateTokenByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error generate token": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) logout(c *gin.Context) {
	c.Header("Authorization", "")
	c.JSON(http.StatusOK, "Successfully logged out")
}
