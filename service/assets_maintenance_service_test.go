package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/domain"
	mockRepo "github.com/vkhichar/asset-management/repository/mocks"
	"github.com/vkhichar/asset-management/service"
	mockEventService "github.com/vkhichar/asset-management/service/mocks"
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
	mockEventSvc := &mockEventService.MockEventService{}

	mockAssetMaintenanceRepo.On("InsertMaintenanceActivity", ctx, req).Return(nil, errors.New("servicelayer:"))

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)

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
	mockAssetMaintenanceRepo.On("PostAssetMaintenanceActivityEvent", ctx, res).Return(nil, errors.New("error in event service"))

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

	maintenanceActivities, err := maintenanceActivity.CreateAssetMaintenance(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, &res, maintenanceActivities)
}
func TestAssetsMaintenanceService_DetailedMaintenanceActivity_When_DetailedMaintenanceActivityReturnsError(t *testing.T) {
	ctx := context.Background()
	activityId := 123
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	mockEventSvc := &mockEventService.MockEventService{}

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(nil, errors.New("repolayer:Failed to fetch asset maintenance activities"))

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)

	maintenanceActivities, err := maintenanceActivity.DetailedMaintenanceActivity(ctx, activityId)

	assert.Error(t, err)
	assert.Equal(t, "repolayer:Failed to fetch asset maintenance activities", err.Error())
	assert.Nil(t, maintenanceActivities)
}

func TestAssetsMaintenanceService_DetailedMaintenanceActivity_When_Success(t *testing.T) {
	ctx := context.Background()
	activityId := 123
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

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(&res, nil)

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)

	maintenanceActivities, err := maintenanceActivity.DetailedMaintenanceActivity(ctx, activityId)

	assert.NoError(t, err)
	assert.Equal(t, &res, maintenanceActivities)
}

func TestAssetsMaintenanceService_DetailedMaintenanceActivity_When_DetailedMaintenanceActivityReturnsNilNil(t *testing.T) {
	ctx := context.Background()
	activityId := 123
	mockAssetMaintenanceRepo := &mockRepo.MockMaintenanceActivityRepo{}
	mockEventSvc := &mockEventService.MockEventService{}

	mockAssetMaintenanceRepo.On("DetailedMaintenanceActivity", ctx, activityId).Return(nil, nil)

	maintenanceActivity := service.NewAssetForMaintenance(mockAssetMaintenanceRepo, mockEventSvc)

	maintenanceActivities, err := maintenanceActivity.DetailedMaintenanceActivity(ctx, activityId)

	assert.Error(t, err)
	assert.Nil(t, maintenanceActivities)
}
