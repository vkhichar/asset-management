package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type AssetService interface {
	ListAssets(ctx context.Context) ([]domain.Asset, error)
	UpdateAsset(ctx context.Context, Id uuid.UUID, req contract.UpdateRequest) (*domain.Asset, error)
	DeleteAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{
		assetRepo: repo,
	}
}
func (service *assetService) DeleteAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error) {
	asset, err := service.assetRepo.DeleteAsset(ctx, Id)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (service *assetService) UpdateAsset(ctx context.Context, Id uuid.UUID, req contract.UpdateRequest) (*domain.Asset, error) {
	asset, err := service.assetRepo.UpdateAsset(ctx, Id, req)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (service *assetService) ListAssets(ctx context.Context) ([]domain.Asset, error) {

	asset, err := service.assetRepo.ListAssets(ctx)
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, customerrors.NoAssetsExist
	}
	return asset, nil
}
