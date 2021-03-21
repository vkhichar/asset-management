package service

import (
	"context"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
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
	assetAllocation, err := service.assetAllocationsRepo.CreateAssetAllocation(ctx, req)

	if err != nil {
		if err == customerrors.UserNotExist {
			return nil, customerrors.UserNotExist
		}
		if err == customerrors.AssetDoesNotExist {
			return nil, customerrors.AssetDoesNotExist
		}
		if err == customerrors.AssetCannotBeAllocated {
			return nil, customerrors.AssetCannotBeAllocated
		}
		if err == customerrors.AssetAlreadyAllocated {
			return nil, customerrors.AssetAlreadyAllocated
		}
		return nil, err
	}
	return assetAllocation, nil
}
func (service *assetAllocationsService) AssetDeallocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error) {
	return nil, nil
}
