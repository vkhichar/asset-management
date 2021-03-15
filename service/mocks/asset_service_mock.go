package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

type MockAssetService struct {
	mock.Mock
}

func (m *MockAssetService) ListAssets(ctx context.Context) ([]domain.Asset, error) {
	args := m.Called(ctx)

	var assets []domain.Asset
	if args[0] != nil {
		assets = args[0].([]domain.Asset)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	if args[0] == nil && args[1] == nil {
		return assets, customerrors.NoAssetsExist
	}

	return assets, err

}

func (m *MockAssetService) UpdateAsset(ctx context.Context, Id uuid.UUID, req contract.UpdateRequest) (*domain.Asset, error) {
	args := m.Called(ctx, Id, req)

	var asset *domain.Asset
	if args[0] != nil {
		asset = args[0].(*domain.Asset)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	if args[0] == nil && args[1] == nil {
		return asset, customerrors.NoAssetsExist
	}
	return asset, err
}
func (m *MockAssetService) DeleteAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error) {
	args := m.Called(ctx, Id)

	var asset *domain.Asset
	if args[0] != nil {
		asset = args[0].(*domain.Asset)
	}
	var err error
	if args[1] != nil {
		err = args[1].(error)
	}
	if args[0] == nil && args[1] == nil {
		return asset, customerrors.NoAssetsExist
	}
	return asset, err
}
