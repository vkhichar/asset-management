package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type AssetAllocationService interface {
	CreateAssetAllocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error)
	AssetDeallocation(ctx context.Context, id uuid.UUID) (*string, error)
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
		return nil, err
	}
	return assetAllocation, nil
}
func (service *assetAllocationsService) AssetDeallocation(ctx context.Context, id uuid.UUID) (*string, error) {

	msg, err := service.assetAllocationsRepo.AssetDeallocation(ctx, id)
	if err == customerrors.ErrDeallocatedAlready {
		fmt.Printf("asset allocation service: already deallocated: %s", err.Error())
		return nil, customerrors.ErrDeallocatedAlready
	}
	if err != nil {
		fmt.Printf("asset allocation service: error while calling asset_deallocation: %s", err.Error())
		return nil, err
	}
	return msg, nil
}
