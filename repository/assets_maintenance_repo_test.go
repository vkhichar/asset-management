package repository_test

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/vkhichar/asset-management/config"
// 	"github.com/vkhichar/asset-management/domain"
// 	"github.com/vkhichar/asset-management/repository"
// )

// func TestAssetsMaintenanceRepo_InsertMaintenanceActivity_When_Success(t *testing.T) {
// 	ctx := context.Background()
// 	expectedMaintenanceActivity := domain.MaintenanceActivity{}
// 	id, _ := uuid.Parse("ffb4b1a4-7bf5-11eb-9439-0242ac130002")
// 	startedDate := "28-02-1996"
// 	tym, err := time.Parse("02-01-2006", startedDate)
// 	req := domain.MaintenanceActivity{
// 		AssetId:     id,
// 		Cost:        100,
// 		StartedAt:   tym,
// 		Description: "hardware issue",
// 	}

// 	config.Init()
// 	repository.InitDB()
// 	db := repository.GetDB()
// 	tx := db.MustBegin()

// 	assetMaintainRepo := repository.NewAssetMaintainRepository()
// 	maintenanceActivity, err := assetMaintainRepo.InsertMaintenanceActivity(ctx, req)

// 	tx.MustExec("DELETE FROM maintenance_activities")
// 	tx.MustExec("INSERT INTO maintenance_activities (id,asset_id,cost,started_at,description) VALUES ($1,$2,$3,$4,$5)", maintenanceActivity.ID, maintenanceActivity.AssetId, maintenanceActivity.Cost, maintenanceActivity.StartedAt, maintenanceActivity.Description)
// 	tx.Commit()
// 	db.Get(&expectedMaintenanceActivity, "SELECT id,asset_id,cost,started_at,description FROM maintenance_activities WHERE id=$1", maintenanceActivity.ID)
// 	assert.Nil(t, err)
// 	assert.Equal(t, &expectedMaintenanceActivity, maintenanceActivity)
// }

// func TestAssetsMaintenanceRepo_InsertMaintenanceActivity_When_DbReturnsError(t *testing.T) {
// 	ctx := context.Background()
// 	id, _ := uuid.Parse("ffb4ba50-7bf5-11eb-9439-0242ac130002")
// 	startedDate := "28-02-1996"
// 	tym, _ := time.Parse("02-01-2006", startedDate)
// 	req := domain.MaintenanceActivity{
// 		AssetId:     id,
// 		Cost:        100,
// 		StartedAt:   tym,
// 		Description: "hardware issue",
// 	}

// 	config.Init()
// 	repository.InitDB()
// 	assetMaintainRepo := repository.NewAssetMaintainRepository()
// 	maintenanceActivity, err := assetMaintainRepo.InsertMaintenanceActivity(ctx, req)
// 	assert.Error(t, err)
// 	assert.Nil(t, maintenanceActivity)
// }

// func TestAssetsMaintenanceRepo_DetailedMaintenanceActivity_When_Success(t *testing.T) {
// 	ctx := context.Background()
// 	expectedMaintenanceActivity := domain.MaintenanceActivity{}
// 	id := 76
// 	assetId, _ := uuid.Parse("ffb4b1a4-7bf5-11eb-9439-0242ac130002")
// 	cost := 1000
// 	startedDate := "28-02-1996"
// 	tym, err := time.Parse("02-01-2006", startedDate)
// 	description := "hardware issue"
// 	config.Init()
// 	repository.InitDB()
// 	db := repository.GetDB()
// 	tx := db.MustBegin()
// 	tx.MustExec("DELETE FROM maintenance_activities")
// 	tx.MustExec("INSERT INTO maintenance_activities (id,asset_id,cost,started_at,ended_at,description) VALUES ($1,$2,$3,$4,$5,$6)", id, assetId, cost, tym, tym, description)
// 	tx.Commit()
// 	db.Get(&expectedMaintenanceActivity, "SELECT * FROM maintenance_activities WHERE id=$1", id)
// 	assetMaintainRepo := repository.NewAssetMaintainRepository()
// 	maintenanceActivity, err := assetMaintainRepo.DetailedMaintenanceActivity(ctx, id)
// 	assert.Nil(t, err)
// 	assert.Equal(t, &expectedMaintenanceActivity, maintenanceActivity)
// }

// func TestAssetsMaintenanceRepo_DetailedMaintenanceActivity_When_IdDoesNotExists(t *testing.T) {
// 	ctx := context.Background()
// 	config.Init()
// 	repository.InitDB()
// 	assetMaintainRepo := repository.NewAssetMaintainRepository()
// 	maintenanceActivity, err := assetMaintainRepo.DetailedMaintenanceActivity(ctx, 9888)
// 	assert.Nil(t, err)
// 	assert.Nil(t, maintenanceActivity)

// }
