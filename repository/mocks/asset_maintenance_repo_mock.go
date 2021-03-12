package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockAssetMaintenanceRepo struct {
	mock.Mock
}

func (mock *MockAssetMaintenanceRepo) InsertMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
	return nil, nil
}

func (mock *MockAssetMaintenanceRepo) DeleteMaintenanceActivity(ctx context.Context, activityId int) error {
	args := mock.Called(ctx, activityId)

	if args[0] != nil {
		return args[0].(error)
	}
	return nil
}

func (mock *MockAssetMaintenanceRepo) GetAllByAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error) {
	args := mock.Called(ctx, assetId)
	if args[1] != nil {
		return nil, args[1].(error)
	}

	return args[0].([]domain.MaintenanceActivity), nil

}

func (mock *MockAssetMaintenanceRepo) UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {

	args := mock.Called(ctx, req)
	if args[1] != nil {
		return nil, args[1].(error)
	}
	return args[0].(*domain.MaintenanceActivity), nil
}
