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

	var EventId string
	if args[0] != nil {
		EventId = args[0].(string)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	return EventId, err
}
