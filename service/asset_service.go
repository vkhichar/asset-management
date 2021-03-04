package service

import (
	"context"
	"errors"

	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

var ErrNoUserExists = errors.New("No asset exists here")

type AssetService interface {
	ListAssets(ctx context.Context) ([]domain.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{
		assetRepo: repo,
	}
}

// func (service *assetService) CreateAsset(ctx context.Context, assetCategory string, assetStatus string, purchaseCost float64, assetName string, specifiations domain.Asset) (*domain.Asset, error) {
// 	asv, err := service.assetRepo.CreateAsset(ctx, assetCategory, assetStatus, purchaseCost, assetName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &asv, nil
// }

func (service *assetService) ListAssets(ctx context.Context) ([]domain.Asset, error) {

	asset, err := service.assetRepo.ListAssets(ctx)

	// if err == repository.NoUserExists {
	// 	return nil, ErrNoUserExists
	// }

	if err != nil {
		return nil, err
	}

	if asset == nil {
		return nil, customerrors.NoAssetsExist
	}

	return asset, nil

}
