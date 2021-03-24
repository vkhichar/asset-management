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

	args := m.Called(ctx, user)

	var newUser *domain.User
	if args[0] != nil {
		newUser = args[0].(*domain.User)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return newUser, err
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

func (m *MockUserRepo) DeleteUser(ctx context.Context, id int) (string, error) {
	args := m.Called(ctx, id)

	var result string
	if args[0] != "" {
		result = args[0].(string)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	if args[0] == "" && args[1] == nil {
		return result, customerrors.UserDoesNotExist
	}
	return result, err
}

func (m *MockUserRepo) GetUserByID(ctx context.Context, ID int) (*domain.User, error) {

	var newUser *domain.User
	args := m.Called(ctx, ID)

	if args[0] == nil && args[1] == nil {
		return nil, customerrors.UserNotExist
	}
	if args[0] != nil {
		newUser = args[0].(*domain.User)
	}

	var err error

	if args[1] != nil {
		err = args[1].(error)
	}
	return newUser, err
}
