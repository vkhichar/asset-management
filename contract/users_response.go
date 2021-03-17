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

func DomainToContract(d *domain.User) User {
	user := User{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		IsAdmin:   d.IsAdmin,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
	return user
}
