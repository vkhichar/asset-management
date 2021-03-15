package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
	mockRepo "github.com/vkhichar/asset-management/repository/mocks"
	"github.com/vkhichar/asset-management/service"
)

func TestAssetService_UpdateAsset_When_Success(t *testing.T) {
	ctx := context.Background()
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	dat, _ := time.Parse(layout, str)
	cost, _ := strconv.ParseFloat("5000", 32)
	m := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	b, _ := json.Marshal(m)
	asset := domain.Asset{

		Id:             Id,
		Status:         "active",
		Category:       "Laptop",
		PurchaseAt:     dat,
		PurchaseCost:   cost,
		Name:           "Dell Latitude E5550",
		Specifications: b,
	}
	js := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	jsr, _ := json.Marshal(js)
	Specifications := jsr
	Status := "active"
	req := contract.UpdateRequest{
		Status:         &Status,
		Specifications: Specifications,
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(&asset, nil)
	assetService := service.NewAssetService(mockAssetRepo)
	DBasset, issue := assetService.UpdateAsset(ctx, Id, req)

	assert.Nil(t, issue)
	assert.Equal(t, &asset, DBasset)

}
func TestAssetService_DeleteAsset_Success(t *testing.T) {
	ctx := context.Background()
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	dat, _ := time.Parse(layout, str)
	cost, _ := strconv.ParseFloat("5000", 32)
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, _ := json.Marshal(m)

	asset := domain.Asset{

		Id:             Id,
		Status:         "active",
		Category:       "Laptop",
		PurchaseAt:     dat,
		PurchaseCost:   cost,
		Name:           "Dell Latitude E5550",
		Specifications: b,
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(&asset, nil)
	assetService := service.NewAssetService(mockAssetRepo)
	DBasset, issue := assetService.DeleteAsset(ctx, Id)

	assert.NoError(t, issue)
	assert.Equal(t, &asset, DBasset)

}

func TestAssetService_ListAllAsset_Success(t *testing.T) {
	ctx := context.Background()
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	fl, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	dat, _ := time.Parse(layout, str)
	cost, _ := strconv.ParseFloat("5000", 32)
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, _ := json.Marshal(m)

	asset := []domain.Asset{
		{
			Id:             fl,
			Status:         "active",
			Category:       "Laptop",
			PurchaseAt:     dat,
			PurchaseCost:   cost,
			Name:           "Dell Latitude E5550",
			Specifications: b,
		},
	}

	mockAssetRepo.On("ListAssets", ctx).Return(asset, nil)

	assetService := service.NewAssetService(mockAssetRepo)
	DBasset, issue := assetService.ListAssets(ctx)

	assert.NoError(t, issue)
	assert.Equal(t, asset, DBasset)
}
func TestAssetService_DeleteAsset_When_DeleteAssetReturnsError(t *testing.T) {
	ctx := context.Background()
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(nil, errors.New("Some DB Error"))

	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.DeleteAsset(ctx, Id)

	assert.Error(t, err)
	assert.Equal(t, "Some DB Error", err.Error())
	assert.Nil(t, asset)

}
func TestAssetService_DeleteAsset_When_DeleteAssetReturnsAlready(t *testing.T) {
	ctx := context.Background()
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(nil, customerrors.AssetAlreadyDeleted)

	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.DeleteAsset(ctx, Id)

	assert.Error(t, err)
	assert.Equal(t, "Asset Already Deleted", err.Error())
	assert.Nil(t, asset)

}

func TestAssetService_DeleteAsset_When_DeleteAssetReturnsNil(t *testing.T) {
	ctx := context.Background()
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")

	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(nil, nil)
	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.DeleteAsset(ctx, Id)

	assert.Nil(t, asset)
	assert.NotNil(t, err)

}
func TestAssetService_UpdateAsset_When_ReturnsError(t *testing.T) {
	ctx := context.Background()
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	Status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, _ := json.Marshal(m)
	Specifications := b
	req := contract.UpdateRequest{
		Status:         &Status,
		Specifications: Specifications,
	}

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(nil, errors.New("some DB error"))

	AssetService := service.NewAssetService(mockAssetRepo)
	asset, err := AssetService.UpdateAsset(ctx, Id, req)
	assert.Nil(t, asset)
	assert.Error(t, err)
	assert.Equal(t, "some DB error", err.Error())

}
func TestAssetService_UpdateAsset_When_ReturnsAlreadyDeleted(t *testing.T) {
	ctx := context.Background()
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	mockAssetRepo := &mockRepo.MockAssetRepo{}
	Status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, _ := json.Marshal(m)
	Specifications := b
	req := contract.UpdateRequest{
		Status:         &Status,
		Specifications: Specifications,
	}
	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(nil, customerrors.AssetAlreadyDeleted)

	AssetService := service.NewAssetService(mockAssetRepo)
	asset, err := AssetService.UpdateAsset(ctx, Id, req)
	assert.Nil(t, asset)
	assert.Error(t, err)
	assert.Equal(t, "Asset Already Deleted", err.Error())

}
func TestAssetService_ListAllAsset_When_ListAssetReturnsError(t *testing.T) {
	ctx := context.Background()

	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("ListAssets", ctx).Return(nil, errors.New("No Asset Exists"))
	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.ListAssets(ctx)

	assert.Error(t, err)
	assert.Equal(t, "No Asset Exists", err.Error())
	assert.Nil(t, asset)
}
func TestAssetService_ListAllAsset_When_ListAssetReturnsNil(t *testing.T) {
	ctx := context.Background()

	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("ListAssets", ctx).Return(nil, nil)
	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.ListAssets(ctx)

	assert.Nil(t, asset)
	assert.NotNil(t, err)

}
func TestAssetService_UpdateAsset_When_ReturnsNil(t *testing.T) {
	ctx := context.Background()
	Id, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	Status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, _ := json.Marshal(m)
	Specifications := b
	req := contract.UpdateRequest{
		Status:         &Status,
		Specifications: Specifications,
	}

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(nil, nil)

	AssetService := service.NewAssetService(mockAssetRepo)
	asset, err := AssetService.UpdateAsset(ctx, Id, req)

	assert.Nil(t, asset)
	assert.NotNil(t, err)

}
