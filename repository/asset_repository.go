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
	getAssetDetails    = "SELECT id, category, status, purchase_at, purchase_cost, name, specifications FROM assets"
	getAssetName       = "SELECT status, specifications from assets WHERE id=$1 AND status!=$2"
	UpdateAssetDetails = "UPDATE assets SET status=$1, specifications=$2 WHERE id=$3"
	getAssetDelete     = "SELECT status FROM assets WHERE id=$1 AND status!=$2"
	getAssetDeletefun  = "UPDATE assets SET status=$1 WHERE id=$2"
	getAsset           = "SELECT id, category, status, purchase_at, purchase_cost, name, specifications FROM assets WHERE id=$1"
)

type AssetRepository interface {
	ListAssets(ctx context.Context) ([]domain.Asset, error)
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
	var asset domain.Asset
	sample := "retired"
	err := repo.db.Get(&m, getAssetDelete, Id, sample)
	if err != nil {
		return nil, customerrors.NoAssetsExist
	}
	tx := repo.db.MustBegin()
	tx.MustExec(getAssetDeletefun, sample, Id)
	tx.Commit()
	err = repo.db.Get(&asset, getAsset, Id)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &asset, nil

}
func (repo *assetRepo) UpdateAsset(ctx context.Context, Id uuid.UUID, req contract.UpdateRequest) (*domain.Asset, error) {
	var m domain.Asset
	var asset domain.Asset
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
	err = repo.db.Get(&asset, getAsset, Id)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}

func (repo *assetRepo) ListAssets(ctx context.Context) ([]domain.Asset, error) {
	var as []domain.Asset
	err := repo.db.Select(&as, getAssetDetails)
	if err == sql.ErrNoRows {

		return nil, customerrors.NoAssetsExist
	}
	if err != nil {
		return nil, err
	}
	return as, nil

}
