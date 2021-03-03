package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

type AssetMaintenanceRepo interface {
	InsertMaintenance(ctx context.Context, assets_id string) (*domain.Maintenance, error)
	GetMaintenanceDetail(ctx context.Context, assets_id string) (*domain.Maintenance, error)
}

type assetMaintainRepo struct {
	db *sqlx.DB
}

func NewAssetMaintainRepository() AssetMaintenanceRepo {
	return &assetMaintainRepo{
		db: GetDB(),
	}
}
