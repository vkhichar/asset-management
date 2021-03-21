package handler_test

import (
	"bytes"
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

func TestAssetHandler_UpdateAssets_When_ReturnsError(t *testing.T) {

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing %s", errParse.Error())
	}
	Id := fl
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error While Marshal %s", errMarshal.Error())
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

func TestAssetHandler_UpdateAssets_When_ReturnsNoAssetExist(t *testing.T) {

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing %s", errParse.Error())
	}
	Id := fl
	status := "active"
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Generation"] = "i8"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error While Marshal %s", errMarshal.Error())
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
	mockAssetService.On("UpdateAsset", mock.Anything, Id, assetReq).Return(nil, customerrors.NoAssetsExist)

	r := mux.NewRouter()
	r.HandleFunc("/assets/{Id}", handler.UpdateAssetHandler(mockAssetService)).Methods("PUT")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"no asset found"}`)

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
		fmt.Printf("Error While Parsing %s", errParseFloat.Error())
	}
	layout := "2006-01-02T15:04:05.000Z"
	str := "2020-07-01T00:00:00Z"
	dat, errParseDate := time.Parse(layout, str)
	if errParseDate != nil {
		fmt.Printf("Error While Parsing %s", errParseDate.Error())
	}
	cost, errParseFloat := strconv.ParseFloat("500", 32)
	if errParseFloat != nil {
		fmt.Printf("Error While Parsing %s", errParseFloat.Error())
	}
	m := make(map[string]interface{})
	m["RAM"] = "4GB"
	m["HDD"] = "500GB"
	m["Genration"] = "i8"
	b, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error While Marshal %s", errMarshal.Error())
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
		fmt.Printf("Error While Marshal %s", errMarshal.Error())
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
		fmt.Printf("Error While Parsing %s", errParsing.Error())
	}
	Id := fl
	status := "retired"
	m := make(map[string]interface{})
	m["RAM"] = "8GB"
	m["HDD"] = "1TB"
	m["Generation"] = "i8"
	ss, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		fmt.Printf("Error While Parsing %s", errMarshal.Error())
	}
	specifications := ss
	body := fmt.Sprintf(`{"status":"retired","specifications":{"Generation":"i8","HDD":"1TB","RAM":"8GB"}}`)
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
	expectedout := `{"success":"Update Operation Done Successfully"}`
	assert.JSONEq(t, expectedout, rr.Body.String())

}

func TestAssetHandler_Delete_When_ReturnsError(t *testing.T) {

	fl, errParse := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParse != nil {
		fmt.Printf("Error While Parsing %s", errParse.Error())
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
		fmt.Printf("Error While Parsing%s", errParse.Error())
	}
	Id := fl

	req, err := http.NewRequest("DELETE", "/assets/delete/ffb4b1a4-7bf5-11ee-9339-0242ac130002", nil)
	if err != nil {
		t.Fatal(err)
	}
	expect := `{"success":"Delete Operation Done Successfully"}`

	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", mock.Anything, Id).Return(nil, nil)

	r := mux.NewRouter()
	r.HandleFunc("/assets/delete/{Id}", handler.DeleteAssetHandler(mockAssetService)).Methods("DELETE")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expect, rr.Body.String())

}

func TestAssetHandler_Delete_When_ReturnsNil(t *testing.T) {

	fl, errParsing := uuid.Parse("ffb4b1a4-7bf5-11ee-9339-0242ac130002")
	if errParsing != nil {
		fmt.Printf("Error While Parsing %s", errParsing.Error())
	}
	Id := fl

	req, err := http.NewRequest("DELETE", "/assets/delete/ffb4b1a4-7bf5-11ee-9339-0242ac130002", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockAssetService := &mockService.MockAssetService{}
	mockAssetService.On("DeleteAsset", mock.Anything, Id).Return(nil, customerrors.NoAssetsExist)

	r := mux.NewRouter()
	r.HandleFunc("/assets/delete/{Id}", handler.DeleteAssetHandler(mockAssetService)).Methods("DELETE")
	r.ServeHTTP(rr, req)

	expectedErr := string(`{"error":"no asset found"}`)

	assert.JSONEq(t, expectedErr, rr.Body.String())

}
