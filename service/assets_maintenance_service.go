package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type AssetMaintenanceService interface {
	CreateAssetMaintenance(ctx context.Context, assetId uuid.UUID, req domain.MaintenanceActivity) (user *domain.MaintenanceActivity, err error)
}

type assetMaintenanceService struct {
	assetMaintainRepo repository.AssetMaintenanceRepo
}

func NewAssetForMaintenance(repo repository.AssetMaintenanceRepo) AssetMaintenanceService {
	return &assetMaintenanceService{
		assetMaintainRepo: repo,
	}
}

func (service *assetMaintenanceService) CreateAssetMaintenance(ctx context.Context, assetId uuid.UUID, req domain.MaintenanceActivity) (user *domain.MaintenanceActivity, err error) {
	maintain, err := service.assetMaintainRepo.InsertMaintenanceActivity(ctx, assetId, req)
	if err != nil {
		fmt.Printf("servicelayer:%s", err.Error())
		return nil, err
	}

	return maintain, nil
}
