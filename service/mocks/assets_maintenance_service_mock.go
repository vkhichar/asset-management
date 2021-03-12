package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockAssetMaintenanceService struct {
	mock.Mock
}

func (mock *MockAssetMaintenanceService) CreateAssetMaintenance(ctx context.Context, req domain.MaintenanceActivity) (user *domain.MaintenanceActivity, err error) {
	return nil, nil
}

func (mock *MockAssetMaintenanceService) DeleteMaintenanceActivity(ctx context.Context, id int) (err error) {
	args := mock.Called(ctx, id)

	if args[0] != nil {
		return args[0].(error)
	}
	return nil
}

func (mock *MockAssetMaintenanceService) GetAllForAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error) {
	args := mock.Called(ctx, assetId)
	if args[1] != nil {
		return nil, args[1].(error)
	}

	return args[0].([]domain.MaintenanceActivity), nil
}

func (mock *MockAssetMaintenanceService) UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
	args := mock.Called(ctx, req)
	if args[1] != nil {
		return nil, args[1].(error)
	}
	return args[0].(*domain.MaintenanceActivity), nil
}
