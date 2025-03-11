package repository

import (
	_ "github.com/lib/pq"
	"github.com/wrtgvr/todoapi/models"
)

func GetUsers() ([]models.User, error) {
	query := `SELECT id, username FROM users;`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

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
