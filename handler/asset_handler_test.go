package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/handler"
	mockService "github.com/vkhichar/asset-management/service/mocks"
)

func TestCreateAssetHandler_When_InvalidRequest(t *testing.T) {
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
	requestByte, _ := json.Marshal(obj)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/assets", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	expectedErr := string(`{"error":"invalid request"}`)
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("CreateAsset", ctx, &obj).Return(nil, errors.New("invalid request"))

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateAssetHandler(mockAssetService))
	handler.ServeHTTP(resp, req)
	//fmt.Println(resp.Body.String())

	assert.JSONEq(t, expectedErr, resp.Body.String())
}

func TestCreateAssetHandler_When_BadRequestError(t *testing.T) {
	ctx := context.Background()
	obj := domain.Asset{
		Id:             uuid.New(),
		Status:         "ihvih",
		Category:       "Laptops",
		PurchaseAt:     time.Now(),
		PurchaseCost:   50000.00,
		Name:           "aspire-5",
		Specifications: json.RawMessage{},
	}
	requestByte, _ := json.Marshal(obj)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/assets", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	expectedErr := string(`{"error":"invalid request"}`)
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("CreateAsset", ctx, obj).Return(nil, nil)

	resp := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/assets", handler.CreateAssetHandler(mockAssetService)).Methods("POST")
	r.ServeHTTP(resp, req)

	assert.JSONEq(t, expectedErr, resp.Body.String())
}
func TestAssetHandler_CreateAssetHandler_When_Success(t *testing.T) {
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
	requestByte, _ := json.Marshal(obj)
	requestReader := bytes.NewReader(requestByte)
	req, err := http.NewRequest("POST", "/assets", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("CreateAsset", ctx, obj).Return(&obj, nil)

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateAssetHandler(mockAssetService))
	handler.ServeHTTP(resp, req)

	assert.NotNil(t, &obj)
	assert.NoError(t, err)
}

func TestGetAssetHandler_When_ReturnsInvalidUUID(t *testing.T) {
	ctx := context.Background()
	id, err := uuid.Parse("642fc397-abec-4e1e-8473-69803dbb901")
	if err != nil {
		fmt.Printf("%s:", err.Error())
		return

	}
	req, err := http.NewRequest("GET", "/assets/642fc397-abec-4e1e-8473-69803dbb901", nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedErr := string(`{"error":"Invalid UUID format"}`)
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("GetAsset", ctx, id).Return(nil, errors.New("Invalid UUID format"))

	resp := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/assets/{id}", handler.GetAssetHandler(mockAssetService)).Methods("GET")
	r.ServeHTTP(resp, req)

	fmt.Println(resp.Body.String())
	assert.JSONEq(t, expectedErr, resp.Body.String())
}

func TestGetAssetHandler_When_AssetDoesNotExist(t *testing.T) {
	//ctx := context.Background()
	id, err := uuid.Parse("69c9886b-fdf7-4434-8c28-2639d4122871")
	if err != nil {
		fmt.Printf("%s:", err.Error())
		return

	}
	req, err := http.NewRequest("GET", "/assets/69c9886b-fdf7-4434-8c28-2639d4122871", nil)
	if err != nil {
		t.Fatal(err)
	}

	expectedErr := string(`{"error":"No assets exist"}`)
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("GetAsset", mock.Anything, id).Return(nil, customerrors.NoAssetsExist)

	resp := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/assets/{id}", handler.GetAssetHandler(mockAssetService)).Methods("GET")
	r.ServeHTTP(resp, req)

	fmt.Println(resp.Body.String())
	assert.JSONEq(t, expectedErr, resp.Body.String())
}

func TestAssetHandler_GetAssetHandler_When_Success(t *testing.T) {
	id, _ := uuid.Parse("642fc397-abec-4e1e-8473-69803dbb9016")
	duration, _ := time.Parse("01/01/0001", "01/01/0001")
	specs := []byte(`{"ram":"4GB","brand":"acer"}`)

	obj := domain.Asset{
		Id:             id,
		Status:         "active",
		Category:       "laptop",
		PurchaseAt:     duration,
		PurchaseCost:   45000.00,
		Name:           "aspire-5",
		Specifications: specs,
	}

	expectedObj, err := json.Marshal(contract.GetAssetResponse{
		ID:             id,
		Status:         "active",
		Category:       "laptop",
		PurchaseAt:     duration.String(),
		PurchaseCost:   45000.00,
		Name:           "aspire-5",
		Specifications: specs,
	})

	req, err := http.NewRequest("GET", "/assets/642fc397-abec-4e1e-8473-69803dbb9016", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("GetAsset", mock.Anything, id).Return(&obj, nil)

	resp := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/assets/{id}", handler.GetAssetHandler(mockAssetService)).Methods("GET")
	r.ServeHTTP(resp, req)

	assert.NotNil(t, &obj)
	assert.JSONEq(t, string(expectedObj), resp.Body.String())
}
