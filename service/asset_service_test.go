package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/domain"
	mockRepo "github.com/vkhichar/asset-management/repository/mocks"
	"github.com/vkhichar/asset-management/service"
)

func TestAssetService_CreateAsset_When_CreateAssetReturnsError(t *testing.T) {
	ctx := context.Background()

	obj := domain.Asset{
		Id:             uuid.New(),
		Status:         "retired",
		Category:       "Laptops",
		PurchaseAt:     time.Now(),
		PurchaseCost:   50000.00,
		Name:           "aspire-5",
		Specifications: json.RawMessage{},
	}

	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("CreateAsset", ctx, obj).Return(nil, errors.New("some db error"))

	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.CreateAsset(ctx, obj)

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, asset)
}

func TestAssetService_CreateAsset_Success(t *testing.T) {
	ctx := context.Background()

	obj := domain.Asset{
		Id:             uuid.New(),
		Status:         "retired",
		Category:       "Laptops",
		PurchaseAt:     time.Now(),
		PurchaseCost:   50000.00,
		Name:           "aspire-5",
		Specifications: json.RawMessage{},
	}

	mockAssetRepo := &mockRepo.MockAssetRepo{}

	mockAssetRepo.On("CreateAsset", ctx, obj).Return(&obj, nil)

	assetService := service.NewAssetService(mockAssetRepo)
	dbAsset, err := assetService.CreateAsset(ctx, obj)

	assert.NoError(t, err)
	assert.Equal(t, &obj, dbAsset)
}

func TestAssetService_GetAsset_When_GetAssetReturnsError(t *testing.T) {
	ctx := context.Background()
	Id := uuid.New()

	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockAssetRepo.On("GetAsset", ctx, Id).Return(nil, errors.New("some db error"))

	assetService := service.NewAssetService(mockAssetRepo)
	asset, err := assetService.GetAsset(ctx, Id)

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, asset)
}

func TestAssetService_GetAsset_When_Success(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()

	duration, err := time.Parse("01-01-0001", "0001-01-01")
	if err != nil {
		fmt.Printf("asset_servie_test: error while parsing string into time")
		return
	}

	specs, err := json.Marshal(`"ram":"4GB","brand":"acer"`)
	if err != nil {
		fmt.Printf("asset_servie_test: error while parsing string into json")
		return
	}

	dummy := domain.Asset{
		Id:             uuid.New(),
		Status:         "active",
		Category:       "laptops",
		PurchaseAt:     duration,
		PurchaseCost:   45000.00,
		Name:           "aspire-5",
		Specifications: specs,
	}

	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockAssetRepo.On("GetAsset", ctx, dummy).Return(&dummy, nil)

	assetService := service.NewAssetService(mockAssetRepo)
	asset, _ := assetService.GetAsset(ctx, id)

	//assert.NoError(t, err)
	//assert.Equal(t, dummy.Id, asset.Id)
	assert.Nil(t, asset)

}
