package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

const (
	getAssetDetails    = "SELECT id,category,status,purchase_at,purchase_cost,name,specifications FROM assets"
	createAssetQuery   = "INSERT INTO assets (id, status, category, purchase_at, purchase_cost, name, specifications) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *"
	getAssetByIDQuery  = "SELECT * FROM assets WHERE id=$1"
	getAssetName       = "SELECT status, specifications from assets WHERE id=$1 AND status!=$2"
	UpdateAssetDetails = "UPDATE assets SET status=$1, specifications=$2 WHERE id=$3"
	getAssetDelete     = "SELECT status FROM assets WHERE id=$1 AND status!=$2"
	getAssetDeletefun  = "UPDATE assets SET status=$1 WHERE id=$2"
	getAsset           = "SELECT id, category, status, purchase_at, purchase_cost, name, specifications FROM assets WHERE id=$1"
)

type AssetRepository interface {
	ListAssets(ctx context.Context) ([]domain.Asset, error)
	CreateAsset(ctx context.Context, asset_param *domain.Asset) (*domain.Asset, error)
	GetAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error)
	UpdateAsset(ctx context.Context, Id uuid.UUID, req contract.UpdateRequest) (*domain.Asset, error)
	DeleteAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error)
}

type assetRepo struct {
	db *sqlx.DB
}

func NewAssetRepository() AssetRepository {
	return &assetRepo{
		db: GetDB(),
	}
}

func (repo *assetRepo) DeleteAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error) {
	var m domain.Asset
	sample := "retired"
	err := repo.db.Get(&m, getAssetDelete, Id, sample)
	if err != nil {
		return nil, customerrors.NoAssetsExist
	}
	tx := repo.db.MustBegin()
	tx.MustExec(getAssetDeletefun, sample, Id)
	tx.Commit()

	return nil, nil

}
func (repo *assetRepo) UpdateAsset(ctx context.Context, Id uuid.UUID, req contract.UpdateRequest) (*domain.Asset, error) {
	var m domain.Asset
	sample := "retired"
	err := repo.db.Get(&m, getAssetName, Id, sample)

	if err != nil {
		return nil, customerrors.NoAssetsExist
	}

	if req.Status == nil {
		req.Status = &m.Status
	}

	if req.Specifications == nil {

		req.Specifications = m.Specifications

	}

	tx := repo.db.MustBegin()

	tx.MustExec(UpdateAssetDetails, *req.Status, req.Specifications, Id)
	tx.Commit()

	return nil, nil
}

func (repo *assetRepo) ListAssets(ctx context.Context) ([]domain.Asset, error) {
	var as []domain.Asset

	err := repo.db.Select(&as, getAssetDetails)

	if err == sql.ErrNoRows {
		fmt.Println("In repo.", as)
		return nil, customerrors.NoAssetsExist
	}
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (repo *assetRepo) CreateAsset(ctx context.Context, asset_param *domain.Asset) (*domain.Asset, error) {
	var asset domain.Asset
	err := repo.db.Get(&asset, createAssetQuery,
		asset_param.Id,
		asset_param.Status,
		asset_param.Category,
		asset_param.PurchaseAt,
		asset_param.PurchaseCost,
		asset_param.Name,
		asset_param.Specifications,
	)

	if err != nil {
		if err == customerrors.NoAssetsExist {
			fmt.Printf("Asset Repository: No asset exist : %s", err.Error())
			return nil, err
		}
		fmt.Printf("error in asset repository")
		return nil, err
	}

	return &asset, nil
}

func (repo *assetRepo) GetAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error) {
	var asset domain.Asset

	err := repo.db.Get(&asset, getAssetByIDQuery, Id)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("repository: couldn't find asset for Asset ID: %s", Id)
			return nil, customerrors.NoAssetsExist
		}
		return nil, err
	}

	return &asset, nil
}
