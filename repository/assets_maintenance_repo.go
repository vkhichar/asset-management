package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

const (
	createMaintainActivityByQuery = "INSERT INTO maintenance_activities (asset_id,cost,started_at,description) VALUES ($1,$2,$3,$4) RETURNING id,asset_id,cost,started_at,description"
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

	err := repo.db.Get(&maintenance, createMaintainActivityByQuery, req.AssetId, req.Cost, req.StartedAt, req.Description)
	if err != nil {
		fmt.Println("repolayer:%s", err.Error())
		return nil, err //have to do error check
	}
	return &maintenance, nil
}
