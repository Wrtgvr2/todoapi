package handlers_test

import (
	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

type MockUserRepo struct{}

var (
	TestUsername        string = "testUser"
	TestPassword        string = "testPassword"
	TestUserID          uint64 = 1
	TestUsername_BadReq string = "err"
	TestPassword_BadReq string = "err"
)

func (m MockUserRepo) GetUsers() ([]models.UserResponse, error) {
	return []models.UserResponse{
		{ID: 1, Username: TestUsername},
		{ID: 2, Username: TestUsername},
	}, nil
}

func (m MockUserRepo) DeleteUser(id uint64) error {
	if id == TestUserID {
		return nil
	}
	return rep.ErrUserNotFound
}

func (m MockUserRepo) GetFullUser(id uint64) (*models.User, error) {
	if id == TestUserID {
		return &models.User{
			ID:       TestUserID,
			Username: TestUsername,
			Password: TestPassword,
		}, nil
	}
	return nil, rep.ErrUserNotFound
}

func (m MockUserRepo) GetUserByUsername(username string) (*models.UserResponse, error) {
	if username == TestUsername {
		return &models.UserResponse{
			ID:       TestUserID,
			Username: TestUsername,
		}, nil
	}
	return nil, rep.ErrUserNotFound
}

func (m MockUserRepo) GetUserById(id uint64) (*models.UserResponse, error) {
	if id == TestUserID {
		return &models.UserResponse{
			ID:       TestUserID,
			Username: TestUsername,
		}, nil
	}
	return nil, rep.ErrUserNotFound
}

func (m MockUserRepo) CreateUser(userData *models.UserRequest) (*models.UserResponse, error) {
	return &models.UserResponse{
		ID:       TestUserID,
		Username: *userData.Username,
	}, nil
}

func (m MockUserRepo) UpdateUser(userData *models.User) (*models.UserResponse, error) {
	if userData.ID == TestUserID {
		return &models.UserResponse{
			ID:       TestUserID,
			Username: TestUsername,
		}, nil
	}
	return nil, rep.ErrUserNotFound
}

func (m MockUserRepo) GetUserTodos(id uint64) ([]models.Todo, error) {
	return nil, nil
}
