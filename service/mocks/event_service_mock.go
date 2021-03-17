package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockEventService struct {
	mock.Mock
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
