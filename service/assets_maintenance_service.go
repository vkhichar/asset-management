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
}

func NewAssetForMaintenance(repo repository.AssetMaintenanceRepo) AssetMaintenanceService {
	return &assetMaintenanceService{
		assetMaintainRepo: repo,
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
	return service.assetMaintainRepo.UpdateMaintenanceActivity(ctx, req)
}
