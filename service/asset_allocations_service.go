package service

import (
	"context"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type AssetAllocationService interface {
	CreateAssetAllocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error)
	AssetDeallocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error)
}

type assetAllocationsService struct {
	assetAllocationsRepo repository.AssetAllocationsRepository
}

func NewAssetAllocationService(repo repository.AssetAllocationsRepository) AssetAllocationService {
	return &assetAllocationsService{
		assetAllocationsRepo: repo,
	}
}

func (service *assetAllocationsService) CreateAssetAllocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error) {
	return nil, nil
}
func (service *assetAllocationsService) AssetDeallocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error) {
	return nil, nil
}
