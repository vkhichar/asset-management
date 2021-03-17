package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) PostCreateUserEvent(ctx context.Context, user *domain.User) (string, error) {

	args := m.Called(ctx, user)

	var newUser string
	if args[0] != nil {
		newUser = args[0].(string)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return newUser, err
}

func (m *MockEventService) PostUpdateUserEvent(ctx context.Context, user *domain.User) (string, error) {

	args := m.Called(ctx, user)

	var eventId string
	if args[0] != "" {
		eventId = args[0].(string)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return eventId, err
}
