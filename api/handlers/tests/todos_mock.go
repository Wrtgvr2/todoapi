package handlers_test

import (
	"time"

	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

type MockTodoRepo struct{}

var (
	TestTodoID           uint64    = 1
	TestTodoUserID       uint64    = 1
	TestTodoTitle        string    = "testtitle"
	TestTodoTitle_BadReq string    = " "
	TestTodoDescription  string    = "test"
	TestTodoCompleted    bool      = false
	TestTodoCreatedAt    time.Time = time.Date(2008, time.September, 16, 15, 19, 26, 0, time.UTC)
)

var TestTodoUpdateData = models.UpdateTodoData{
	Title:       &TestTodoTitle,
	Description: &TestTodoDescription,
	Completed:   &TestTodoCompleted,
}

var TestTodoCreateData = models.CreateTodoData{
	User_ID:     TestUserID,
	Title:       TestTodoTitle,
	Description: &TestTodoDescription,
}

var TestTodoData = models.Todo{
	ID:          TestTodoID,
	User_ID:     TestTodoUserID,
	Title:       TestTodoTitle,
	Description: &TestTodoDescription,
	Completed:   &TestTodoCompleted,
	Created_At:  TestTodoCreatedAt,
}

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
