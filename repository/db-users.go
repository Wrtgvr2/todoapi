package repository

import (
	"database/sql"
	"strings"

	_ "github.com/lib/pq"
	"github.com/wrtgvr/todoapi/models"
)

type PostgresUserRepo struct {
	DB *sql.DB
}

func (p *PostgresUserRepo) DeleteUser(id uint64) error {
	query := `DELETE FROM users WHERE id=$1`

	err := p.DB.QueryRow(query, id).Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrUserNotFound
		}
		return err
	}

	return nil
}

func (p *PostgresUserRepo) UpdateUser(newUserData *models.User) (*models.UserResponse, error) {
	query := `UPDATE users SET username=$1, password=$2 WHERE id=$3 RETURNING id, username`
	user := models.UserResponse{}

	err := p.DB.QueryRow(query, strings.ToLower(newUserData.Username), newUserData.Password, newUserData.ID).Scan(&user.ID, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (p *PostgresUserRepo) CreateUser(userData *models.UserRequest) (*models.UserResponse, error) {
	var user models.UserResponse
	query := `
	INSERT INTO users(username, password)
	VALUES($1, $2)
	RETURNING id, username;
	`

	err := p.DB.QueryRow(query, strings.ToLower(*userData.Username), userData.Password).Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PostgresUserRepo) GetFullUser(id uint64) (*models.User, error) {
	var user models.User

	query := `SELECT id, username, password FROM users WHERE id=$1;`

	err := p.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (p *PostgresUserRepo) GetUserByUsername(username string) (*models.UserResponse, error) {
	var user models.UserResponse

	lowerUsername := strings.ToLower(username)
	query := `SELECT id, username FROM users WHERE username=$1;`

	row := p.DB.QueryRow(query, lowerUsername)

	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (p *PostgresUserRepo) GetUserById(id uint64) (*models.UserResponse, error) {
	query := `SELECT id, username FROM users WHERE id=$1;`

	row := p.DB.QueryRow(query, id)

	var user models.UserResponse

	if err := row.Scan(&user.ID, &user.Username); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (p *PostgresUserRepo) GetUsers() ([]models.UserResponse, error) {
	query := `SELECT id, username FROM users;`

	rows, err := p.DB.Query(query)
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

func (p *PostgresUserRepo) GetUserTodos(id uint64) ([]models.Todo, error) {
	query := `SELECT * FROM todos WHERE user_id=$1`

	rows, err := DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(
			&todo.ID,
			&todo.User_ID,
			&todo.Title,
			&todo.Description,
			&todo.Completed,
			&todo.Created_At,
		); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}
