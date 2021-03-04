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
