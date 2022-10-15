package repository

import (
	"github.com/jmoiron/sqlx"
	todo_list_golang "todo-list-golang"
)

type Authorization interface {
	CreateUser(user todo_list_golang.User) (int, error)
	GetUser(username, password string) (todo_list_golang.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}

}
