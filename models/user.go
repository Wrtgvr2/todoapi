package models

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
}

type UserRequest struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}
