package service

import (
	"context"
	"github.com/vkhichar/asset-management/domain"
)

type UserService interface {
	Login(ctx context.Context, email, password string) (user *domain.User, token string, err error)
}

func NewUserService() UserService {
	return nil
}
