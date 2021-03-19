package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
)

type AssetAllocationsRepository interface {
	CreateAssetAllocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error)
	AssetDeallocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error)
}

type assetAllocationsRepo struct {
	db *sqlx.DB
}

func NewAssetAllocationRepository() AssetAllocationsRepository {
	return &assetAllocationsRepo{
		db: GetDB(),
	}
}

func (repo *assetAllocationsRepo) CreateAssetAllocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error) {
	return nil, nil
}
func (repo *assetAllocationsRepo) AssetDeallocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error) {
	return nil, nil
}
