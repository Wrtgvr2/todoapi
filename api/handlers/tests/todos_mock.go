package handlers_test

import (
	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

type MockTodoRepo struct{}

func (m MockTodoRepo) UpdateTodo(id uint64, updateData *models.UpdateTodoData) (*models.Todo, error) {
	if id == TestTodoID {
		return &TestTodoData, nil
	}
	return nil, rep.ErrTodoNotFound
}

func (m MockTodoRepo) DeleteTodo(id uint64) error {
	if id == TestTodoID {
		return nil
	}
	return rep.ErrTodoNotFound
}

func (m MockTodoRepo) GetTodo(id uint64) (*models.Todo, error) {
	if id == TestTodoID {
		return &TestTodoData, nil
	}
	return nil, rep.ErrTodoNotFound
}

func (m MockTodoRepo) GetTodos() ([]models.Todo, error) {
	return []models.Todo{
		TestTodoData,
	}, nil
}

func (m MockTodoRepo) CreateToDo(TodoData *models.CreateTodoData) (*models.Todo, error) {
	return &TestTodoData, nil
}
