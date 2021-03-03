package service

import (
	"context"

	"github.com/vkhichar/asset-management/repository"
)

type AssetService interface {
	CreateAsset(ctx context.Context, assetCategory string, assetStatus string, purchaseCost float64, assetName string, spec *domain.Specs) (*domain.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{
		assetRepo: repo,
	}
}
