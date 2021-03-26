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
	eventSvc          EventService
	assetRepo         repository.AssetRepository
}

func NewAssetForMaintenance(repo repository.AssetMaintenanceRepo, es EventService, assetRepo repository.AssetRepository) AssetMaintenanceService {
	return &assetMaintenanceService{
		assetMaintainRepo: repo,
		eventSvc:          es,
		assetRepo:         assetRepo,
	}
}

func (service *assetMaintenanceService) CreateAssetMaintenance(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
	assetsMaintenance, err := service.assetMaintainRepo.InsertMaintenanceActivity(ctx, req)

	if err != nil {
		fmt.Printf("servicelayer:%s", err.Error())
		return nil, err
	}

	eventID, errEvent := service.eventSvc.PostAssetMaintenanceActivityEvent(ctx, assetsMaintenance)
	if errEvent == customerrors.ResponseTimeLimitExceeded {
		fmt.Printf("servicelayere events:%s", errEvent.Error())
		return nil, errEvent
	}

	if errEvent != nil {
		fmt.Println("Event cannot be created")
		fmt.Printf("servicelayer:%s", errEvent.Error())

		return assetsMaintenance, errEvent

	}
	fmt.Println("Event created with id:", eventID)
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
	_, err := service.assetRepo.GetAsset(ctx, assetId)
	if err != nil {
		return nil, err
	}
	return service.assetMaintainRepo.GetAllByAssetId(ctx, assetId)
}

func (service *assetMaintenanceService) UpdateMaintenanceActivity(ctx context.Context, req domain.MaintenanceActivity) (*domain.MaintenanceActivity, error) {
	activity, err := service.assetMaintainRepo.UpdateMaintenanceActivity(ctx, req)
	if err != nil {
		return nil, err
	}
	id, err := service.eventSvc.PostMaintenanceActivity(ctx, *activity)

	if err != nil {
		fmt.Println("Failed to submit event: ", err)
	} else {
		fmt.Println("Successfully submitted event ", id)
	}
	return activity, nil
}
