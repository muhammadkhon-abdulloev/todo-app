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
	Create(userID int, listID int, item todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetByID(userID, itemID int) (todo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID int, itemID int, input todo.UpdateItemInput) error

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
		TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}