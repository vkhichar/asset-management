package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
	ExportAssetsToCSV(ctx context.Context) ([]contract.CSVAsset, error)
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
	res, errevent := service.eventSvc.PostAssetEvent(ctx, asset)
	if errevent != nil {
		fmt.Println("Service :Error in PostAssetEvent")

	} else {
		fmt.Println("New Id Created", res)
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
		fmt.Printf("asset_service error while creating asset: %s\n", err.Error())
		return nil, err
	}

	id, err := service.eventSvc.PostCreateAssetEvent(ctx, asset)
	if err != nil {
		fmt.Printf("asset service: error during post create asset event: %s\n", err.Error())
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

func (service *assetService) ExportAssetsToCSV(ctx context.Context) ([]contract.CSVAsset, error) {
	assets, err := service.assetRepo.ListAssets(ctx)

	if err != nil {
		return nil, err
	}

	if assets == nil {
		return nil, customerrors.NoAssetsExist
	}

	assetList := make([]contract.CSVAsset, len(assets))
	for i, val := range assets {
		assetList[i] = contract.CSVAsset{
			ID:             val.Id.String(),
			Status:         val.Status,
			Category:       val.Category,
			PurchaseAt:     val.PurchaseAt,
			PurchaseCost:   val.PurchaseCost,
			Name:           val.Name,
			Specifications: toString(val.Specifications),
		}
	}

	return assetList, nil
}

func toString(specs json.RawMessage) string {
	var specifications interface{}

	bytes, _ := specs.MarshalJSON()

	_ = json.Unmarshal(bytes, &specifications)

	result := Decode(specifications.(map[string]interface{}))

	return result
}

func Decode(specs map[string]interface{}) string {
	var specifications strings.Builder

	for k, v := range specs {

		var keyvalue strings.Builder
		keyvalue.WriteString(k + ":")

		switch vv := v.(type) {
		case string:
			keyvalue.WriteString(vv)
		case float64:
			keyvalue.WriteString(strconv.FormatFloat(vv, 'f', -1, 64))
		case bool:
			keyvalue.WriteString(strconv.FormatBool(vv))
		case nil:
			keyvalue.WriteString("nil")
		}

		specifications.WriteString(keyvalue.String() + ",")
	}

	result := specifications.String()
	size := len(result)

	if size > 0 && result[size-1] == ',' {
		result = result[:size-1]
	}

	return result
}
