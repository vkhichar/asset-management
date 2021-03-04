package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

const (
	GetAssetDetails = "SELECT id,category,status,purchase_at,purchase_cost,name,specifications FROM assets"
)

type AssetRepository interface {
	// CreateAsset(ctx context.Context, assetCategory string, assetStatus string, purchaseCost float64, assetName string) (domain.Asset, error)
	ListAssets(ctx context.Context) ([]domain.Asset, error)
}

type assetRepo struct {
	db *sqlx.DB
}

func NewAssetRepository() AssetRepository {
	return &assetRepo{
		db: GetDB(),
	}
}

// func (repo *assetRepo) UpdateAsset(ctx context.Context,Id domain.UUID) (*domain.Asset,error){

// 	var asset domain.Asset
// 	err:=repo.db.

// }
// func (repo *assetRepo) CreateAsset(ctx context.Context, assetCategory string, assetStatus string, purchaseCost float64, assetName string) (domain.Asset, error) {

// 	var as domain.Asset

// 	return as, nil
// }

func (repo *assetRepo) ListAssets(ctx context.Context) ([]domain.Asset, error) {
	var as []domain.Asset

	err := repo.db.Select(&as, GetAssetDetails)
	if err == sql.ErrNoRows {
		fmt.Printf("repository: No asset present")

		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return as, nil

}
