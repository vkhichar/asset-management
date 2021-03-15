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

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Println("Error While Parsing")
	}
	Id := fl
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Println("Error While Marshal")
	}
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
	mockAssetService.On("UpdateAsset", mock.Anything, Id, assetReq).Return(nil, errors.New("something went wrong"))

	r := mux.NewRouter()
	r.HandleFunc("/assets/{Id}", handler.UpdateAssetHandler(mockAssetService)).Methods("PUT")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"something went wrong"}`)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}

func TestAssetHandler_UpdateAssets_When_ReturnsNil(t *testing.T) {

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Println("Error While Parsing")
	}
	Id := fl
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Println("Error While Marshal")
	}
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

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Println("Error While Parsing")
	}
	Id := fl
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Println("Error While MArshaling")
	}
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

	fl, errParseFloat := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParseFloat != nil {
		fmt.Println("Error While Parsing")
	}
	layout := "2006-01-02T15:04:05.000Z"
	str := "2020-07-01T00:00:00Z"
	dat, errParseDate := time.Parse(layout, str)
	if errParseDate != nil {
		fmt.Println("Error While Parsing")
	}
	cost, errParseFloat := strconv.ParseFloat("500", 32)
	if errParseFloat != nil {
		fmt.Println("Error While Parsing")
	}
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Genration"] = "i8"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Println("Error While Marshal")
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
	expect, errMarshal := json.Marshal(expectedasset)
	if errMarshal != nil {
		fmt.Println("Error While Marshal")
	}
	expectedout := string(expect)

	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("ListAssets", ctx).Return(asset, nil)

	handlerTest := http.HandlerFunc(handler.ListAssetHandler(mockAssetService))
	handlerTest.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedout, rr.Body.String())

}

func TestAssetHandler_UpdateAssets_When_Success(t *testing.T) {

	fl, errParsing := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParsing != nil {
		fmt.Println("Error While Parsing")
	}
	Id := fl
	status := "retired"
	m := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	m["Generation"] = "i8"
	ss, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Println("Error While Parsing")
	}
	specifications := ss
	td := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	m["Generation"] = "i8"

	TD, errMarshal := json.Marshal(td)
	if errMarshal != nil {
		fmt.Println("Error While Marshal")
	}

	layout := "2006-01-02T15:04:05.000Z"
	str := "2020-07-01T00:00:00Z"
	dat, errParseDate := time.Parse(layout, str)
	if errParseDate != nil {
		fmt.Println("Error While Marshal")
	}
	cost, errParseFloat := strconv.ParseFloat("500", 32)
	if errParseFloat != nil {
		fmt.Println("Error While Parsing float")
	}

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
	expect, errMarshal := json.Marshal(expectedasset)
	if errMarshal != nil {
		fmt.Println("Error While Marshaling")
	}
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

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Println("Error While Parsing")
	}
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

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Println("Error While Parsing")
	}
	Id := fl

	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	ss, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Println("Error While Marshaling")
	}
	layout := "2006-01-02T15:04:05.000Z"
	str := "2020-07-01T00:00:00Z"
	dat, errParseDate := time.Parse(layout, str)
	if errParseDate != nil {
		fmt.Println("Error While Marshal")
	}
	cost, errParseFloat := strconv.ParseFloat("500", 32)
	if errParseFloat != nil {
		fmt.Println("Error While Parsing")
	}

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
	expect, errMarshal := json.Marshal(expectedasset)
	if errMarshal != nil {
		fmt.Println("Error While Marshal")
	}
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

	fl, errParsing := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParsing != nil {
		fmt.Println("Error While Parsing")
	}
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

	fl, errParsing := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParsing != nil {
		fmt.Println("Error While Parsing")
	}
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
