package handlers

import (
	"Skipper/pkg/models/forms"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"os"
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
	user, _ := h.services.GetUserData(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ошибка чтения токена"})
		return
	}
	token, _, err := h.services.Authorization.GenerateTokenByID(userId, user.IsMentor)
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
	file, header, err := c.Request.FormFile("file")
	filename := header.Filename
	_, err = h.services.Authorization.SaveProfilePicture(file, filename)
	err = h.services.UserData.UpdateProfilePicture(filename, userId.(uint))
	if id, ok := userId.(uint); ok {
		err := h.services.Authorization.UpgradeUserToMentor(id, input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления профиля"})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован как менти"})
		return
	}
	token, refreshToken, err := h.services.Authorization.GenerateTokenByID(userId.(uint), true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":        token,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) verifyEmail(c *gin.Context) {
	token := c.Query("token")
	userId, _, err := h.services.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильная ссылка на подтверждение почты"})
		return
	}
	err = h.services.VerifyEmail(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось подтвердить почту"})
		return
	}
	c.Redirect(http.StatusMovedPermanently, os.Getenv("FRONTEND"))
}

// ResetPassword
// @Description  Забыл пароль
// @Accept       json
// @Produce      json
// @Param 		 request 	body 		forms.ForgotPasswordInput	true 	"query params"
// @Success      200  		{object} 	forms.SuccessResponse
// @Failure 	 400 		{object}	forms.ErrorResponse
// @Router       /auth/reset-password [post]
func (h *Handler) ResetPassword(c *gin.Context) {
	var input forms.ForgotPasswordInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, forms.ErrorResponse{Error: "Неправильная форма сброса пароля"})
		return
	}
	err := h.services.ResetPassword(input.Login)
	if err != nil {
		c.JSON(http.StatusBadRequest, forms.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, forms.SuccessResponse{Status: "Форма сброса пароля успешно отправлена"})
}

// setNewPassword
// @Description  Установка нового пароля
// @Accept       json
// @Produce      json
// @Param 		 request 	body 		forms.SetNewPassword	true 	"query params"
// @Success      200  		{object} 	forms.SuccessResponse
// @Failure 	 400 		{object}	forms.ErrorResponse
// @Router       /auth/new-password [post]
func (h *Handler) setNewPassword(c *gin.Context) {
	var input forms.SetNewPassword
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, forms.ErrorResponse{Error: "Неправильная форма установк нового пароля"})
		return
	}
	userId, _, err := h.services.ParseToken(input.Token)
	err = h.services.SetNewPassword(userId, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, forms.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, forms.SuccessResponse{Status: "Пароль успешно обновлён"})
}
