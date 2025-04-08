package handlers_test

import (
	"github.com/wrtgvr/todoapi/internal/errdefs"
	"github.com/wrtgvr/todoapi/models"
)

type MockUserRepo struct{}

func (m MockUserRepo) GetUsers() ([]models.UserResponse, error) {
	return []models.UserResponse{
		TestUserRespData,
	}, nil
}

func (m MockUserRepo) DeleteUser(id uint64) error {
	if id == TestUserID {
		return nil
	}
	return errdefs.ErrUserNotFound
}

func (m MockUserRepo) GetFullUser(id uint64) (*models.User, error) {
	if id == TestUserID {
		return &TestUserData, nil
	}
	return nil, errdefs.ErrUserNotFound
}

func (m MockUserRepo) GetUserByUsername(username string) (*models.UserResponse, error) {
	if username == TestUsername {
		return &TestUserRespData, nil
	}
	return nil, errdefs.ErrUserNotFound
}

func (m MockUserRepo) GetUserById(id uint64) (*models.UserResponse, error) {
	if id == TestUserID {
		return &TestUserRespData, nil
	}
	return nil, errdefs.ErrUserNotFound
}

func (m MockUserRepo) CreateUser(userData *models.UserRequest) (*models.UserResponse, error) {
	return &TestUserRespData, nil
}

func (m MockUserRepo) UpdateUser(userData *models.User) (*models.UserResponse, error) {
	if userData.ID == TestUserID {
		return &TestUserRespData, nil
	}
	return nil, errdefs.ErrUserNotFound
}

func (m MockUserRepo) GetUserTodos(id uint64) ([]models.Todo, error) {
	if id == TestUserID {
		return []models.Todo{
			TestTodoData,
		}, nil
	}
	return nil, errdefs.ErrUserNotFound
}
