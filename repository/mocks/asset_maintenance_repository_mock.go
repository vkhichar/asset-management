package mocks

import (
	"context"

	"github.com/google/uuid"
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

func (mock *MockMaintenanceActivityRepo) DeleteMaintenanceActivity(ctx context.Context, activityId int) error {
	args := mock.Called(ctx, activityId)

	if args[0] != nil {
		return args[0].(error)
	}
	return nil
}

func (mock *MockMaintenanceActivityRepo) GetAllByAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error) {
	args := mock.Called(ctx, assetId)
	if args[1] != nil {
		return nil, args[1].(error)
	}

	return args[0].([]domain.MaintenanceActivity), nil

}

func (mock *MockMaintenanceActivityRepo) UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {

	args := mock.Called(ctx, req)
	if args[1] != nil {
		return nil, args[1].(error)
	}
	return args[0].(*domain.MaintenanceActivity), nil
}
