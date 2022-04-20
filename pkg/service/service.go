package service

import (
	"github.com/muhammadkhon-abdulloev/todo-app"
	"github.com/muhammadkhon-abdulloev/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetByID(userID, id int) (todo.TodoList, error)
	Delete(userID, id int) error
	Update(userID int, id int, input todo.UpdateListInput) error
}

type TodoItem interface {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService (repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList: NewTodoListService(repos.TodoList),
	}
}