package service

import (
	"github.com/vkhichar/asset-management/repository"
)

type AssetService interface {
}

type assetService struct {
	assetRepo repository.AssetRepository
}

func NewAssetService(repo repository.AssetRepository) AssetService {
	return &assetService{
		assetRepo: repo,
	}
}
