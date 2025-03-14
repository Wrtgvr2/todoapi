package models

import "time"

type ToDo struct {
	ID          uint64    `json:"id"`
	User_ID     uint64    `json:"userid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdat"`
}
