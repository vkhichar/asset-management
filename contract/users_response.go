package contract

import (
	"time"

	"github.com/vkhichar/asset-management/domain"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUsersResponse(user domain.User) User {
	//Map domain to contract
	var u User
	u.ID = user.ID
	u.Name = user.Name
	u.Email = user.Email
	u.IsAdmin = user.IsAdmin
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
	return u
}
