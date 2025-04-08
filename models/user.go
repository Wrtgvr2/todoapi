package models

type User struct {
	ID              uint64 `json:"id"`
	Username        string `json:"username"`
	DisplayUsername string `json:"displayusername"`
	Password        string `json:"password"`
}

type UserResponse struct {
	ID              uint64 `json:"id"`
	Username        string `json:"username"`
	DisplayUsername string `json:"displayusername"`
}

type UserRequest struct {
	Username        *string `json:"username,omitempty"`
	DisplayUsername *string `json:"displayusername,omitempty"`
	Password        *string `json:"password,omitempty"`
}
