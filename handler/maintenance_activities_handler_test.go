package handler_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
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

func TestMaintenanceActivitiesHandler_CreateMaintenanceHandler_When_CreateAssetMaintenanceReturnsError(t *testing.T) {

	body := fmt.Sprintf(`{"cost": 100, "started_at":"28-02-1996","description": "hardware issue"}`)
	id, _ := uuid.Parse("ffb4b1a4-7bf5-11eb-9439-0242ac130002")
	startedDate := "28-02-1996"
	tym, err := time.Parse("02-01-2006", startedDate)

	req, err := http.NewRequest("POST", "/assets/ffb4b1a4-7bf5-11eb-9439-0242ac130002/maintenance", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	inpt := domain.MaintenanceActivity{
		AssetId:     id,
		Cost:        100,
		StartedAt:   tym,
		Description: "hardware issue",
	}
	expectedErr := string(`{"error":"This asset id does not exist... "}`)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockAssetMaintenanceService.On("CreateAssetMaintenance", mock.Anything, inpt).Return(nil, errors.New("This asset id does not exist..."))
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.CreateMaintenanceHandler(mockAssetMaintenanceService)).Methods("POST")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}

func TestMaintenanceActivitiesHandler_CreateMaintenanceHandler_When_Success(t *testing.T) {

	body := fmt.Sprintf(`{"cost": 100, "started_at":"28-02-1996","description": "hardware issue"}`)
	id, _ := uuid.Parse("ffb4b1a4-7bf5-11eb-9439-0242ac130002")
	startedDate := "28-02-1996"
	tym, err := time.Parse("02-01-2006", startedDate)

	req, err := http.NewRequest("POST", "/assets/ffb4b1a4-7bf5-11eb-9439-0242ac130002/maintenance", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	inpt := domain.MaintenanceActivity{
		AssetId:     id,
		Cost:        100,
		StartedAt:   tym,
		Description: "hardware issue",
	}

	outpt := domain.MaintenanceActivity{
		ID:          1,
		AssetId:     id,
		Cost:        100,
		StartedAt:   tym,
		Description: "hardware issue",
	}
	outptt := contract.AssetMaintenanceResponse{
		ID:          1,
		AssetId:     id,
		Cost:        100,
		StartedAt:   tym.Format("02-01-2006"),
		Description: "hardware issue",
	}
	finalOutpt, _ := json.Marshal(outptt)
	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockAssetMaintenanceService.On("CreateAssetMaintenance", mock.Anything, inpt).Return(&outpt, nil)
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.CreateMaintenanceHandler(mockAssetMaintenanceService)).Methods("POST")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, string(finalOutpt), rr.Body.String())
}

func TestMaintenanceActivitiesHandler_CreateMaintenanceHandler_When_ParsingUuidReturnsError(t *testing.T) {

	body := fmt.Sprintf(`{"cost": 100, "started_at":"28-02-1996","description": "hardware issue"}`)

	req, err := http.NewRequest("POST", "/assets/ffb4b1a4-7bf5-11eb-9439-0242ac13000/maintenance", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	expectedErr := string(`{"error":"invalid asset id"}`)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.CreateMaintenanceHandler(mockAssetMaintenanceService)).Methods("POST")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}

func TestMaintenanceActivitiesHandler_CreateMaintenanceHandler_When_DecodingReturnsError(t *testing.T) {

	body := fmt.Sprintf(`{"cost": "100", "started_at":"28-02-1996","description": "hardware issue"}`)

	req, err := http.NewRequest("POST", "/assets/ffb4b1a4-7bf5-11eb-9439-0242ac130002/maintenance", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	expectedErr := string(`{"error":"invalid request"}`)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.CreateMaintenanceHandler(mockAssetMaintenanceService)).Methods("POST")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}
func TestMaintenanceActivitiesHandler_CreateMaintenanceHandler_When_ValidateReturnsError(t *testing.T) {

	body := fmt.Sprintf(`{"cost": 100, "started_at":"28-02-1996","description":""}`)

	req, err := http.NewRequest("POST", "/assets/ffb4b1a4-7bf5-11eb-9439-0242ac130002/maintenance", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	expectedErr := string(`{"error":"description is required"}`)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.CreateMaintenanceHandler(mockAssetMaintenanceService)).Methods("POST")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}

func TestMaintenanceActivitiesHandler_CreateMaintenanceHandler_When_ConvertReqFormatReturnsError(t *testing.T) {

	body := fmt.Sprintf(`{"cost": 100, "started_at":"28-02-199","description":"hardware issue"}`)

	req, err := http.NewRequest("POST", "/assets/ffb4b1a4-7bf5-11eb-9439-0242ac130002/maintenance", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	expectedErr := string(`{"error":"incorrect date format"}`)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.CreateMaintenanceHandler(mockAssetMaintenanceService)).Methods("POST")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}

func TestMaintenanceActivitiesHandler_DetailedMaintenanceActivityHandler_When_ServicesDetailedMaintenanceActivityReturnsError(t *testing.T) {

	req, err := http.NewRequest("GET", "/maintenance_activities/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	id := 1
	expectedErr := string(`{"error":"something went wrong"}`)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockAssetMaintenanceService.On("DetailedMaintenanceActivity", mock.Anything, id).Return(nil, errors.New("something went wrong"))
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.DetailedMaintenanceActivityHandler(mockAssetMaintenanceService)).Methods("GET")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}
func TestMaintenanceActivitiesHandler_DetailedMaintenanceActivityHandler_When_ServicesDetailedMaintenanceActivityReturnsCustomError(t *testing.T) {

	req, err := http.NewRequest("GET", "/maintenance_activities/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	id := 1
	expectedErr := string(`{"error":"id not found"}`)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockAssetMaintenanceService.On("DetailedMaintenanceActivity", mock.Anything, id).Return(nil, customerrors.MaintenanceIdDoesNotExist)
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.DetailedMaintenanceActivityHandler(mockAssetMaintenanceService)).Methods("GET")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}

func TestMaintenanceActivitiesHandler_DetailedMaintenanceActivityHandler_When_Success(t *testing.T) {

	req, err := http.NewRequest("GET", "/maintenance_activities/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	id := 1
	assetId, _ := uuid.Parse("ffb4b1a4-7bf5-11eb-9439-0242ac130002")
	startedDate := "28-02-1996"
	tym, err := time.Parse("02-01-2006", startedDate)

	outpt := domain.MaintenanceActivity{
		ID:          1,
		AssetId:     assetId,
		Cost:        100,
		StartedAt:   tym,
		EndedAt:     tym,
		Description: "hardware issue",
	}
	outptt := contract.DetailAssetMaintenanceActivityResponse{
		ID:          1,
		AssetId:     assetId,
		Cost:        100,
		StartedAt:   tym.Format("02-01-2006"),
		EndedAt:     tym.Format("02-01-2006"),
		Description: "hardware issue",
	}
	finalOutpt, _ := json.Marshal(outptt)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockAssetMaintenanceService.On("DetailedMaintenanceActivity", mock.Anything, id).Return(&outpt, nil)
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.DetailedMaintenanceActivityHandler(mockAssetMaintenanceService)).Methods("GET")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, string(finalOutpt), rr.Body.String())
}

func TestMaintenanceActivitiesHandler_DetailedMaintenanceActivityHandler_When_IdIsReturnsError(t *testing.T) {

	req, err := http.NewRequest("GET", "/maintenance_activities/wrongID", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	expectedErr := string(`{"error":"wrong id format"}`)

	mockAssetMaintenanceService := &mockService.MockMaintenanceActivityService{}
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.DetailedMaintenanceActivityHandler(mockAssetMaintenanceService)).Methods("GET")
	r.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}
