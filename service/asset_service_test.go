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
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/domain"
	mockRepo "github.com/vkhichar/asset-management/repository/mocks"
	"github.com/vkhichar/asset-management/service"
	mockEventSvc "github.com/vkhichar/asset-management/service/mocks"
	"gopkg.in/h2non/gock.v1"
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
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("CreateAsset", ctx, obj).Return(nil, errors.New("some db error"))

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := assetService.CreateAsset(ctx, &obj)

	assert.Error(t, err)
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
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("CreateAsset", ctx, obj).Return(&obj, nil)

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	dbAsset, err := assetService.CreateAsset(ctx, &obj)

	assert.NoError(t, err)
	assert.Equal(t, &obj, dbAsset)
}

func TestAssetService_GetAsset_When_GetAssetReturnsError(t *testing.T) {
	ctx := context.Background()
	Id := uuid.New()

	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventSvc := &mockEventSvc.MockEventService{}
	mockAssetRepo.On("GetAsset", ctx, Id).Return(nil, errors.New("some db error"))

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := assetService.GetAsset(ctx, Id)

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, asset)
}

func TestAssetService_GetAsset_When_Success(t *testing.T) {
	ctx := context.Background()
	id, err := uuid.Parse("ca99be07-30ef-4853-a6b6-bfa38d254a29")
	if err != nil {
		fmt.Printf("asset_servie_test: error while parsing string into uuid")
		return
	}

	dummy := domain.Asset{
		Id:             id,
		Status:         "inactive",
		Category:       "laptops",
		PurchaseAt:     time.Now(),
		PurchaseCost:   45000.00,
		Name:           "aspire-5",
		Specifications: json.RawMessage{},
	}

	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventSvc := &mockEventSvc.MockEventService{}
	mockAssetRepo.On("GetAsset", ctx, id).Return(&dummy, nil)

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := assetService.GetAsset(ctx, id)

	fmt.Println(asset)
	fmt.Println()

	assert.NoError(t, err)
	assert.Equal(t, &dummy, asset)
	assert.NotNil(t, asset)
}

func TestAssetService_CreateAsset_When_PostAssetEventReturnsEventID(t *testing.T) {
	ctx := context.Background()
	eventID := "120"
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
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("CreateAsset", ctx, obj).Return(&obj, nil)
	mockEventSvc.On("PostAssetEvent", ctx, &obj).Return(eventID, nil)

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	dbAsset, err := assetService.CreateAsset(ctx, &obj)

	assert.NoError(t, err)
	assert.Equal(t, &obj, dbAsset)
}

func TestAssetService_CreateAsset_When_PostAssetEventReturnsError(t *testing.T) {
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
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("CreateAsset", ctx, &obj).Return(nil, errors.New("some db error"))
	mockEventSvc.On("PostAssetEvent", ctx, nil).Return("", errors.New("Error during Event"))

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	dbAsset, err := assetService.CreateAsset(ctx, &obj)
	fmt.Printf(err.Error())

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, dbAsset)
}

func TestAssetService_When_PostAssetEventSuccess(t *testing.T) {
	ctx := context.Background()

	gock.New("http://34.70.86.33:" + config.GetEventPort()).
		Post("/events").
		Reply(200).
		JSON(map[string]string{"id": "123"})

	m := make(map[string]interface{})
	m["ram"] = "4GB"
	m["brand"] = "acer"
	b, _ := json.Marshal(m)

	obj := domain.Asset{
		Id:             uuid.New(),
		Status:         "retired",
		Category:       "Laptops",
		PurchaseAt:     time.Now(),
		PurchaseCost:   50000.00,
		Name:           "aspire-5",
		Specifications: b,
	}

	eventSvc := service.NewEventService()
	id, err := eventSvc.PostAssetEvent(ctx, &obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"id": "123"}`, id)
}

func TestAssetService_When_PostAssetEventReturnsError(t *testing.T) {
	ctx := context.Background()

	gock.New("http://34.70.86.33:" + config.GetEventPort()).Post("/events").Reply(400)

	m := make(map[string]interface{})
	m["ram"] = "4GB"
	m["brand"] = "acer"
	b, _ := json.Marshal(m)

	obj := domain.Asset{
		Id:             uuid.New(),
		Status:         "retired",
		Category:       "Laptops",
		PurchaseAt:     time.Now(),
		PurchaseCost:   50000.00,
		Name:           "aspire-5",
		Specifications: b,
	}

	eventSvc := service.NewEventService()
	id, err := eventSvc.PostAssetEvent(ctx, &obj)

	assert.Nil(t, err)
	assert.Equal(t, "", id)
}
