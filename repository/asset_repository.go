package repository

import (
	"github.com/jmoiron/sqlx"
)

type AssetRepository interface {
}

type assetRepo struct {
	db *sqlx.DB
}

func NewAssetRepository() AssetRepository {
	return &assetRepo{
		db: GetDB(),
	}
}
