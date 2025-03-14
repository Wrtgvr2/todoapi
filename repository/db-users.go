package repository

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	"github.com/wrtgvr/todoapi/models"
)

func CreateUser(userData models.UserRequest) (*models.UserResponse, error) {
	var user models.UserResponse
	query := `
	INSERT INTO users(username, password)
	VALUES($1, $2)
	RETURNING id, username;
	`

	err := DB.QueryRow(query, strings.ToLower(userData.Username), userData.Password).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*models.UserResponse, error) {
	var user models.UserResponse

	lowerUsername := strings.ToLower(username)
	query := `SELECT id, username FROM users WHERE username=$1;`

	row := DB.QueryRow(query, lowerUsername)

	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func GetUserById(id uint64) (*models.UserResponse, error) {
	query := `SELECT id, username FROM users WHERE id=$1;`

	row := DB.QueryRow(query, id)

	var user models.UserResponse

	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			fmt.Print(4)
			return nil, &ErrUserNotFound{RequestedID: id}
		}
		fmt.Print(5)
		return nil, err
	}

	return &user, nil
}

func GetUsers() ([]models.UserResponse, error) {
	query := `SELECT id, username FROM users;`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse

	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
