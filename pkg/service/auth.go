package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/alextsa22/to-do-api/pkg/model"
	"github.com/alextsa22/to-do-api/pkg/repository"
)

const (
	salt = "ahdga24wnefmb2535iuhrg"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
