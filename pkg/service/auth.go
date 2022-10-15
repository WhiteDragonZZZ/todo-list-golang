package service

import (
	"crypto/sha1"
	"fmt"
	todo_list_golang "todo-list-golang"
	"todo-list-golang/pkg/repository"
)

const salt = "jhbfajhfdshfhladk"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo_list_golang.User) (int, error) {
	user.Password = s.generatedPasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatedPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
