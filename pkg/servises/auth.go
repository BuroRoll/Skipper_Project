package service

import (
	"Skipper/pkg/models"
	"Skipper/pkg/repository"
	"crypto/sha256"
	"fmt"
)

const (
	salt = "14hjqrhj1231qw124617ajfha1123ssfqa3ssjs190"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUser(email, password string) (uint, error) {
	return s.repo.GetUser(email, generatePasswordHash(password))
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
