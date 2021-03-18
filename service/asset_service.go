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

type AssetService interface {
	ListAssets(ctx context.Context) ([]domain.Asset, error)
	CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error)
	GetAsset(ctx context.Context, ID uuid.UUID) (*domain.Asset, error)
	UpdateAsset(ctx context.Context, Id uuid.UUID, req contract.UpdateRequest) (*domain.Asset, error)
	DeleteAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
	eventSvc  EventService
}

func NewAssetService(repo repository.AssetRepository, event EventService) AssetService {
	return &assetService{
		assetRepo: repo,
		eventSvc:  event,
	}
}
func (service *assetService) DeleteAsset(ctx context.Context, Id uuid.UUID) (*domain.Asset, error) {
	asset, err := service.assetRepo.DeleteAsset(ctx, Id)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (service *assetService) UpdateAsset(ctx context.Context, Id uuid.UUID, req contract.UpdateRequest) (*domain.Asset, error) {
	asset, err := service.assetRepo.UpdateAsset(ctx, Id, req)
	if err != nil {
		return nil, err
	}
	return asset, nil
}

func (service *assetService) ListAssets(ctx context.Context) ([]domain.Asset, error) {
	asset, err := service.assetRepo.ListAssets(ctx)
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, customerrors.NoAssetsExist
	}
	return asset, nil
}

func (service *assetService) CreateAsset(ctx context.Context, assetParam *domain.Asset) (*domain.Asset, error) {
	asset, err := service.assetRepo.CreateAsset(ctx, assetParam)
	if err != nil {
		if err == customerrors.NoAssetsExist {
			fmt.Printf("Asset service: asset does not exist: %s", err.Error())
			return nil, err
		}
		fmt.Printf("asset_service error while creating asset: %s", err.Error())
		return nil, err
	}

	id, err := service.eventSvc.PostAssetEventCreateAsset(ctx, asset)
	if err != nil {
		fmt.Printf("asset service: error during post create asset event: %s", err.Error())
		return asset, err
	} else {
		fmt.Println("New event created successfully:", id)
	}

	return asset, err
}

func (service *assetService) GetAsset(ctx context.Context, ID uuid.UUID) (*domain.Asset, error) {
	asset, err := service.assetRepo.GetAsset(ctx, ID)
	if err != nil {
		fmt.Printf("asset_service error while getting asset by it's ID: %s", err.Error())
		return nil, err
	}

	return asset, err
}
