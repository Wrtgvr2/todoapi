package repository

import (
	_ "github.com/lib/pq"
	"github.com/wrtgvr/todoapi/models"
)

/*
ID          uint64    `json:"id"`
User_ID     uint64    `json:"userid"`
Title       string    `json:"title"`
Description string    `json:"description"`
Completed   bool      `json:"completed"`
Created_At   time.Time `json:"createdat"`
*/
func GetTodos() ([]models.Todo, error) {
	query := `SELECT * from todos;`

	rows, err := DB.Query(query)
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

func CreateToDo(TodoData models.CreateTodoData) (*models.Todo, error) {
	var todo models.Todo
	query := `
	INSERT INTO todos(user_id, title, description)
	VALUES($1, $2, $3)
	RETURNING id, user_id, title, description, completed, created_at;
	`

	err := DB.QueryRow(query, TodoData.User_ID, TodoData.Title, TodoData.Description).Scan(
		&todo.ID,
		&todo.User_ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.Created_At,
	)
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
