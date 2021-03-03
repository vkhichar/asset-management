package repository

import (
	"context"

	"github.com/vkhichar/asset-management/domain"
)

func (repo *userRepo) ShowUsers(ctx context.Context) (*domain.UserList, error) {
	var user []domain.UserList
	err := repo.db.Select(&user, "Select * from users")

	if err != nil {
		return nil, err
	}

	return &user, nil
}
