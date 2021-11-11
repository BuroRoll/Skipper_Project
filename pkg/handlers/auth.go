package handlers

import (
	"Skipper/pkg/models/forms"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input forms.SignUpUserForm
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
	var input forms.SignInInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, refreshToken, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка генерации токена"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) mentorSignUp(c *gin.Context) {
	file, header, err := c.Request.FormFile("profile_picture")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	filename := header.Filename
	var input forms.SignUpMentorForm
	if err := c.MustBindWith(&input, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	_, err = h.services.Authorization.SaveProfilePicture(file, filename)
	_, err = h.services.Authorization.CreateMentorUser(input, filename)
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
	var input forms.TokenReqBody
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
	var input forms.SignUpUserToMentorForm
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
