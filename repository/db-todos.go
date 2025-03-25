package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/wrtgvr/todoapi/models"
)

/*
ID          uint64    `json:"id"`
User_ID     uint64    `json:"userid"`
Title       string    `json:"title"`
Description string    `json:"description"`
Completed   bool      `json:"completed"`
Created_At  time.Time `json:"createdat"`
*/

func UpdateTodo(id uint64, updateData *models.UpdateTodoData) (*models.Todo, error) {
	query := `UPDATE todos SET title=$1, description=$2, completed=$3 WHERE id=$4 RETURNING *`
	todo := models.Todo{}

	err := DB.QueryRow(query, updateData.Title, updateData.Description, updateData.Completed, id).Scan(
		&todo.ID,
		&todo.User_ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.Created_At,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &ErrTodoNotFound{}
		}
		return nil, err
	}

	return &todo, nil
}

func DeleteTodo(id uint64) error {
	query := `DELETE FROM todos WHERE id=$1`

	err := DB.QueryRow(query, id).Err()
	if err != nil {
		if err == sql.ErrNoRows {
			return &ErrTodoNotFound{}
		}
		return err
	}

	return nil
}

func GetTodo(id uint64) (*models.Todo, error) {
	query := `SELECT * FROM todos WHERE id=$1`

	var todo models.Todo

	if err := DB.QueryRow(query, id).Scan(
		&todo.ID,
		&todo.User_ID,
		&todo.Title,
		&todo.Description,
		&todo.Completed,
		&todo.Created_At,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, &ErrTodoNotFound{}
		}
		return nil, err
	}

	return &todo, nil
}

func GetTodos() ([]models.Todo, error) {
	query := `SELECT * FROM todos;`

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

func CreateToDo(TodoData *models.CreateTodoData) (*models.Todo, error) {
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
