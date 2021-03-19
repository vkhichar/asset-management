package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) PostAssetMaintenanceActivityEvent(ctx context.Context, resBody *domain.MaintenanceActivity) (string, error) {
	args := m.Called(ctx, resBody)

	var EventMaintenanceActivity string
	if args[0] != nil {
		EventMaintenanceActivity = args[0].(string)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return EventMaintenanceActivity, err
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

func (m *MockEventService) PostAssetEventCreateAsset(ctx context.Context, asset *domain.Asset) (string, error) {
	args := m.Called(ctx, asset)

	var eventID string
	if args[0] != nil {
		eventID = args[0].(string)
	}

	var errorString error
	if args[1] != nil {
		errorString = args[1].(error)
	}
	return eventID, errorString
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

func (service *MockEventService) PostMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (string, error) {
	args := service.Called(ctx, req)

	if args[1] != nil {
		return "", args[1].(error)
	}
	return args[0].(string), nil
}
