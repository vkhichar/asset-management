package repository

import (
	"context"
	"github.com/vkhichar/asset-management/domain"
)

type UserRepository interface {
	FindUser(ctx context.Context, email string) (*domain.User, error)
}
