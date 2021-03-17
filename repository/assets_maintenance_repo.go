package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

const (
	createMaintainActivityByQuery      = "INSERT INTO maintenance_activities (asset_id,cost,started_at,description) VALUES ($1,$2,$3,$4) RETURNING id,asset_id,cost,started_at,description"
	detailedMaintainActivityByQuery    = "SELECT id, asset_id, cost, started_at, ended_at, description FROM maintenance_activities WHERE id = $1"
	withoutDateMaintainActivityByQuery = "SELECT id, asset_id, cost, started_at,description FROM maintenance_activities WHERE id = $1"
	deleteById                         = "DELETE FROM maintenance_activities WHERE id = $1"
	getAllByAssetId                    = "SELECT id, asset_id, cost, started_at, ended_at, description FROM maintenance_activities WHERE asset_id = $1"
	updateQuery                        = "UPDATE maintenance_activities SET cost = $1, ended_at = $2 ,description = $3 WHERE id = $4"
	findByIdQuery                      = "SELECT id, asset_id, cost, started_at, ended_at, description FROM maintenance_activities WHERE id = $1"
)

var (
	NoActivityRecordFound = errors.New("Records not found")
)

type AssetMaintenanceRepo interface {
	InsertMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error)
	DetailedMaintenanceActivity(ctx context.Context, activityId int) (*domain.MaintenanceActivity, error)
	DeleteMaintenanceActivity(ctx context.Context, activityId int) error
	GetAllByAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error)
	UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error)
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

		errInDate := repo.db.Get(&maintenanceActivity, withoutDateMaintainActivityByQuery, activityId)
		if errInDate != nil {
			fmt.Printf("repolayer:Failed to fetch asset maintenance activities due to %s", err.Error())
			return nil, err
		}

		maintenanceActivity.EndedAt = &time.Time{}
		return &maintenanceActivity, nil

	}

	return &maintenanceActivity, nil
}

func (repo *assetMaintainRepo) DeleteMaintenanceActivity(ctx context.Context, activityId int) error {
	_, err := repo.db.Exec(deleteById, activityId)
	if err != nil {
		fmt.Printf("repository: Failed to delete activity due to %s", err.Error())
		return errors.New("Failed to delete activity")
	}
	return nil
}

func (repo *assetMaintainRepo) GetAllByAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error) {
	var activities []domain.MaintenanceActivity
	err := repo.db.Select(&activities, getAllByAssetId, assetId)
	if err != nil {
		fmt.Printf("repository: Failed to fetch asset maintenance activities due to %s", err.Error())
		return nil, errors.New("Failed to fetch asset maintenance activities")
	}
	return activities, nil
}

func (repo *assetMaintainRepo) UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Printf("Repository: Failed to begin trasanction: %s", err.Error())
		return nil, errors.New("Failed to update maintenance activity ")
	}
	tx.ExecContext(ctx, updateQuery, req.Cost, req.EndedAt, req.Description, req.ID)
	res, err := tx.ExecContext(ctx, updateQuery, req.Cost, req.EndedAt, req.Description, req.ID)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Failed to update asset maintenance activity")
	}
	if rc, _ := res.RowsAffected(); rc == 0 {
		fmt.Printf("repository: record not found with id %d \n", req.ID)
		return nil, customerrors.ErrNotFound
	}

	activity := domain.MaintenanceActivity{}
	row := tx.QueryRowContext(ctx, findByIdQuery, req.ID)

	err = row.Scan(&activity.ID, &activity.AssetId, &activity.Cost, &activity.StartedAt, &activity.EndedAt, &activity.Description)

	if err != nil {
		fmt.Printf("repository: record not found with id %s", err.Error())
		tx.Rollback()
		return nil, errors.New("Failed to update asset maintenance activity")
	}
	err = tx.Commit()

	if err != nil {
		fmt.Printf("repository: failed to commit trasaction %s \n", err.Error())
		tx.Rollback()
		return nil, errors.New("Failed to update asset maintenance activity")
	}
	return &activity, nil
}
