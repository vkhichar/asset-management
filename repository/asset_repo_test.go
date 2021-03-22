package repository_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

func TestAssetRepository_ListAssetsRepo_When_Success(t *testing.T) {
	ctx := context.Background()
	var expectedasset []domain.Asset

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()

	tx.MustExec("Delete from asset_allocations")
	tx.MustExec("Delete from maintenance_activities")
	tx.MustExec("Delete from assets")

	tx.MustExec("insert into assets(id,status,category,purchase_at,purchase_cost,name,specifications) values($1,$2,$3,$4,$5,$6,$7)", "ffb4b1a4-7bf5-11ee-9339-0242ac130002", "active", "Laptop", "01/07/2020", 500, "Dell Latitude E5550", `{"RAM":"4GB","HDD":"500GB","Generation":"i8"}`)
	tx.Commit()
	db.Select(&expectedasset, "Select id,status,category,purchase_at,purchase_cost,name,specifications from assets")

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.ListAssets(ctx)

	assert.Equal(t, expectedasset, asset)
	assert.Nil(t, err)
}

func TestAssetRepository_ListAssetsRepo_When_ReturnsError(t *testing.T) {
	ctx := context.Background()

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()

	tx.MustExec("Delete from assets")
	tx.Commit()

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.ListAssets(ctx)

	assert.Nil(t, asset)
	assert.Nil(t, err)

}

func TestAssetRepository_UpdateAssetsRepo_When_ReturnsError(t *testing.T) {
	ctx := context.Background()
	expectedasset := "No assets exist"
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130022")
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Genration"] = "i8"
	b, _ := json.Marshal(m)
	specifications := b

	req := contract.UpdateRequest{
		Status:         &status,
		Specifications: specifications,
	}

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("Delete from assets")

	tx.MustExec("insert into assets(id,status,category,purchase_at,purchase_cost,name,specifications) values($1,$2,$3,$4,$5,$6,$7)", "ffb4b1a4-7bf5-11ee-9339-0242ac130002", "active", "Laptop", "01/07/2020", 500, "Dell Latitude E5550", `{"RAM":"4GB","HDD":"500GB","Generation":"i8"}`)
	tx.Commit()
	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.UpdateAsset(ctx, Id, req)

	assert.Equal(t, expectedasset, err.Error())
	assert.Nil(t, asset)

}

func TestAssetRepository_DeleteRepo_When_ReturnsError(t *testing.T) {
	ctx := context.Background()
	expectedasset := "No assets exist"
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130022")

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("Delete from assets")
	tx.Commit()

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.DeleteAsset(ctx, Id)

	assert.Equal(t, expectedasset, err.Error())
	assert.Nil(t, asset)

}

func TestAssetRepository_UpdateAssetsRepo_When_Success(t *testing.T) {
	ctx := context.Background()
	var expectedasset domain.Asset
	fmt.Println("Above")
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	m["Genration"] = "i8"
	b, _ := json.Marshal(m)
	specifications := b

	req := contract.UpdateRequest{
		Status:         &status,
		Specifications: specifications,
	}

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()

	tx.MustExec("Delete from assets")

	tx.MustExec("insert into assets(id,status,category,purchase_at,purchase_cost,name,specifications) values($1,$2,$3,$4,$5,$6,$7)", "ffb4b1a4-7bf5-11ee-9339-0242ac130002", "active", "Laptop", "01/07/2020", 500, "Dell Latitude E5550", `{"RAM":"4GB","HDD":"500GB","Generation":"i8"}`)
	tx.Commit()
	fmt.Println("below")

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.UpdateAsset(ctx, Id, req)
	db.Get(&expectedasset, "Select id,status,category,purchase_at,purchase_cost,name,specifications from assets where id=$1", Id)

	assert.Equal(t, &expectedasset, asset)
	assert.Nil(t, err)

}

func TestAssetRepository_DeleteRepo_When_Success(t *testing.T) {
	ctx := context.Background()
	var expectedasset domain.Asset
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()

	tx.MustExec("Delete from assets")
	tx.MustExec("Delete from maintenance_activities")
	tx.MustExec("insert into assets(id,status,category,purchase_at,purchase_cost,name,specifications) values($1,$2,$3,$4,$5,$6,$7)", Id, "active", "Laptop", "01/07/2020", 500, "Dell Latitude E5550", `{"RAM":"4GB","HDD":"500GB","Generation":"i8"}`)
	tx.Commit()

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.DeleteAsset(ctx, Id)
	db.Get(&expectedasset, "select id,status,category,purchase_at,purchase_cost,name,specifications from assets where id=$1", Id)

	assert.Equal(t, &expectedasset, asset)
	assert.Nil(t, err)

}
