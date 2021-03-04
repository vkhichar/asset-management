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

func (u User) DomainToContract(d *domain.User) User {
	var u User
ID:
}
