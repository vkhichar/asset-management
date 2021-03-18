package service_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
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

	mockAssetRepo.On("CreateAsset", ctx, &obj).Return(nil, errors.New("some db error"))

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

	mockAssetRepo.On("CreateAsset", ctx, &obj).Return(&obj, nil)
	mockEventSvc.On("PostAssetEventCreateAsset", ctx, &obj).Return("12", nil)

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

func TestAssetService_CreateAsset_When_PostAssetEventCreateAssetReturnsEventID(t *testing.T) {
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

	mockAssetRepo.On("CreateAsset", ctx, &obj).Return(&obj, nil)
	mockEventSvc.On("PostAssetEventCreateAsset", ctx, &obj).Return(eventID, nil)

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	dbAsset, err := assetService.CreateAsset(ctx, &obj)

	assert.NoError(t, err)
	assert.Equal(t, &obj, dbAsset)
}

func TestAssetService_CreateAsset_When_PostAssetEventCreateAssetReturnsError(t *testing.T) {
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
	mockEventSvc.On("PostAssetEventCreateAsset", ctx, nil).Return("", errors.New("Error during Event"))

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	dbAsset, err := assetService.CreateAsset(ctx, &obj)
	fmt.Printf(err.Error())

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, dbAsset)
}

func TestAssetService_When_PostAssetEventSuccess(t *testing.T) {
	ctx := context.Background()

	gock.New(config.GetEventServiceUrl()).
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
	id, err := eventSvc.PostAssetEventCreateAsset(ctx, &obj)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"id": "123"}`, id)
}

func TestAssetService_When_PostAssetEventReturnsError(t *testing.T) {
	ctx := context.Background()

	gock.New(config.GetEventServiceUrl()).Post("/events").Reply(400)

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
	id, err := eventSvc.PostAssetEventCreateAsset(ctx, &obj)

	assert.Nil(t, err)
	assert.Equal(t, "", id)
}

func TestAssetService_UpdateAsset_When_Success(t *testing.T) {
	ctx := context.Background()
	Id, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing String to UUID %s", errParse.Error())
	}
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	dat, errParseDate := time.Parse(layout, str)
	if errParseDate != nil {
		fmt.Printf("Error While Parsing %s", errParseDate.Error())
	}
	cost, errParseFloat := strconv.ParseFloat("5000", 32)
	if errParseFloat != nil {
		fmt.Printf("Error While Parsing %s", errParseFloat.Error())
	}
	m := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error while Marshaling %s", errMarshal.Error())
	}
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
	jsr, errMarshal := json.Marshal(js)
	if errMarshal != nil {
		fmt.Printf("Error While Marshaling %s", errMarshal.Error())
	}
	Specifications := jsr
	Status := "active"
	req := contract.UpdateRequest{
		Status:         &Status,
		Specifications: Specifications,
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(&asset, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	DBasset, err := assetService.UpdateAsset(ctx, Id, req)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}

	assert.Nil(t, err)
	assert.Equal(t, &asset, DBasset)

}
func TestAssetService_DeleteAsset_Success(t *testing.T) {
	ctx := context.Background()
	Id, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing String to UUID %s", errParse.Error())
	}
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	dat, errParseDate := time.Parse(layout, str)
	if errParseDate != nil {
		fmt.Printf("Error While Parsing %s", errParseDate.Error())
	}
	cost, errParseFloat := strconv.ParseFloat("5000", 32)
	if errParseFloat != nil {
		fmt.Printf("Error While Parsing %s", errParseFloat.Error())
	}
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error While Marshaling %s", errMarshal.Error())
	}

	asset := domain.Asset{

		Id:           Id,
		Status:       "active",
		Category:     "Laptop",
		PurchaseAt:   dat,
		PurchaseCost: cost,

		Name:           "Dell Latitude E5550",
		Specifications: b,
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(&asset, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	DBasset, err := assetService.DeleteAsset(ctx, Id)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}
	assert.NoError(t, err)
	assert.Equal(t, &asset, DBasset)

}

func TestAssetService_ListAllAsset_Success(t *testing.T) {
	ctx := context.Background()
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing String to UUID")
	}
	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	dat, errParseDate := time.Parse(layout, str)
	if errParseDate != nil {
		fmt.Printf("Error While Parsing")
	}
	cost, errParseFloat := strconv.ParseFloat("5000", 32)
	if errParseFloat != nil {
		fmt.Printf("Error While Parsing %s", errParseFloat.Error())
	}
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error While Marshaling %s", errMarshal.Error())
	}

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
	mockEventSvc := &mockEventSvc.MockEventService{}

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	DBasset, err := assetService.ListAssets(ctx)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}

	assert.NoError(t, err)
	assert.Equal(t, asset, DBasset)
}
func TestAssetService_DeleteAsset_When_DeleteAssetReturnsError(t *testing.T) {
	ctx := context.Background()
	Id, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing String to UUID %s", errParse.Error())
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(nil, errors.New("Some DB Error"))

	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := assetService.DeleteAsset(ctx, Id)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}
	assert.Error(t, err)
	assert.Equal(t, "Some DB Error", err.Error())
	assert.Nil(t, asset)

}

func TestAssetService_DeleteAsset_When_DeleteAssetReturnsNil(t *testing.T) {
	ctx := context.Background()
	Id, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing String to UUID %s", errParse.Error())
	}

	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(nil, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := assetService.DeleteAsset(ctx, Id)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}

	assert.Nil(t, asset)
	assert.NotNil(t, err)

}
func TestAssetService_UpdateAsset_When_ReturnsError(t *testing.T) {
	ctx := context.Background()
	Id, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing String to UUID %s", errParse.Error())
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	Status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error While Marshaling %s", errMarshal.Error())
	}
	Specifications := b
	req := contract.UpdateRequest{
		Status:         &Status,
		Specifications: Specifications,
	}

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(nil, errors.New("some DB error"))
	mockEventSvc := &mockEventSvc.MockEventService{}

	AssetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := AssetService.UpdateAsset(ctx, Id, req)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}
	assert.Nil(t, asset)
	assert.Error(t, err)
	assert.Equal(t, "some DB error", err.Error())

}

func TestAssetService_ListAllAsset_When_ListAssetReturnsError(t *testing.T) {
	ctx := context.Background()

	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("ListAssets", ctx).Return(nil, errors.New("No Asset Exists"))
	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := assetService.ListAssets(ctx)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}

	assert.Error(t, err)
	assert.Equal(t, "No Asset Exists", err.Error())
	assert.Nil(t, asset)
}
func TestAssetService_ListAllAsset_When_ListAssetReturnsNil(t *testing.T) {
	ctx := context.Background()

	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("ListAssets", ctx).Return(nil, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := assetService.ListAssets(ctx)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}

	assert.Nil(t, asset)
	assert.NotNil(t, err)

}
func TestAssetService_UpdateAsset_When_ReturnsNil(t *testing.T) {
	ctx := context.Background()
	Id, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing String to UUID %s", errParse.Error())
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}

	Status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "1TB"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error While Marshaling %s", errMarshal.Error())
	}
	Specifications := b
	req := contract.UpdateRequest{
		Status:         &Status,
		Specifications: Specifications,
	}
	mockEventSvc := &mockEventSvc.MockEventService{}

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(nil, nil)

	AssetService := service.NewAssetService(mockAssetRepo, mockEventSvc)
	asset, err := AssetService.UpdateAsset(ctx, Id, req)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}

	assert.Nil(t, asset)
	assert.NotNil(t, err)
}
