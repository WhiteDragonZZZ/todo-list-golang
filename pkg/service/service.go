package service

import (
	todo_list_golang "todo-list-golang"
	"todo-list-golang/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo_list_golang.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}

}
