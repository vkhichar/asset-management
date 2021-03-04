package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

type AssetMaintenanceRepo interface {
	InsertMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error)
}

type assetMaintainRepo struct {
	db *sqlx.DB
}

func NewAssetMaintainRepository() AssetMaintenanceRepo {
	return &assetMaintainRepo{
		db: GetDB(),
	}
}

func (repo *assetMaintainRepo) InsertMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
	var maintenance domain.MaintenanceActivity
	err := repo.db.Get(&maintenance, "Insert into maintenance_activities (assets_id,cost,description) values ($1,$2,$3) returning *", req.AssetId, req.Cost, req.Description)
	if err != nil {
		return nil, err //error check
	}
	return &maintenance, nil
}
