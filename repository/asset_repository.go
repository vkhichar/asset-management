package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type AssetRepository interface {
	CreateAsset(ctx context.Context, assetCategory string, assetStatus string, purchaseCost float64, assetName string, spec *domain.Specs) (*domain.Asset, error)
}

type assetRepo struct {
	db *sqlx.DB
}

func NewAssetRepository() AssetRepository {
	return &assetRepo{
		db: GetDB(),
	}
}
