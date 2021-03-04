package contract

import (
	"errors"
	"strings"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (req CreateUserRequest) Validate() error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("Name is required")
	}

	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}
