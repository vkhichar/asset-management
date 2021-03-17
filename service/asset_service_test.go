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
	mockService "github.com/vkhichar/asset-management/service/mocks"
	"gopkg.in/h2non/gock.v1"
)

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
	eventId := "123"
	Specifications := jsr
	Status := "active"
	req := contract.UpdateRequest{
		Status:         &Status,
		Specifications: Specifications,
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventService := &mockService.MockEventService{}

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(&asset, nil)
	mockEventService.On("PostAssetEvent", ctx, &asset).Return(eventId, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventService)
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
	mockEventService := &mockService.MockEventService{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(&asset, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventService)
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
	mockEventService := &mockService.MockEventService{}

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

	assetService := service.NewAssetService(mockAssetRepo, mockEventService)
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
	mockEventService := &mockService.MockEventService{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(nil, errors.New("Some DB Error"))

	assetService := service.NewAssetService(mockAssetRepo, mockEventService)
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
	mockEventService := &mockService.MockEventService{}

	mockAssetRepo.On("DeleteAsset", ctx, Id).Return(nil, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventService)
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
	mockEventService := &mockService.MockEventService{}

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
	mockEventService.On("")

	AssetService := service.NewAssetService(mockAssetRepo, mockEventService)
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
	mockEventService := &mockService.MockEventService{}

	mockAssetRepo.On("ListAssets", ctx).Return(nil, errors.New("No Asset Exists"))
	assetService := service.NewAssetService(mockAssetRepo, mockEventService)
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
	mockEventService := &mockService.MockEventService{}

	mockAssetRepo.On("ListAssets", ctx).Return(nil, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventService)
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
	mockEventService := &mockService.MockEventService{}

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

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(nil, nil)

	AssetService := service.NewAssetService(mockAssetRepo, mockEventService)
	asset, err := AssetService.UpdateAsset(ctx, Id, req)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}

	assert.Nil(t, asset)
	assert.NotNil(t, err)

}
func TestAssetService_UpdateAsset_When_PostAssetReturnSucess(t *testing.T) {
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
	eventId := "112"
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
	mockEventService := &mockService.MockEventService{}

	mockAssetRepo.On("UpdateAsset", ctx, Id, req).Return(&asset, nil)
	mockEventService.On("PostAssetEvent", ctx, &asset).Return(eventId, nil)
	assetService := service.NewAssetService(mockAssetRepo, mockEventService)
	DBasset, err := assetService.UpdateAsset(ctx, Id, req)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}

	assert.Nil(t, err)
	assert.Equal(t, &asset, DBasset)

}
func TestAssetService_UpdateAsset_When_PostAssetEventReturnError(t *testing.T) {
	ctx := context.Background()
	Id, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing String to UUID %s", errParse.Error())
	}
	mockAssetRepo := &mockRepo.MockAssetRepo{}
	mockEventService := &mockService.MockEventService{}

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
	mockEventService.On("PostAssetEvent", ctx, nil).Return("", errors.New("Error in PostAssetEvent"))

	AssetService := service.NewAssetService(mockAssetRepo, mockEventService)
	asset, err := AssetService.UpdateAsset(ctx, Id, req)
	if err != nil {
		fmt.Printf("Something went Wrong %s", err.Error())
	}
	assert.Nil(t, asset)
	assert.Error(t, err)
	assert.Equal(t, "some DB error", err.Error())

}
func TestAssetService_PostAssetEvent_When_Success(t *testing.T) {
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
	defer gock.Off()

	gock.New("http://34.70.86.33:" + config.GetEventAppPort()).
		Post("/events").
		Reply(200).
		JSON(map[string]string{"id": "120"})
	asset := domain.Asset{

		Id:             Id,
		Status:         "active",
		Category:       "Laptop",
		PurchaseAt:     dat,
		PurchaseCost:   cost,
		Name:           "Dell Latitude E5550",
		Specifications: b,
	}

	eventService := service.NewEventService()
	id, err := eventService.PostAssetEvent(ctx, &asset)

	assert.Nil(t, err)
	assert.JSONEq(t, `{"id":"120"}`, id)

}
func TestAssetService_PostAssetEvent_When_ReturnsError(t *testing.T) {
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
	defer gock.Off()

	gock.New("http://34.70.86.33:" + config.GetEventAppPort()).
		Post("/events").
		Reply(400)
	asset := domain.Asset{

		Id:             Id,
		Status:         "active",
		Category:       "Laptop",
		PurchaseAt:     dat,
		PurchaseCost:   cost,
		Name:           "Dell Latitude E5550",
		Specifications: b,
	}

	eventService := service.NewEventService()
	id, err := eventService.PostAssetEvent(ctx, &asset)

	assert.Equal(t, "", id)
	assert.Nil(t, err)

}
