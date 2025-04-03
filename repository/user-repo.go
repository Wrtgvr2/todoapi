package repository

import (
	"github.com/wrtgvr/todoapi/models"
)

type UserRepo interface {
	GetUsers() ([]models.UserResponse, error)
	DeleteUser(id uint64) error
	GetFullUser(id uint64) (*models.User, error)
	GetUserByUsername(username string) (*models.UserResponse, error)
	GetUserById(id uint64) (*models.UserResponse, error)
	CreateUser(userData *models.UserRequest) (*models.UserResponse, error)
	UpdateUser(userData *models.User) (*models.UserResponse, error)
	GetUserTodos(id uint64) ([]models.Todo, error)
}
