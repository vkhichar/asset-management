package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

const (
	createMaintainActivityByQuery   = "INSERT INTO maintenance_activities (asset_id,cost,started_at,description) VALUES ($1,$2,$3,$4) RETURNING id,asset_id,cost,started_at,description"
	detailedMaintainActivityByQuery = "SELECT id, asset_id, cost, started_at, ended_at, description FROM maintenance_activities WHERE id = $1"
	deleteById                      = "DELETE FROM maintenance_activities WHERE id = $1"
	getAllByAssetId                 = "SELECT id, asset_id, cost, started_at, ended_at, description FROM maintenance_activities WHERE asset_id = $1"
)

type AssetMaintenanceRepo interface {
	InsertMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error)
	DetailedMaintenanceActivity(ctx context.Context, activityId int) (*domain.MaintenanceActivity, error)
	DeleteMaintenanceActivity(ctx context.Context, activityId int) error
	GetAllByAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error)
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
	var maintenanceActivity domain.MaintenanceActivity

	err := repo.db.Get(&maintenanceActivity, createMaintainActivityByQuery, req.AssetId, req.Cost, req.StartedAt, req.Description)
	if err != nil {
		fmt.Printf("repolayer:Failed to insert asset maintenance activities due to %s", err.Error())
		return nil, err
	}

	return &maintenanceActivity, nil
}

func (repo *assetMaintainRepo) DetailedMaintenanceActivity(ctx context.Context, activityId int) (*domain.MaintenanceActivity, error) {
	var maintenanceActivity domain.MaintenanceActivity
	err := repo.db.Get(&maintenanceActivity, detailedMaintainActivityByQuery, activityId)
	if err == sql.ErrNoRows {
		fmt.Printf("repository: activity not present")

		return nil, nil
	}

	if err != nil {
		fmt.Printf("repolayer:Failed to fetch asset maintenance activities due to %s", err.Error())
		return nil, err
	}

	return &maintenanceActivity, nil
}

func (repo *assetMaintainRepo) DeleteMaintenanceActivity(ctx context.Context, activityId int) error {
	_, err := repo.db.Exec(deleteById, activityId)
	if err != nil {
		fmt.Printf("Failed to delete activity due to %s", err.Error())
		return errors.New("Failed to delete activity")
	}
	return nil
}

func (repo *assetMaintainRepo) GetAllByAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error) {
	var activities []domain.MaintenanceActivity
	err := repo.db.Select(&activities, getAllByAssetId, assetId)
	if err != nil {
		fmt.Printf("Failed to fetch asset maintenance activities due to %s", err.Error())
		return nil, errors.New("Failed to fetch asset maintenance activities")
	}
	return activities, nil
}
