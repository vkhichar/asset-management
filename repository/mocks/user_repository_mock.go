package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) FindUser(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)

	var user *domain.User
	if args[0] != nil {
		user = args[0].(*domain.User)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return user, err
}

func (m *MockUserRepo) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	// TODO: define mock method
	return nil, nil
}

func (m *MockUserRepo) ListUsers(ctx context.Context) ([]domain.User, error) {
	// TODO: define mock method
	return nil, nil
}
