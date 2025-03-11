package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/wrtgvr/todoapi/models"
)

func GetUser(id uint64) (models.UserResponse, error) {
	query := `SELECT id, username FROM users WHERE id=$1;`

	row := DB.QueryRow(query, id)

	var user models.UserResponse

	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return models.UserResponse{}, &ErrUserNotFound{RequestedID: id}
		}
		return models.UserResponse{}, err
	}

	return user, nil
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
