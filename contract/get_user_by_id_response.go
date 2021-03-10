package contract

import (
	"time"

	"github.com/vkhichar/asset-management/domain"
)

type GetUserByID struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DomaintoContractUserID(user *domain.User) (contractUser GetUserByID) {

	contractUser = GetUserByID{Name: user.Name, Email: user.Email, IsAdmin: user.IsAdmin, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt}
	return contractUser

}
