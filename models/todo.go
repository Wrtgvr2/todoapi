package models

import "time"

type UpdateTodoData struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Completed   *bool   `json:"completed,omitempty"`
}

type CreateTodoData struct {
	User_ID     *uint64 `json:"user_id"`
	Title       *string `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
}

type Todo struct {
	ID          uint64    `json:"id"`
	User_ID     uint64    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Created_At  time.Time `json:"createdat"`
}
