package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
	//repository "github.com/vkhichar/asset-management/repository"
)

const (
	getAssetIDQuery       = "SELECT MAX(id) FROM asset_allocations  GROUP BY asset_id HAVING asset_id=$1 ORDER BY COUNT(id) DESC"
	assetDeallocateQuery  = "UPDATE asset_allocations SET allocated_till=$1 WHERE id=$2"
	getAssetAllocatedTime = "SELECT allocated_till from asset_allocations where id=$1"
	assetAllocatedQuery   = "INSERT INTO asset_allocations (user_id, asset_id, allocated_by VALUES ($1,$2,$3)"
)

type AssetAllocationsRepository interface {
	CreateAssetAllocation(ctx context.Context, req contract.CreateAssetAllocationRequest) (*domain.AssetAllocations, error)
	AssetDeallocation(ctx context.Context, id uuid.UUID) (*string, error)
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

	return nil, nil
}
func (repo *assetAllocationsRepo) AssetDeallocation(ctx context.Context, asset_id uuid.UUID) (*string, error) {
	var ID int
	var date *string
	err := repo.db.Get(&ID, getAssetIDQuery, asset_id)
	if err != nil {
		fmt.Printf("asset allocation repo: error while getting id from asset allocation table: %s", err.Error())
		return nil, err
	}
	err = repo.db.Get(&date, getAssetAllocatedTime, ID)

	if date != nil {
		fmt.Printf("asset deallocated already:")
		return nil, customerrors.ErrDeallocatedAlready
	}
	if err != nil {
		fmt.Println("error while getting date")
		return nil, err
	}
	err = repo.db.Get(&ID, getAssetIDQuery, asset_id)

	tx := repo.db.MustBegin()
	tx.MustExec(assetDeallocateQuery, time.Now(), ID)
	tx.Commit()
	var msg = "asset dellocated successfully"

	return &msg, nil
}
