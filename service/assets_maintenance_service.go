package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type AssetMaintenanceService interface {
	CreateAssetMaintenance(ctx context.Context, req domain.MaintenanceActivity) (user *domain.MaintenanceActivity, err error)
	DetailedMaintenanceActivity(ctx context.Context, activityId int) (user *domain.MaintenanceActivity, err error)
	DeleteMaintenanceActivity(ctx context.Context, id int) (err error)
	GetAllForAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error)
	UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error)
}

type assetMaintenanceService struct {
	assetMaintainRepo repository.AssetMaintenanceRepo
	eventService      EventService
}

func NewAssetForMaintenance(repo repository.AssetMaintenanceRepo, eventService EventService) AssetMaintenanceService {
	return &assetMaintenanceService{
		assetMaintainRepo: repo,
		eventService:      eventService,
	}
}

func (service *assetMaintenanceService) CreateAssetMaintenance(ctx context.Context, req domain.MaintenanceActivity) (user *domain.MaintenanceActivity, err error) {
	assetsMaintenance, err := service.assetMaintainRepo.InsertMaintenanceActivity(ctx, req)

	if err != nil {
		fmt.Printf("servicelayer:%s", err.Error())
		return nil, err
	}

	return assetsMaintenance, nil
}

func (service *assetMaintenanceService) DetailedMaintenanceActivity(ctx context.Context, activityId int) (user *domain.MaintenanceActivity, err error) {
	assetsMaintenance, err := service.assetMaintainRepo.DetailedMaintenanceActivity(ctx, activityId)

	if err != nil {
		fmt.Printf("servicelayer:%s", err.Error())
		return nil, err
	}

	if assetsMaintenance == nil {
		return assetsMaintenance, customerrors.MaintenanceIdDoesNotExist
	}

	return assetsMaintenance, nil
}

func (service *assetMaintenanceService) DeleteMaintenanceActivity(ctx context.Context, activityId int) error {
	return service.assetMaintainRepo.DeleteMaintenanceActivity(ctx, activityId)
}

func (service *assetMaintenanceService) GetAllForAssetId(ctx context.Context, assetId uuid.UUID) ([]domain.MaintenanceActivity, error) {
	return service.assetMaintainRepo.GetAllByAssetId(ctx, assetId)
}

func (service *assetMaintenanceService) UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
	activity, err := service.assetMaintainRepo.UpdateMaintenanceActivity(ctx, req)
	if err != nil {
		return nil, err
	}
	id, err := service.eventService.PostEvent(ctx, *activity)

	if err != nil {
		fmt.Println("Failed to submit event: ", err)
	} else {
		fmt.Println("Successfully submitted event ", id)
	}
	return activity, nil
}
