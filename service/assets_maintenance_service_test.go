package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
	mockRepo "github.com/vkhichar/asset-management/repository/mocks"
	"github.com/vkhichar/asset-management/service"
<<<<<<< HEAD
	mockEventService "github.com/vkhichar/asset-management/service/mocks"
=======
	mockService "github.com/vkhichar/asset-management/service/mocks"
>>>>>>> a57ceeea7a603f523eb02e7c113394f9f64b67ee
)

func TestAssetsMaintenanceService_CreateAssetMaintenance_When_InsertMaintenanceActivityReturnsError(t *testing.T) {
	ctx := context.Background()
	req := domain.MaintenanceActivity{
		AssetId:     uuid.New(),
		Cost:        100,
		StartedAt:   time.Now(),
		Description: "hardware corrupted",
	}
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
<<<<<<< HEAD
	mockEventSvc := &mockEventService.MockEventService{}

	mockAssetMaintenanceRepo.On("InsertMaintenanceActivity", ctx, req).Return(nil, errors.New("servicelayer:"))

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)
=======
	mockEventService := &mockService.MockEventService{}

	mockAssetMaintenanceRepo.On("InsertMaintenanceActivity", ctx, req).Return(nil, errors.New("servicelayer:"))

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
>>>>>>> a57ceeea7a603f523eb02e7c113394f9f64b67ee

	maintenanceActivities, err := maintenanceActivity.CreateAssetMaintenance(ctx, req)

	assert.Error(t, err)
	assert.Equal(t, "servicelayer:", err.Error())
	assert.Nil(t, maintenanceActivities)
}

func TestAssetsMaintenanceService_CreateAssetMaintenance_When_EventServiceReturnsError(t *testing.T) {
	ctx := context.Background()
	req := domain.MaintenanceActivity{
		AssetId:     uuid.New(),
		Cost:        100,
		StartedAt:   time.Now(),
		Description: "hardware corrupted",
	}
	date := time.Now()

	res := domain.MaintenanceActivity{
		ID:          1,
		AssetId:     uuid.New(),
		Cost:        100,
		StartedAt:   date,
		EndedAt:     &date,
		Description: "hardware corrupted",
	}
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
<<<<<<< HEAD
	mockEventSvc := &mockEventService.MockEventService{}
=======
	mockEventService := &mockService.MockEventService{}
>>>>>>> a57ceeea7a603f523eb02e7c113394f9f64b67ee

	mockAssetMaintenanceRepo.On("InsertMaintenanceActivity", ctx, req).Return(&res, nil)
	mockAssetMaintenanceRepo.On("PostAssetMaintenanceActivityEvent", ctx, res).Return(nil, errors.New("error in event service"))

<<<<<<< HEAD
	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)

	maintenanceActivities, err := maintenanceActivity.CreateAssetMaintenance(ctx, req)

	assert.Nil(t, maintenanceActivities)
	assert.Equal(t, "error in event service", err.Error())
}
func TestAssetsMaintenanceService_CreateAssetMaintenance_When_Success(t *testing.T) {
	ctx := context.Background()
	req := domain.MaintenanceActivity{
		AssetId:     uuid.New(),
		Cost:        100,
		StartedAt:   time.Now(),
		Description: "hardware corrupted",
	}

	res := domain.MaintenanceActivity{
		ID:          1,
		AssetId:     uuid.New(),
		Cost:        100,
		StartedAt:   time.Now(),
		EndedAt:     time.Now(),
		Description: "hardware corrupted",
	}
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	mockEventSvc := &mockEventService.MockEventService{}
	mockAssetMaintenanceRepo.On("InsertMaintenanceActivity", ctx, req).Return(&res, nil)
	mockEventSvc.On("PostAssetMaintenanceActivityEvent", ctx, &res).Return("14", nil)

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)
=======
	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
>>>>>>> a57ceeea7a603f523eb02e7c113394f9f64b67ee

	maintenanceActivities, err := maintenanceActivity.CreateAssetMaintenance(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, &res, maintenanceActivities)
}
func TestAssetsMaintenanceService_DetailedMaintenanceActivity_When_DetailedMaintenanceActivityReturnsError(t *testing.T) {
	ctx := context.Background()
	activityId := 123
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
<<<<<<< HEAD
	mockEventSvc := &mockEventService.MockEventService{}

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(nil, errors.New("repolayer:Failed to fetch asset maintenance activities"))

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)
=======
	mockEventService := &mockService.MockEventService{}

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(nil, errors.New("repolayer:Failed to fetch asset maintenance activities"))

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
>>>>>>> a57ceeea7a603f523eb02e7c113394f9f64b67ee

	maintenanceActivities, err := maintenanceActivity.DetailedMaintenanceActivity(ctx, activityId)

	assert.Error(t, err)
	assert.Equal(t, "repolayer:Failed to fetch asset maintenance activities", err.Error())
	assert.Nil(t, maintenanceActivities)
}

func TestAssetsMaintenanceService_DetailedMaintenanceActivity_When_Success(t *testing.T) {
	ctx := context.Background()
	activityId := 123
	date := time.Now()
	res := domain.MaintenanceActivity{
		ID:          1,
		AssetId:     uuid.New(),
		Cost:        100,
		StartedAt:   date,
		EndedAt:     &date,
		Description: "hardware corrupted",
	}
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
<<<<<<< HEAD
	mockEventSvc := &mockEventService.MockEventService{}

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(&res, nil)

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)
=======
	mockEventService := &mockService.MockEventService{}

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(&res, nil)

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
>>>>>>> a57ceeea7a603f523eb02e7c113394f9f64b67ee

	maintenanceActivities, err := maintenanceActivity.DetailedMaintenanceActivity(ctx, activityId)

	assert.NoError(t, err)
	assert.Equal(t, &res, maintenanceActivities)
}

func TestAssetsMaintenanceService_DetailedMaintenanceActivity_When_DetailedMaintenanceActivityReturnsNilNil(t *testing.T) {
	ctx := context.Background()
	activityId := 123
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
<<<<<<< HEAD
	mockEventSvc := &mockEventService.MockEventService{}

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(nil, nil)

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)
=======
	mockEventService := &mockService.MockEventService{}

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(nil, nil)

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
>>>>>>> a57ceeea7a603f523eb02e7c113394f9f64b67ee

	maintenanceActivities, err := maintenanceActivity.DetailedMaintenanceActivity(ctx, activityId)

	assert.Error(t, err)
	assert.Nil(t, maintenanceActivities)
}

func TestMaintenanceActivitiesHandler_DeleteById_When_DeleteReturnsError(t *testing.T) {
	ctx := context.Background()
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	mockAssetMaintenanceRepo.On("DeleteMaintenanceActivity", ctx, 1).Return(errors.New("Failed to delete activity"))

	mockEventService := &mockService.MockEventService{}

	service := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)

	err := service.DeleteMaintenanceActivity(ctx, 1)

	assert.Error(t, err)
	assert.Equal(t, "Failed to delete activity", err.Error())
}

func TestMaintenanceActivitiesHandler_DeleteById_When_Success(t *testing.T) {
	ctx := context.Background()
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	mockAssetMaintenanceRepo.On("DeleteMaintenanceActivity", ctx, 1).Return(nil)
	mockEventService := &mockService.MockEventService{}

	service := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)

	err := service.DeleteMaintenanceActivity(ctx, 1)

	assert.NoError(t, err)
}

func TestMaintenanceActivitiesHandler_GetAllByAssetId_When_GetAllReturnsError(t *testing.T) {
	ctx := context.Background()
	assetId, _ := uuid.NewUUID()
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	mockAssetMaintenanceRepo.On("GetAllByAssetId", ctx, assetId).Return(nil, errors.New("Failed to fetch activities"))
	mockEventService := &mockService.MockEventService{}

	service := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)

	activities, err := service.GetAllForAssetId(ctx, assetId)

	assert.Error(t, err)
	assert.Equal(t, "Failed to fetch activities", err.Error())
	assert.Nil(t, activities)

}

func TestMaintenanceActivitiesHandler_GetAllByAssetId_When_GetAllReturnsNoData(t *testing.T) {
	ctx := context.Background()
	assetId, _ := uuid.NewUUID()
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	mockAssetMaintenanceRepo.On("GetAllByAssetId", ctx, assetId).Return([]domain.MaintenanceActivity{}, nil)
	mockEventService := &mockService.MockEventService{}

	service := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
	activities, err := service.GetAllForAssetId(ctx, assetId)

	assert.NoError(t, err)
	assert.Len(t, activities, 0)

}

func TestMaintenanceActivitiesHandler_GetAllByAssetId_When_GetAllReturnsData(t *testing.T) {
	ctx := context.Background()
	assetId, _ := uuid.NewUUID()
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}

	activities := make([]domain.MaintenanceActivity, 1)
	date := time.Now()
	activities[0] = domain.MaintenanceActivity{
		ID:        1,
		AssetId:   assetId,
		Cost:      20,
		StartedAt: time.Now(),
		EndedAt:   &date,
	}

	mockAssetMaintenanceRepo.On("GetAllByAssetId", ctx, assetId).Return(activities, nil)

	mockEventService := &mockService.MockEventService{}

	maintenanceService := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
	activities, err := maintenanceService.GetAllForAssetId(ctx, assetId)

	assert.Nil(t, err)
	assert.NotNil(t, activities)
	assert.Len(t, activities, 1)
}

func TestMaintenanceActivitiesHandler_UpdateMaintenanceActivity_When_UpdateReturnsError(t *testing.T) {

	ctx := context.Background()
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	assetId, _ := uuid.NewUUID()

	date := time.Now()
	activity := domain.MaintenanceActivity{
		ID:        1,
		AssetId:   assetId,
		Cost:      20,
		StartedAt: date,
		EndedAt:   &date,
	}
	mockAssetMaintenanceRepo.On("UpdateMaintenanceActivity", ctx, mock.Anything).Return(nil, errors.New("Failed to fetch activity"))

	mockEventService := &mockService.MockEventService{}
	mockEventService.On("PostMaintenanceActivity", ctx, mock.Anything).Return("", nil)

	maintenanceService := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
	_, err := maintenanceService.UpdateMaintenanceActivity(ctx, activity)

	assert.NotNil(t, err)
}

func TestMaintenanceActivitiesHandler_UpdateMaintenanceActivity_When_UpdateIsSuccessful(t *testing.T) {

	ctx := context.Background()
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	assetId, _ := uuid.NewUUID()
	date := time.Now()
	activity := domain.MaintenanceActivity{
		ID:        1,
		AssetId:   assetId,
		Cost:      20,
		StartedAt: date,
		EndedAt:   &date,
	}
	mockAssetMaintenanceRepo.On("UpdateMaintenanceActivity", ctx, mock.Anything).Return(&activity, nil)

	mockEventService := &mockService.MockEventService{}
	mockEventService.On("PostMaintenanceActivity", ctx, mock.Anything).Return("", nil)

	maintenanceService := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventService)
	output, err := maintenanceService.UpdateMaintenanceActivity(ctx, activity)

	assert.Nil(t, err)
	assert.NotNil(t, output)

}
