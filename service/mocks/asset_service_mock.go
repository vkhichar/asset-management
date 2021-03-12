package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockAssetService struct {
	mock.Mock
}

func (m *MockAssetService) ListAssets(ctx context.Context) ([]domain.Asset, error) {
	return nil, nil
}

func (m *MockAssetService) CreateAsset(ctx context.Context, asset_param domain.Asset) (*domain.Asset, error) {
	args := m.Called(ctx, asset_param)
	var asset *domain.Asset
	if args[0] != nil {
		asset = args[0].(*domain.Asset)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return asset, err
}

func (m *MockAssetService) GetAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error) {
	args := m.Called(ctx, Id)

	var asset *domain.Asset
	if args[0] != nil {
		asset = args[0].(*domain.Asset)
	}

	var err error
	if args[1] != nil {
		err = args[1].(error)
	}

	return asset, err
}
