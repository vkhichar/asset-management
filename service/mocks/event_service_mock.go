package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockEventService struct {
	mock.Mock
}

func (m *MockEventService) PostAssetEvent(ctx context.Context, asset *domain.Asset) (string, error) {
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
