package service

import (
	"context"

	"fmt"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type AssetService interface {
	ListAssets(ctx context.Context) ([]domain.Asset, error)
	CreateAsset(ctx context.Context, asset domain.Asset) (*domain.Asset, error)
	GetAsset(ctx context.Context, ID uuid.UUID) (*domain.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{
		assetRepo: repo,
	}
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

func (service *assetService) CreateAsset(ctx context.Context, asset_param domain.Asset) (*domain.Asset, error) {
	asset, err := service.assetRepo.CreateAsset(ctx, asset_param)
	if err != nil {
		fmt.Printf("asset_service error while creating asset: %s", err.Error())
		return nil, err
	}

	return asset, err
}

func (service *assetService) GetAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error) {
	asset, err := service.assetRepo.GetAsset(ctx, Id)
	if err != nil {
		fmt.Printf("asset_service error while getting asset by it's ID: %s", err.Error())
		return nil, err
	}

	return asset, err
}

//test cover tool
