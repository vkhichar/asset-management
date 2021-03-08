package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
)

type MockAssetRepo struct {
	mock.Mock
}

func (m *MockAssetRepo) CreateAsset(ctx context.Context, asset_param domain.Asset) (*domain.Asset, error) {
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

func (m *MockAssetRepo) ListAssets(ctx context.Context) ([]domain.Asset, error) {
	return nil, nil
}
