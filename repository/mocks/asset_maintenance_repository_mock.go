package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockMaintenanceActivityRepo struct {
	mock.Mock
}

func (m *MockMaintenanceActivityRepo) InsertMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
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

func (m *MockMaintenanceActivityRepo) DetailedMaintenanceActivity(ctx context.Context, activityId int) (*domain.MaintenanceActivity, error) {
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
