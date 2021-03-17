package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
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
	args := m.Called(ctx)

	var users []domain.User
	if args[0] != nil {
		users = args[0].([]domain.User)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return users, err
}

func (m *MockUserRepo) UpdateUser(ctx context.Context, id int, req contract.UpdateUserRequest) (*domain.User, error) {
	args := m.Called(ctx, id, req)

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

func (m *MockUserRepo) DeleteUser(ctx context.Context, id int) (*domain.User, error) {
	args := m.Called(ctx, id)

	var user *domain.User
	if args[0] != nil {
		user = args[0].(*domain.User)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	if args[0] == nil && args[1] == nil {
		return user, customerrors.UserDoesNotExist
	}
	return user, err
}
