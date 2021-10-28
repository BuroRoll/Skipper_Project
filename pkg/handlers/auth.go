package handlers

import (
	"Skipper/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpUserForm
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, "invalid input body")
		return
	}
	_, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, refreshToken, err := h.services.Authorization.GenerateToken(input.Phone, input.Password)
	c.JSON(http.StatusOK, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, refreshToken, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error generate token": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) mentorSignUp(c *gin.Context) {
	var input models.SignUpMentorForm
	if err := c.MustBindWith(&input, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.services.Authorization.CreateMentorUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, refreshToken, err := h.services.Authorization.GenerateToken(input.Phone, input.Password)
	c.JSON(http.StatusOK, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) refreshToken(c *gin.Context) {
	var input models.TokenReqBody
	err := c.Bind(&input)
	userId, err := h.services.ParseRefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error generate token": err.Error()})
		return
	}
	token, _, err := h.services.Authorization.GenerateTokenByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error generate token": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) userToMentorSignUp(c *gin.Context) {
	var input models.SignUpUserToMentorForm
	if err := c.MustBindWith(&input, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := c.Get(userCtx)
	if id, ok := userId.(uint); ok {
		err := h.services.Authorization.UpgradeUserToMentor(id, input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error upgrade user": err.Error()})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": ok})
	}

}
