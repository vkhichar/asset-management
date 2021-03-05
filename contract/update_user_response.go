package contract

import (
	"github.com/vkhichar/asset-management/domain"
)

type UpdateUserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func DomainToContractUpdate(d *domain.User) UpdateUserResponse {
	u := UpdateUserResponse{
		ID:        d.ID,
		Name:      d.Name,
		Email:     d.Email,
		IsAdmin:   d.IsAdmin,
		CreatedAt: d.CreatedAt.String(),
		UpdatedAt: d.UpdatedAt.String(),
	}
	return u
}
