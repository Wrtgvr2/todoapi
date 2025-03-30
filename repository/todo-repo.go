package repository

import (
	"github.com/wrtgvr/todoapi/models"
)

type TodoRepo interface {
	UpdateTodo(id uint64, updateData *models.UpdateTodoData) (*models.Todo, error)
	DeleteTodo(id uint64) error
	GetTodo(id uint64) (*models.Todo, error)
	GetTodos() ([]models.Todo, error)
	CreateToDo(TodoData *models.CreateTodoData) (*models.Todo, error)
}
