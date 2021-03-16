package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) PostUserEvent(ctx context.Context, user *domain.User) (string, error) {
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

func (service *MockTokenService) PostEvent(ctx context.Context, req domain.MaintenanceActivity) (int, error) {
	args := service.Called(ctx, req)

	if args[1] != nil {
		return 0, args[1].(error)
	}
	return args[0].(int), nil
}
