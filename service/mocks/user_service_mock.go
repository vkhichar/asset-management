package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Login(ctx context.Context, email, password string) (*domain.User, string, error) {
	//login service
	return nil, "", nil
}

func (m *MockUserService) ListUsers(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	var users []domain.User
	if args[0] != nil {
		users = args[0].([]domain.User)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	if args[0] == nil && args[1] == nil {
		return users, customerrors.NoUsersExist
	}
	return users, err
}

func (m *MockUserService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	//create user service
	return nil, nil
}

func (m *MockUserService) UpdateUser(ctx context.Context, id int, req contract.UpdateUserRequest) (*domain.User, error) {
	args := m.Called(ctx, id, req)
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

func (m *MockUserService) DeleteUser(ctx context.Context, id int) (*domain.User, error) {
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
		return nil, customerrors.UserDoesNotExist
	}
	return user, err
}
