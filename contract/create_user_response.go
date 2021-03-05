package contract

import (
	"time"

	"github.com/vkhichar/asset-management/domain"
)

type CreateUserResponse struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DomaintoContract(user *domain.User) (contractUser CreateUserResponse) {

	contractUser = CreateUserResponse{user.Name, user.Email, user.Password, user.IsAdmin, user.CreatedAt, user.UpdatedAt}
	return contractUser

}
