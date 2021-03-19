package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

func TestAssetsMaintenanceRepo_InsertMaintenanceActivity_When_Success(t *testing.T) {
	ctx := context.Background()
	expectedMaintenanceActivity := domain.MaintenanceActivity{}
	id := uuid.New()

	status := "active"
	category := "laptop"
	purchaseAt := time.Now()
	purchaseCost := 45000.00
	name := "aspire-5"
	specifications := []byte(`{"ram":"4GB","brand":"acer"}`)

	startedDate := "28-02-1996"
	tym, _ := time.Parse("02-01-2006", startedDate)
	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM maintenance_activities")
	tx.MustExec("INSERT INTO assets (id, status, category, purchase_at, purchase_cost, name, specifications) VALUES ($1, $2, $3, $4, $5, $6, $7)", id, status, category, purchaseAt, purchaseCost, name, specifications)
	tx.Commit()
	req := domain.MaintenanceActivity{
		AssetId:     id,
		Cost:        100,
		StartedAt:   tym,
		Description: "hardware issue",
	}

	assetMaintainRepo := repository.NewAssetMaintainRepository()
	maintenanceActivity, err := assetMaintainRepo.InsertMaintenanceActivity(ctx, req)
	db.Get(&expectedMaintenanceActivity, "SELECT id,asset_id,cost,started_at,description FROM maintenance_activities WHERE id=$1", maintenanceActivity.ID)
	assert.Nil(t, err)
	assert.Equal(t, &expectedMaintenanceActivity, maintenanceActivity)

}

func TestAssetsMaintenanceRepo_InsertMaintenanceActivity_When_DbReturnsError(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()
	startedDate := "28-02-1996"
	tym, _ := time.Parse("02-01-2006", startedDate)
	req := domain.MaintenanceActivity{
		AssetId:     id,
		Cost:        100,
		StartedAt:   tym,
		Description: "hardware issue",
	}

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM maintenance_activities")
	tx.MustExec("DELETE FROM assets")
	tx.Commit()
	assetMaintainRepo := repository.NewAssetMaintainRepository()
	maintenanceActivity, err := assetMaintainRepo.InsertMaintenanceActivity(ctx, req)
	assert.Error(t, err)
	assert.Nil(t, maintenanceActivity)
}

func TestAssetsMaintenanceRepo_DetailedMaintenanceActivity_When_Success(t *testing.T) {
	ctx := context.Background()
	expectedMaintenanceActivity := domain.MaintenanceActivity{}
	id := 76
	assetId := uuid.New()

	status := "active"
	category := "laptop"
	purchaseAt := time.Now()
	purchaseCost := 5000.00
	name := "LG-5"
	specifications := []byte(`{"ram":"5GB","brand":"LG"}`)
	cost := 100
	description := "hardware issue"
	startedDate := "18-02-2020"
	tym, _ := time.Parse("02-01-2006", startedDate)
	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM maintenance_activities")
	tx.MustExec("INSERT INTO assets (id, status, category, purchase_at, purchase_cost, name, specifications) VALUES ($1, $2, $3, $4, $5, $6, $7)", assetId, status, category, purchaseAt, purchaseCost, name, specifications)
	tx.MustExec("INSERT INTO maintenance_activities (id,asset_id,cost,started_at,ended_at,description) VALUES ($1,$2,$3,$4,$5,$6)", id, assetId, cost, tym, tym, description)
	tx.Commit()
	db.Get(&expectedMaintenanceActivity, "SELECT * FROM maintenance_activities WHERE id=$1", id)
	assetMaintainRepo := repository.NewAssetMaintainRepository()
	maintenanceActivity, err := assetMaintainRepo.DetailedMaintenanceActivity(ctx, id)
	assert.Nil(t, err)
	assert.Equal(t, &expectedMaintenanceActivity, maintenanceActivity)
}

func TestAssetsMaintenanceRepo_DetailedMaintenanceActivity_When_IdDoesNotExists(t *testing.T) {
	ctx := context.Background()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM maintenance_activities")
	tx.Commit()
	assetMaintainRepo := repository.NewAssetMaintainRepository()
	maintenanceActivity, err := assetMaintainRepo.DetailedMaintenanceActivity(ctx, 100)
	assert.Nil(t, err)
	assert.Nil(t, maintenanceActivity)

}
