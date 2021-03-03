package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
)

type UUID [16]byte
type AssetMaintenanceRepo interface {
	InsertMaintenanceActivity(ctx context.Context, assets_id UUID, req contract.AssetMaintain) (*domain.Maintenance, error)
}

type assetMaintainRepo struct {
	db *sqlx.DB
}

func NewAssetMaintainRepository() AssetMaintenanceRepo {
	return &assetMaintainRepo{
		db: GetDB(),
	}
}
