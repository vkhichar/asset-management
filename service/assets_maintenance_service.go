package service

import (
	"context"

	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type AssetMaintenanceService interface {
	CreateAssetMaintenance(ctx context.Context, assets_id string) (user *domain.Maintenance, err error)
}

type assetMaintenanceService struct {
	assetMaintainRepo repository.AssetMaintenanceRepo
}

func NewAssetMaintenance(repo repository.AssetMaintenanceRepo) AssetMaintenanceService {
	return &assetMaintenanceService{
		assetMaintainRepo: repo,
	}
}
