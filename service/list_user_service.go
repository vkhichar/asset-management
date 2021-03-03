package service

import (
	"context"

	"github.com/vkhichar/asset-management/domain"
)

func (service *userService) ListUser(ctx context.Context) ([]domain.User, error) {
	user, err := service.userRepo.ShowUsers(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
