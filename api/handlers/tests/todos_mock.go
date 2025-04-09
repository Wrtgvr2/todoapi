package handlers_test

import (
	"github.com/wrtgvr/todoapi/internal/errdefs"
	"github.com/wrtgvr/todoapi/models"
)

type MockTodoRepo struct{}

func (m MockTodoRepo) UpdateTodo(id uint64, updateData *models.UpdateTodoData) (*models.Todo, error) {
	if id == testTodoID {
		return &testTodoData, nil
	}
	return nil, errdefs.ErrTodoNotFound
}

func (m MockTodoRepo) DeleteTodo(id uint64) error {
	if id == testTodoID {
		return nil
	}
	return errdefs.ErrTodoNotFound
}

func (m MockTodoRepo) GetTodo(id uint64) (*models.Todo, error) {
	if id == testTodoID {
		return &testTodoData, nil
	}
	return nil, errdefs.ErrTodoNotFound
}

func (m MockTodoRepo) GetTodos() ([]models.Todo, error) {
	return []models.Todo{
		testTodoData,
	}, nil
}

func (m MockTodoRepo) CreateToDo(TodoData *models.CreateTodoData) (*models.Todo, error) {
	return &testTodoData, nil
}
