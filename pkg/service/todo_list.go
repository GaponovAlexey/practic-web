package service

import (
	todo "github.com/GaponovAlexey/practic-web"
	"github.com/GaponovAlexey/practic-web/pkg/repository"
)

type TodoListsService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListsService {
	return &TodoListsService{repo: repo}
}

func (s *TodoListsService) Create(userId int, list todo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListsService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListsService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}