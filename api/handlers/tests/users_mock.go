package handlers_test

import (
	"github.com/wrtgvr/todoapi/models"
	rep "github.com/wrtgvr/todoapi/repository"
)

type MockUserRepo struct{}

var (
	TestUsername string = "testUser"
	TestPassword string = "TestPassword"
	TestUserID   uint64 = 1
)

func (m MockUserRepo) GetUsers() ([]models.UserResponse, error) {
	return []models.UserResponse{
		models.UserResponse{
			ID:       TestUserID,
			Username: TestUsername,
		},
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
	if *userData.Username == TestUsername && *userData.Password == TestPassword {
		return &models.UserResponse{
			ID:       TestUserID,
			Username: TestUsername,
		}, nil
	}
	return nil, TestInternalError
}

func (m MockUserRepo) UpdateUser(userData *models.User) (*models.UserResponse, error) {
	if userData.ID == TestUserID && userData.Username == TestUsername && userData.Password == TestPassword {
		return &models.UserResponse{
			ID:       TestUserID,
			Username: TestUsername,
		}, nil
	}
	return nil, rep.ErrUserNotFound
}
