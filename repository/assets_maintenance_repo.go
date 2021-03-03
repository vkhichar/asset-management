package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

type AssetMaintenanceRepo interface {
	InsertMaintenanceActivity(ctx context.Context, asset_id domain.UUID, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error)
}

type assetMaintainRepo struct {
	db *sqlx.DB
}

func NewAssetMaintainRepository() AssetMaintenanceRepo {
	return &assetMaintainRepo{
		db: GetDB(),
	}
}
