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

	contractUser = CreateUserResponse{Name: user.Name, Email: user.Email, Password: user.Password, IsAdmin: user.IsAdmin, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	return contractUser

}
