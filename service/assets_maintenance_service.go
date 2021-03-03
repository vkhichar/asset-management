package service

import (
	"context"

	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type AssetMaintenanceService interface {
	CreateAssetMaintenance(ctx context.Context, asset_id domain.UUID, req domain.MaintenanceActivity) (user *domain.MaintenanceActivity, err error)
}

type assetMaintenanceService struct {
	assetMaintainRepo repository.AssetMaintenanceRepo
}

func NewAssetForMaintenance(repo repository.AssetMaintenanceRepo) AssetMaintenanceService {
	return &assetMaintenanceService{
		assetMaintainRepo: repo,
	}
}
