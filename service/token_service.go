package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Claims struct {
	UserID  int
	IsAdmin bool
}

type TokenService interface {
	GenerateToken(c *Claims) (string, error)
	ValidateToken(token string) (*Claims, error)
}

type plainToken struct{}

func NewPlainTokenService() TokenService {
	return &plainToken{}
}

func (p *plainToken) GenerateToken(c *Claims) (string, error) {
	token := fmt.Sprintf("userid:%d;is_admin:%t", c.UserID, c.IsAdmin)
	return token, nil
}

func (p *plainToken) ValidateToken(token string) (*Claims, error) {
	tokenParts := strings.Split(token, ";")
	if len(tokenParts) != 2 {
		return nil, errors.New("invalid token")
	}

	userIDParts := strings.Split(tokenParts[0], ":")
	if len(userIDParts) != 2 {
		return nil, errors.New("invalid token")
	}

	roleParts := strings.Split(tokenParts[1], ":")
	if len(roleParts) != 2 {
		return nil, errors.New("invalid token")
	}

	userID, err := strconv.Atoi(userIDParts[1])
	if err != nil {
		return nil, err
	}

	isAdmin, err := strconv.ParseBool(roleParts[1])
	if err != nil {
		return nil, err
	}

	return &Claims{
		UserID:  userID,
		IsAdmin: isAdmin,
	}, nil
}
