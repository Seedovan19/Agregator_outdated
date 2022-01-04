package service

import (
	"crypto/sha1"
	"fmt"

	agregator "github.com/seedovan19/Agregator"
	"github.com/seedovan19/Agregator/pkg/repository"
)

const salt = "cqjksdsvnqui154981dmsioa"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user agregator.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
