package service

import (
	"github.com/muhammadkhon-abdulloev/todo-app"
	"github.com/muhammadkhon-abdulloev/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, id int) (todo.TodoList, error) {
	return s.repo.GetByID(userID, id)
}

func (s *TodoListService) Delete(userID, id int) error {
	return s.repo.Delete(userID, id)
}

func (s *TodoListService) Update(userID int, id int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, id, input)
}
