package handlers_test

import (
	"github.com/wrtgvr/todoapi/internal/errdefs"
	"github.com/wrtgvr/todoapi/models"
)

type MockUserRepo struct{}

func (m MockUserRepo) GetUsers() ([]models.UserResponse, error) {
	return []models.UserResponse{
		testUserRespData,
	}, nil
}

func (m MockUserRepo) DeleteUser(id uint64) error {
	if id == testUserID {
		return nil
	}
	return errdefs.ErrUserNotFound
}

func (m MockUserRepo) GetFullUser(id uint64) (*models.User, error) {
	if id == testUserID {
		return &testUserData, nil
	}
	return nil, errdefs.ErrUserNotFound
}

func (m MockUserRepo) GetUserByUsername(username string) (*models.UserResponse, error) {
	if username == testUsername {
		return &testUserRespData, nil
	}
	return nil, errdefs.ErrUserNotFound
}

func (m MockUserRepo) GetUserById(id uint64) (*models.UserResponse, error) {
	if id == testUserID {
		return &testUserRespData, nil
	}
	return nil, errdefs.ErrUserNotFound
}

func (m MockUserRepo) CreateUser(userData *models.UserRequest) (*models.UserResponse, error) {
	return &testUserRespData, nil
}

func (m MockUserRepo) UpdateUser(userData *models.User) (*models.UserResponse, error) {
	if userData.ID == testUserID {
		return &testUserRespData, nil
	}
	return nil, errdefs.ErrUserNotFound
}

func (m MockUserRepo) GetUserTodos(id uint64) ([]models.Todo, error) {
	if id == testUserID {
		return []models.Todo{
			testTodoData,
		}, nil
	}
	return nil, errdefs.ErrUserNotFound
}
