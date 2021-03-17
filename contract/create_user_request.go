package contract

import (
	"errors"
	"regexp"
	"strings"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func (req CreateUserRequest) Validate() error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("Name is required")
	}

	email := req.Email

	valid := emailRegex.MatchString(email)

	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}

	if valid == false {

		return errors.New("invalid email")

	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}
