package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

const (
	checkIfAssetIsAllocated = "SELECT * FROM asset_allocations WHERE asset_id=$1 ORDER BY id DESC LIMIT 1"
	createAssetAllocation   = "INSERT INTO asset_allocations(user_id, asset_id, allocated_by, allocated_from) VALUES($1, $2, $3, $4) RETURNING *"
)

type AssetAllocationsRepository interface {
	CreateAssetAllocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error)
	AssetDeallocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error)
}

type assetAllocationsRepo struct {
	db        *sqlx.DB
	userRepo  UserRepository
	assetRepo AssetRepository
}

func NewAssetAllocationRepository(uRepo UserRepository, aRepo AssetRepository) AssetAllocationsRepository {
	return &assetAllocationsRepo{
		db:        GetDB(),
		userRepo:  uRepo,
		assetRepo: aRepo,
	}
}

func (repo *assetAllocationsRepo) CreateAssetAllocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error) {
	var assetAllocated1 *domain.AssetAllocations
	repo.db.Get(&assetAllocated1, "Select * from asset_allocations where user_id=$1", 71)
	fmt.Println(assetAllocated1)

	user, _ := repo.userRepo.GetUserByID(ctx, req.UserId)
	if user == nil {
		return nil, customerrors.UserNotExist
	}

	asset, _ := repo.assetRepo.GetAsset(ctx, req.AssetId)
	if asset == nil {
		return nil, customerrors.AssetDoesNotExist
	}

	if asset.Status != "active" {
		return nil, customerrors.AssetCannotBeAllocated
	}

	var assetAllocated domain.AssetAllocations
	err := repo.db.Get(&assetAllocated, checkIfAssetIsAllocated, req.AssetId)
	if err != sql.ErrNoRows && err != nil {
		fmt.Printf("repo layer line 58: %s", err.Error())
		return nil, err
	}

	if assetAllocated.AllocatedTill == nil {
		return nil, customerrors.AssetAlreadyAllocated
	}

	admin, _ := repo.userRepo.GetUserByID(ctx, req.AllocatedBy)
	err = repo.db.Get(&assetAllocated, createAssetAllocation, req.UserId, req.AssetId, admin.Name, time.Now())
	//fmt.Println(assetAllocated)
	if err != nil {
		fmt.Printf("repo layer line 70: %s", err.Error())
		return nil, err
	}
	return &assetAllocated, nil
}
func (repo *assetAllocationsRepo) AssetDeallocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error) {
	return nil, nil
}
