package contract

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

type assetStatus string

const (
	statusActive           assetStatus = "active"
	statusRetired          assetStatus = "retired"
	statusUnderMaintenance assetStatus = "undermaintenance"
)

type CreateAssetRequest struct {
	//AssetID       int    `json:"id"`
	Status         string          `json:"status"`
	Category       string          `json:"category"`
	PurchaseAt     string          `json:"purchase_at"`
	PurchaseCost   float64         `json:"purchase_cost"`
	AssetName      string          `json:"name"`
	Specifications json.RawMessage `json:"specifications"`
}

func (req CreateAssetRequest) ConvertToAsset() (domain.Asset, error) {
	t, err := time.Parse("02/01/2006", req.PurchaseAt)

	if err != nil {
		return domain.Asset{}, err
	}

	tempAsset := domain.Asset{
		uuid.New(),
		req.Status,
		req.Category,
		t,
		req.PurchaseCost,
		req.AssetName,
		req.Specifications,
	}
	return tempAsset, nil
}

func (req CreateAssetRequest) Validate() error {

	if req.PurchaseAt == "" {
		return errors.New("Purchase At filed required")
	}

	if req.AssetName == "" {
		return errors.New("Name is required")
	}

	var checkSpec map[string]interface{}
	if err := json.Unmarshal([]byte(req.Specifications), &checkSpec); err != nil {
		fmt.Printf("contract:create asset request cannot unmarshal specifications")
		return errors.New(err.Error())
	}

	if len(checkSpec) == 0 {
		return errors.New("Specifications required")
	}

	if req.PurchaseCost == 0.0 {
		return errors.New("Purchase Cost required")
	}

	contractStatus := req.Status

	if contractStatus == string(statusActive) || contractStatus == string(statusRetired) || contractStatus == string(statusUnderMaintenance) {
		return nil
	} else {
		return customerrors.ErrInvalidAssetStatus
	}

}
