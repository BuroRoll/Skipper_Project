package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input forms.SignUpUserForm
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма регистрации"})
		return
	}
	_, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания профиля"})
		return
	}
	token, refreshToken, err := h.services.Authorization.GenerateToken(input.Phone, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания токенов"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input forms.SignInInput

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма авторизации"})
		return
	}

	token, refreshToken, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный логин или пароль"})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) mentorSignUp(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка получения файла"})
		return
	}
	filename := header.Filename
	var input forms.SignUpMentorForm
	if err := c.MustBindWith(&input, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма авторизации"})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка чтения токена"})
		return
	}
	token, _, err := h.services.Authorization.GenerateTokenByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка регенерации токена"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *Handler) userToMentorSignUp(c *gin.Context) {
	var input forms.SignUpUserToMentorForm
	if err := c.MustBindWith(&input, binding.FormMultipart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверная форма регистрации"})
		return
	}
	userId, _ := c.Get(userCtx)
	if id, ok := userId.(uint); ok {
		err := h.services.Authorization.UpgradeUserToMentor(id, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля"})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован как менти"})
	}
}

func (h *Handler) UserVerifyEmail(c *gin.Context) {
	userId, _ := c.Get(userCtx)
	err := h.services.SendVerifyEmail(userId.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось подтвердить почту, попробуйте позже"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Подтверждение успешно отправлено"})
}

func (h *Handler) verifyEmail(c *gin.Context) {
	token := c.Query("token")
	userId, err := h.services.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильная ссылка на подтверждение почты"})
		return
	}
	err = h.services.VerifyEmail(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось подтвердить почту"})
		return
	}
	//c.JSON(http.StatusOK, gin.H{
	//	"status":  "Почта успешно подтверждена",
	//	"user_id": userId,
	//})
	c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
}
