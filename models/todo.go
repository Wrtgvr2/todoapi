package models

import "time"

type ToDo struct {
	ID          uint64
	User_ID     uint64
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
}
