package models

type UserResponse struct {
	Username string
	ID       uint64
}

type UserRequest struct {
	Username string
	Password string
}
