package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/handler"
	mockService "github.com/vkhichar/asset-management/service/mocks"
)

func TestAssetHandler_CreateAssetHandler_When_CreateAssetReturnsError(t *testing.T) {
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
	req, err := http.NewRequest("GET", "/assets", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	expectedErr := string(`{"error":"some error"}`)
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("CreateAsset", ctx, obj).Return(nil, errors.New("invalid request"))

	resp := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateAssetHandler(mockAssetService))
	handler.ServeHTTP(resp, req)

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
	req, err := http.NewRequest("GET", "/assets", requestReader)
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
