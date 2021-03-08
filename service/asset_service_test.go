package service_test

import (
	"context"
	"encoding/json"
	"errors"
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
