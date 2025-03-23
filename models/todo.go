package models

import "time"

type CreateTodoData struct {
	User_ID     uint64 `json:"user_id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type Todo struct {
	ID          uint64    `json:"id"`
	User_ID     uint64    `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	Created_At  time.Time `json:"createdat"`
}
