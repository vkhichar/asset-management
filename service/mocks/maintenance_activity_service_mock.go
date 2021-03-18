package mocks

import (
	"context"

	"github.com/google/uuid"
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

func (mock *MockMaintenanceActivityService) DeleteMaintenanceActivity(ctx context.Context, id int) (err error) {
	args := mock.Called(ctx, id)

	if args[0] != nil {
		return args[0].(error)
	}
	return nil
}

func (mock *MockMaintenanceActivityService) GetAllForAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error) {
	args := mock.Called(ctx, assetId)
	if args[1] != nil {
		return nil, args[1].(error)
	}

	return args[0].([]domain.MaintenanceActivity), nil
}

func (mock *MockMaintenanceActivityService) UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
	args := mock.Called(ctx, req)
	if args[1] != nil {
		return nil, args[1].(error)
	}
	return args[0].(*domain.MaintenanceActivity), nil
}
