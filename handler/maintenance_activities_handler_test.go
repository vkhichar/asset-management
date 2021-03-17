package handler_test

import (
	"bytes"
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
		EndedAt:     &tym,
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

func TestMaintenanceActivitiesHandler_DeleteById_When_Error(t *testing.T) {

	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockMaintenanceService.On("DeleteMaintenanceActivity", mock.Anything, 1).Return(errors.New("Failed to delete activity"))

	req, err := http.NewRequest("DELETE", "/maintenance_activities/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.DeleteMaintenanceActivityHandler(mockMaintenanceService)).Methods("DELETE")
	r.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusInternalServerError, resRec.Result().StatusCode)
	assert.JSONEq(t, string(`{ "error" : "Something went wrong"}`), resRec.Body.String())

}

func TestMaintenanceActivitiesHandler_DeleteById_When_InvalidId(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/maintenance_activities/abcd", nil)

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}",
		handler.DeleteMaintenanceActivityHandler(&mockService.MockMaintenanceActivityService{})).Methods("DELETE")
	r.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusBadRequest, resRec.Result().StatusCode)

}

func TestMaintenanceActivitiesHandler_DeleteById_When_Success(t *testing.T) {
	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockMaintenanceService.On("DeleteMaintenanceActivity", mock.Anything, 1).Return(nil)

	req, err := http.NewRequest("DELETE", "/maintenance_activities/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.DeleteMaintenanceActivityHandler(mockMaintenanceService)).Methods("DELETE")
	r.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusNoContent, resRec.Result().StatusCode)

}

func TestMaintenanceActivitiesHandler_ListAllByAssetId_When_Error(t *testing.T) {
	assetId := uuid.New()

	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockMaintenanceService.On("GetAllForAssetId", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to fetch activities"))

	req, err := http.NewRequest("GET", fmt.Sprintf("/assets/%s/maintenance", assetId.String()), nil)

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.ListMaintenanceActivitiesByAsserId(mockMaintenanceService)).Methods("GET")
	r.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusInternalServerError, resRec.Result().StatusCode)
	assert.JSONEq(t, string(`{ "error" : "Something went wrong"}`), resRec.Body.String())
}

func TestMaintenanceActivitiesHandler_ListAllByAssetId_When_NonEmptyResult(t *testing.T) {
	assetId := uuid.New()

	activities := make([]domain.MaintenanceActivity, 1)
	date := time.Now()
	activities[0] = domain.MaintenanceActivity{
		ID:          1,
		AssetId:     assetId,
		Cost:        20,
		StartedAt:   date,
		EndedAt:     &date,
		Description: "test",
	}

	output := make([]contract.MaintenanceActivityResp, 1)
	output[0] = contract.NewMaintenanceActivityResp(activities[0])

	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockMaintenanceService.On("GetAllForAssetId", mock.Anything, mock.Anything).Return(activities, nil)

	req, err := http.NewRequest("GET", fmt.Sprintf("/assets/%s/maintenance", assetId.String()), nil)

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.ListMaintenanceActivitiesByAsserId(mockMaintenanceService)).Methods("GET")
	r.ServeHTTP(resRec, req)

	res, err := json.Marshal(output)
	if err != nil {
		t.Fatal()
	}

	assert.Equal(t, http.StatusOK, resRec.Result().StatusCode)
	assert.JSONEq(t, string(res), resRec.Body.String())
}

func TestMaintenanceActivitiesHandler_ListAllByAssetId_When_EmptyResult(t *testing.T) {
	assetId := uuid.New()
	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockMaintenanceService.On("GetAllForAssetId", mock.Anything, mock.Anything).Return([]domain.MaintenanceActivity{}, nil)

	req, err := http.NewRequest("GET", fmt.Sprintf("/assets/%s/maintenance", assetId.String()), nil)

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/assets/{asset_id}/maintenance", handler.ListMaintenanceActivitiesByAsserId(mockMaintenanceService)).Methods("GET")
	r.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusOK, resRec.Result().StatusCode)
	assert.JSONEq(t, string(`[]`), resRec.Body.String())
}

func TestMaintenanceActivitiesHandler_UpdateById_When_DbError(t *testing.T) {

	reqBody := contract.UpdateMaintenanceActivityReq{
		Cost:        20,
		Description: "descr",
		EndedAt:     "2020-03-03",
	}

	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockMaintenanceService.On("UpdateMaintenanceActivity", mock.Anything, mock.Anything).Return(nil, errors.New("Failed to update activity"))

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal()
	}

	req, err := http.NewRequest("PUT", "/maintenance_activities/1", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.UpdateMaintenanceActivity(mockMaintenanceService)).Methods("PUT")
	r.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusInternalServerError, resRec.Result().StatusCode)
	assert.JSONEq(t, string(`{ "error" : "Something went wrong"}`), resRec.Body.String())
}

func TestMaintenanceActivitiesHandler_UpdateById_When_ErrNotFound(t *testing.T) {

	reqBody := contract.UpdateMaintenanceActivityReq{
		Cost:        20,
		Description: "description",
		EndedAt:     "2020-03-03",
	}

	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockMaintenanceService.On("UpdateMaintenanceActivity", mock.Anything, mock.Anything).Return(nil, customerrors.ErrNotFound)

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal()
	}

	req, err := http.NewRequest("PUT", "/maintenance_activities/1", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.UpdateMaintenanceActivity(mockMaintenanceService)).Methods("PUT")
	r.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusNotFound, resRec.Result().StatusCode)
	assert.JSONEq(t, string(`{ "error" : "not found"}`), resRec.Body.String())
}

func TestMaintenanceActivitiesHandler_UpdateById_When_InvalidRequest(t *testing.T) {

	reqBody := contract.UpdateMaintenanceActivityReq{
		Cost:        20,
		Description: "description",
		EndedAt:     "dsds",
	}

	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal()
	}

	req, err := http.NewRequest("PUT", "/maintenance_activities/1", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.UpdateMaintenanceActivity(mockMaintenanceService)).Methods("PUT")
	r.ServeHTTP(resRec, req)

	assert.Equal(t, http.StatusBadRequest, resRec.Result().StatusCode)
	assert.JSONEq(t, string(`{ "error" : "Invalid date ended_at"}`), resRec.Body.String())
}

func TestMaintenanceActivitiesHandler_UpdateById_When_Success(t *testing.T) {

	reqBody := contract.UpdateMaintenanceActivityReq{
		Cost:        25,
		Description: "description",
		EndedAt:     "2020-01-28",
	}

	date, err := time.Parse(contract.DateFormat, reqBody.EndedAt)
	assetId := uuid.New()
	updated := domain.MaintenanceActivity{
		ID:          1,
		AssetId:     assetId,
		Cost:        25,
		StartedAt:   date,
		EndedAt:     &date,
		Description: "description",
	}
	mockMaintenanceService := &mockService.MockMaintenanceActivityService{}
	mockMaintenanceService.On("UpdateMaintenanceActivity", mock.Anything, mock.Anything).Return(&updated, nil)

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal()
	}

	req, err := http.NewRequest("PUT", "/maintenance_activities/1", bytes.NewBuffer(body))

	if err != nil {
		t.Fatal(err)
	}

	resRec := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/maintenance_activities/{id}", handler.UpdateMaintenanceActivity(mockMaintenanceService)).Methods("PUT")
	r.ServeHTTP(resRec, req)

	expected := contract.MaintenanceActivityResp{
		Id:          1,
		AssetId:     assetId,
		Cost:        25,
		StartedAt:   "2020-01-28",
		EndedAt:     "2020-01-28",
		Description: "description",
	}

	assert.Equal(t, http.StatusOK, resRec.Result().StatusCode)
	expectedRes, err := json.Marshal(expected)
	assert.JSONEq(t, string(expectedRes), resRec.Body.String())
}
