package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
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

func TestAssetHandler_UpdateAssets_When_ReturnsError(t *testing.T) {

	fl, err := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	Id := fl
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	b, _ := json.Marshal(m)
	specifications := b

	body := fmt.Sprintf(`{"status" :"active","specifications": {"Generation":"i8","HDD":"500GB","RAM":"4GB"}}`)
	req, err := http.NewRequest("PUT", "/assets/ffb4b1a4-7bf5-11ee-9339-0242ac130002", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	assetReq := contract.UpdateRequest{
		Status:         &status,
		Specifications: specifications,
	}
	var e1 map[string]interface{}
	fault := json.Unmarshal([]byte(assetReq.Specifications), &e1)
	fmt.Println("Test:", e1)
	fmt.Println(fault)
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("UpdateAsset", mock.Anything, Id, assetReq).Return(nil, errors.New("something went wrong"))

	r := mux.NewRouter()
	r.HandleFunc("/assets/{Id}", handler.UpdateAssetHandler(mockAssetService)).Methods("PUT")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"something went wrong"}`)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}

func TestAssetHandler_UpdateAssets_When_ReturnsNil(t *testing.T) {

	fl, err := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	Id := fl
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	b, _ := json.Marshal(m)
	specifications := b

	body := fmt.Sprintf(`{"status" :"active","specifications": {"Generation":"i8","HDD":"500GB","RAM":"4GB"}}`)
	req, err := http.NewRequest("PUT", "/assets/ffb4b1a4-7bf5-11ee-9339-0242ac130002", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	assetReq := contract.UpdateRequest{
		Status:         &status,
		Specifications: specifications,
	}
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("UpdateAsset", mock.Anything, Id, assetReq).Return(nil, nil)

	r := mux.NewRouter()
	r.HandleFunc("/assets/{Id}", handler.UpdateAssetHandler(mockAssetService)).Methods("PUT")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"no asset found"}`)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}
func TestAssetHandler_UpdateAssets_When_ReturnsAlreadyDeleted(t *testing.T) {

	fl, err := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	Id := fl
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	b, _ := json.Marshal(m)
	specifications := b

	body := fmt.Sprintf(`{"status" :"active","specifications": {"Generation":"i8","HDD":"500GB","RAM":"4GB"}}`)
	req, err := http.NewRequest("PUT", "/assets/ffb4b1a4-7bf5-11ee-9339-0242ac130002", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	assetReq := contract.UpdateRequest{
		Status:         &status,
		Specifications: specifications,
	}
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("UpdateAsset", mock.Anything, Id, assetReq).Return(nil, customerrors.AssetAlreadyDeleted)

	r := mux.NewRouter()
	r.HandleFunc("/assets/{Id}", handler.UpdateAssetHandler(mockAssetService)).Methods("PUT")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"Asset Already Deleted"}`)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}

func TestAssetHandler_ListAllAssets_When_ReturnsError(t *testing.T) {
	ctx := context.Background()

	req, err := http.NewRequest("GET", "/assets", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	expectedErr := string(`{"error":"something went wrong"}`)

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("ListAssets", ctx).Return(nil, errors.New("Something went wrong"))

	handlerTest := http.HandlerFunc(handler.ListAssetHandler(mockAssetService))
	handlerTest.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}
func TestAssetHandler_ListAllAssets_When_ReturnsNil(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequest("GET", "/assets", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	expectedErr := string(`{"error":"no asset found"}`)
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("ListAssets", ctx).Return(nil, nil)

	handlerTest := http.HandlerFunc(handler.ListAssetHandler(mockAssetService))
	handlerTest.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}

func TestAssetHandler_ListAllAssets_When_Success(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequest("GET", "/assets", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	fl, _ := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	layout := "2006-01-02T15:04:05.000Z"
	str := "2020-07-01T00:00:00Z"
	dat, _ := time.Parse(layout, str)
	cost, _ := strconv.ParseFloat("500", 32)
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Genration"] = "i8"
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
	expectedasset := []contract.Asset{
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
	expect, _ := json.Marshal(expectedasset)
	expectedout := string(expect)

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("ListAssets", ctx).Return(asset, nil)

	handlerTest := http.HandlerFunc(handler.ListAssetHandler(mockAssetService))
	handlerTest.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedout, rr.Body.String())

}

func TestAssetHandler_UpdateAssets_When_Success(t *testing.T) {

	fl, err := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	Id := fl
	status := "retired"
	m := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	m["Generation"] = "i8"
	ss, _ := json.Marshal(m)
	specifications := ss
	td := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	m["Generation"] = "i8"

	TD, _ := json.Marshal(td)

	layout := "2006-01-02T15:04:05.000Z"
	str := "2020-07-01T00:00:00Z"
	dat, _ := time.Parse(layout, str)
	cost, _ := strconv.ParseFloat("500", 32)

	body := fmt.Sprintf(`{"status":"retired","specifications":{"Generation":"i8","HDD":"1TB","RAM":"8GB"}}`)
	req, err := http.NewRequest("PUT", "/assets/ffb4b1a4-7bf5-11ee-9339-0242ac130002", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	asset := domain.Asset{
		Id:             fl,
		Status:         "retired",
		Category:       "Laptop",
		PurchaseAt:     dat,
		PurchaseCost:   cost,
		Name:           "Dell Latitude E5550",
		Specifications: TD,
	}
	expectedasset := contract.Asset{
		Id:             fl,
		Status:         "retired",
		Category:       "Laptop",
		PurchaseAt:     dat,
		PurchaseCost:   cost,
		Name:           "Dell Latitude E5550",
		Specifications: TD,
	}
	expect, _ := json.Marshal(expectedasset)
	expectedout := string(expect)
	assetReq := contract.UpdateRequest{
		Status:         &status,
		Specifications: specifications,
	}
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("UpdateAsset", mock.Anything, Id, assetReq).Return(&asset, nil)

	r := mux.NewRouter()
	r.HandleFunc("/assets/{Id}", handler.UpdateAssetHandler(mockAssetService)).Methods("PUT")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedout, rr.Body.String())

}

func TestAssetHandler_Delete_When_ReturnsError(t *testing.T) {

	fl, err := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	Id := fl

	req, err := http.NewRequest("DELETE", "/assets/delete/ffb4b1a4-7bf5-11ee-9339-0242ac130002", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", mock.Anything, Id).Return(nil, errors.New("something went wrong"))

	r := mux.NewRouter()
	r.HandleFunc("/assets/delete/{Id}", handler.DeleteAssetHandler(mockAssetService)).Methods("DELETE")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"something went wrong"}`)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}

func TestAssetHandler_DeleteAssets_When_Success(t *testing.T) {

	fl, err := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	Id := fl

	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	ss, _ := json.Marshal(m)

	layout := "2006-01-02T15:04:05.000Z"
	str := "2020-07-01T00:00:00Z"
	dat, _ := time.Parse(layout, str)
	cost, _ := strconv.ParseFloat("500", 32)

	req, err := http.NewRequest("DELETE", "/assets/delete/ffb4b1a4-7bf5-11ee-9339-0242ac130002", nil)
	if err != nil {
		t.Fatal(err)
	}
	asset := domain.Asset{
		Id:             fl,
		Status:         "retired",
		Category:       "Laptop",
		PurchaseAt:     dat,
		PurchaseCost:   cost,
		Name:           "Dell Latitude E5550",
		Specifications: ss,
	}
	expectedasset := contract.Asset{
		Id:             fl,
		Status:         "retired",
		Category:       "Laptop",
		PurchaseAt:     dat,
		PurchaseCost:   cost,
		Name:           "Dell Latitude E5550",
		Specifications: ss,
	}
	expect, _ := json.Marshal(expectedasset)
	expectedout := string(expect)
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", mock.Anything, Id).Return(&asset, nil)

	r := mux.NewRouter()
	r.HandleFunc("/assets/delete/{Id}", handler.DeleteAssetHandler(mockAssetService)).Methods("DELETE")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedout, rr.Body.String())

}

func TestAssetHandler_Delete_When_ReturnsNil(t *testing.T) {

	fl, err := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	Id := fl

	req, err := http.NewRequest("DELETE", "/assets/delete/ffb4b1a4-7bf5-11ee-9339-0242ac130002", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", mock.Anything, Id).Return(nil, nil)

	r := mux.NewRouter()
	r.HandleFunc("/assets/delete/{Id}", handler.DeleteAssetHandler(mockAssetService)).Methods("DELETE")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"no asset found"}`)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}
func TestAssetHandler_Delete_When_ReturnsAlreadyDeleted(t *testing.T) {

	fl, err := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	Id := fl

	req, err := http.NewRequest("DELETE", "/assets/delete/ffb4b1a4-7bf5-11ee-9339-0242ac130002", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", mock.Anything, Id).Return(nil, customerrors.AssetAlreadyDeleted)

	r := mux.NewRouter()
	r.HandleFunc("/assets/delete/{Id}", handler.DeleteAssetHandler(mockAssetService)).Methods("DELETE")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"Asset Already Deleted"}`)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}
