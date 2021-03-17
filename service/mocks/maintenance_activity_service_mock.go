package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockMaintenanceActivityService struct {
	mock.Mock
}

func (m *MockMaintenanceActivityService) CreateAssetMaintenance(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
	args := m.Called(ctx, req)

	var maintenanceActivity *domain.MaintenanceActivity
	if args[0] != nil {
		maintenanceActivity = args[0].(*domain.MaintenanceActivity)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return maintenanceActivity, err
}
func (m *MockMaintenanceActivityService) DetailedMaintenanceActivity(ctx context.Context, activityId int) (*domain.MaintenanceActivity, error) {
	args := m.Called(ctx, activityId)

	var maintenanceActivity *domain.MaintenanceActivity
	if args[0] != nil {
		maintenanceActivity = args[0].(*domain.MaintenanceActivity)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return maintenanceActivity, err
}
