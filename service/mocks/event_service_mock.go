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
