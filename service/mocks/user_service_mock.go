package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) FindUser(ctx context.Context, email string) (*domain.User, error) {
	// TODO:
	return nil, nil
}

func (m *MockUserService) Login(ctx context.Context, email, password string) (user *domain.User, token string, err error) {
	return nil, "", nil
}

func (m *MockUserService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {

	var newUser *domain.User
	args := m.Called(ctx, user)
	if args[0] != nil {
		newUser = args[0].(*domain.User)
	}

	var err error

	if args[1] != nil {
		err = args[1].(error)
	}
	return newUser, err
}

func (m *MockUserService) ListUsers(ctx context.Context) ([]domain.User, error) {
	// TODO: define mock method
	return nil, nil
}

func (m *MockUserService) GetUserByID(ctx context.Context, ID int) (*domain.User, error) {
	var newUser *domain.User
	args := m.Called(ctx, ID)
	if args[0] != nil {
		newUser = args[0].(*domain.User)
	}

	var err error

	if args[1] != nil {
		err = args[1].(error)
	}
	return newUser, err
}
