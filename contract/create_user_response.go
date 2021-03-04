package contract

import "time"

type CreateUserResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}
