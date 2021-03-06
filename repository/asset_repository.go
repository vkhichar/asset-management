package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

const (
	getAssetDetails = "SELECT id,category,status,purchase_at,purchase_cost,name,specifications FROM assets"
)

type AssetRepository interface {
	ListAssets(ctx context.Context) ([]domain.Asset, error)
	CreateAsset(ctx context.Context, asset_param domain.Asset) (domain.Asset, error)
}

type assetRepo struct {
	db *sqlx.DB
}

func NewAssetRepository() AssetRepository {
	return &assetRepo{
		db: GetDB(),
	}
}

func (repo *assetRepo) ListAssets(ctx context.Context) ([]domain.Asset, error) {
	var as []domain.Asset

	err := repo.db.Select(&as, getAssetDetails)

	if err == sql.ErrNoRows {
		fmt.Printf("repository: No asset present")

		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (repo *assetRepo) CreateAsset(ctx context.Context, asset_param domain.Asset) (domain.Asset, error) {
	var asset domain.Asset
	err := repo.db.Get(&asset, createAssetQuery,
		asset_param.AssetID,
		asset_param.Status,
		asset_param.Category,
		asset_param.PurchaseAt,
		asset_param.PurchaseCost,
		asset_param.AssetName,
		asset_param.Specifications,
	)

	if err != nil {
		fmt.Printf("error in asset repository")
		return domain.Asset{}, err
	}

	return asset, nil
}
