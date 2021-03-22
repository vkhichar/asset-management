package repository

import (
	"context"
	"database/sql"
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

const (
	checkIfAssetIsAllocated = "SELECT * FROM asset_allocations WHERE asset_id=$1 ORDER BY id DESC LIMIT 1"
	createAssetAllocation   = "INSERT INTO asset_allocations(user_id, asset_id, allocated_by, allocated_from) VALUES($1, $2, $3, $4) RETURNING *"
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
	user, err := repo.userRepo.GetUserByID(ctx, req.UserId)
	if user == nil {
		return nil, customerrors.UserNotExist
	}

	asset, err := repo.assetRepo.GetAsset(ctx, req.AssetId)
	if asset == nil {
		return nil, customerrors.AssetDoesNotExist
	}

	if asset.Status != "active" {
		return nil, customerrors.AssetCannotBeAllocated
	}

	var assetAllocated domain.AssetAllocations
	err = repo.db.Get(&assetAllocated, checkIfAssetIsAllocated, req.AssetId)
	if err != sql.ErrNoRows && err != nil {
		return nil, err
	}

	if assetAllocated.AllocatedTill == nil {
		return nil, customerrors.AssetAlreadyAllocated
	}

	admin, err := repo.userRepo.GetUserByID(ctx, req.AllocatedBy)
	if admin == nil {
		return nil, customerrors.AdminDoesNotExist
	}
	err = repo.db.Get(&assetAllocated, createAssetAllocation, req.UserId, req.AssetId, admin.Name, time.Now())
	repo.db.MustBegin().Commit()
	if err != nil {
		return nil, err
	}

	return &assetAllocated, nil
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
