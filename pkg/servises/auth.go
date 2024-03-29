package service

import (
	"Skipper/pkg/models/forms"
	"Skipper/pkg/repository"
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"gopkg.in/gomail.v2"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

const (
	salt                  = "14hjqrhj1231qw124617ajfha1123ssfqa3ssjs190"
	signingKey            = "qrkjk#4#%35FSFJlja#4353KSFjH"
	signingRefreshKey     = "qrkjk#sdfioh12bkj@nkk3k1axv["
	tokenTTL              = time.Hour * 12
	refreshTokenTTL       = time.Hour * 12 * 365
	resetPasswordTokenTTL = time.Hour
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   uint `json:"user_id"`
	IsMentor bool `json:"is_mentor"`
}

type refreshTokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func (s *AuthService) CreateUser(user forms.SignUpUserForm) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) CreateMentorUser(user forms.SignUpMentorForm, profilePicturePath string) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateMentor(user, profilePicturePath)
}

func (s *AuthService) GetUser(login, password string) (uint, bool, error) {
	return s.repo.GetUser(login, generatePasswordHash(password))
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(login, password string) (string, string, error) {
	userId, isMentor, err := s.GetUser(login, password)
	if err != nil {
		return "", "", err
	}
	return s.GenerateTokenByID(userId, isMentor)
}

func (s *AuthService) GenerateTokenByID(userId uint, isMentor bool) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
		isMentor,
	})
	t, err := token.SignedString([]byte(signingKey))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &refreshTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	rt, err := refreshToken.SignedString([]byte(signingRefreshKey))

	if err != nil {
		return "", "", err
	}

	return t, rt, err
}

func (s *AuthService) ParseRefreshToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &refreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingRefreshKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*refreshTokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (s *AuthService) ParseToken(accessToken string) (uint, bool, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, false, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, false, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.IsMentor, nil
}

func (s *AuthService) UpgradeUserToMentor(userId uint, formData forms.SignUpUserToMentorForm) error {
	return s.repo.UpgradeUserToMentor(userId, formData)
}

func (s *AuthService) SaveProfilePicture(file multipart.File, filename string) (string, error) {
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.Join(filepath.Dir(b), "../..")
	filePath := Root + "/media/user/profile_picture/" + filename
	out, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			return
		}
	}(out)
	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func (s *AuthService) SendVerifyEmail(userId uint) error {
	user, err := s.repo.GetUserById(userId)
	token, _, err := s.GenerateTokenByID(userId, user.IsMentor)
	if err != nil {
		fmt.Println(err)
	}
	userName := user.SecondName + " " + user.FirstName
	err = SendEmail(user.Email,
		token,
		"https://skipper.ga/verify-email?",
		"Подтверждение почты Skipper",
		"Подтвердить",
		"Если вы не регистрировали эту учетную запись, проигнорируйте это сообщение",
		userName)
	return err
}

func (s *AuthService) SendResetPasswordEmail(userId uint, email string, userFirstName string, userSecondName string) error {
	token, err := GenerateTokenForResetPassword(userId)
	userName := userSecondName + " " + userFirstName
	err = SendEmail(email,
		token,
		os.Getenv("FRONTEND")+"/reset-password",
		"Смена пароля Skipper",
		"Сбросить пароль",
		"Если вы не запрашивали смену пароля, проигнорируйте это сообщение",
		userName)
	if err != nil {
		fmt.Println(err)
		return errors.New("Не удалось отправить сообщение ")
	}
	return nil
}

func SendEmail(email string, token string, link string, theme string, buttonText string, description string, userName string) error {
	type data struct {
		Title       string
		Token       string
		Link        string
		ButtonText  string
		Theme       string
		Description string
		UserName    string
	}
	userData := data{
		Title:       theme,
		Token:       token,
		Link:        link,
		ButtonText:  buttonText,
		Theme:       theme,
		Description: description,
		UserName:    userName,
	}
	_, b, _, _ := runtime.Caller(0)
	Root := filepath.Join(filepath.Dir(b), "../..")
	t := template.New("index.html")
	var err error
	t, err = t.ParseFiles(Root + "/" + "index.html")
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, userData); err != nil {
		log.Println(err)
	}
	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", email)
	m.SetHeader("Subject", theme)
	m.SetBody("text/html", result)
	emailPort, _ := strconv.Atoi(os.Getenv("EMAIL_PORT"))
	d := gomail.NewDialer(os.Getenv("EMAIL_HOST"), emailPort, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) VerifyEmail(userId uint) error {
	err := s.repo.VerifyEmail(userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) ResetPassword(login string) error {
	user, err := s.repo.GetUserByEmailOrPhone(login)
	if err != nil {
		return errors.New("Пользователь не найден ")
	}
	err = s.SendResetPasswordEmail(user.ID, user.Email, user.FirstName, user.SecondName)
	return err
}

type resetPasswordTokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

func GenerateTokenForResetPassword(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &resetPasswordTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(resetPasswordTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})
	resetPasswordToken, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", errors.New("Ошибка генерации токена сброса пароля ")
	}
	return resetPasswordToken, nil
}

func (s *AuthService) SetNewPassword(userId uint, newPassword string) error {
	user, err := s.repo.GetUserById(userId)
	password := generatePasswordHash(newPassword)
	if err != nil {
		return errors.New("Ошибка получения данных о пользователе ")
	}
	err = s.repo.ChangeUserPassword(user, password)
	if err != nil {
		return errors.New("Ошибка обновления пароля ")
	}
	return nil
}
