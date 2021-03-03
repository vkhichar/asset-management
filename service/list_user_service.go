package service

import (
	"context"

	"github.com/vkhichar/asset-management/domain"
)

func (user *userService) ListUser(ctx context.Context) (users *domain.UserList, err error) {
	user[], err := service.userRepo.ShowUsers(ctx)
	if err != nil {
		return nil, err
	}
	return user[], nil
}
