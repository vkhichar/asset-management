// package repository

// import (
// 	"context"

// 	"github.com/vkhichar/asset-management/domain"
// )

// func (repo *userRepo) ListUsers(ctx context.Context) ([]domain.User, error) {
// 	var user []domain.User
// 	err := repo.db.Select(&user, "Select * from users")

// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
