package contract

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

type assetCategory string

const (
	categoryLaptop  = "laptop"
	categoryMouse   = "mouse"
	cateoryKeyboard = "keyboard"
	categoryMonitor = "monitor"
	categoryCPU     = "cpu"
	categoryPrinter = "printer"
)

const PurchaseDateRegex = "^[0-9]{2}-[0-9]{2}-[0-9]{4}$"

type CreateAssetRequest struct {
	Category       string          `json:"category"`
	PurchaseAt     string          `json:"purchase_at"`
	PurchaseCost   float64         `json:"purchase_cost"`
	AssetName      string          `json:"name"`
	Specifications json.RawMessage `json:"specifications"`
}

func (req CreateAssetRequest) ConvertToAsset() (domain.Asset, error) {

	t, err := time.Parse("02-01-2006", req.PurchaseAt)

	if err != nil {
		fmt.Printf("create asset request convert to asset: %s", err.Error())
		return domain.Asset{}, err
	}

	tempAsset := domain.Asset{
		Id:             uuid.New(),
		Status:         "active",
		Category:       req.Category,
		PurchaseAt:     t,
		PurchaseCost:   req.PurchaseCost,
		Name:           req.AssetName,
		Specifications: req.Specifications,
	}
	return tempAsset, nil
}

func (req CreateAssetRequest) Validate() error {

	if matched, _ := regexp.MatchString(PurchaseDateRegex, req.PurchaseAt); !matched {
		return errors.New("Invalid purchase_at date format")
	}

	t, err := time.Parse("02-01-2006", req.PurchaseAt)
	if err != nil {
		fmt.Printf("create asset request validate: %s", err.Error())
		return err
	}

	if t.After(time.Now()) {
		fmt.Println("create asset request: purchase date should be on or before current date")
		return customerrors.ErrorInvalidPurchaseDate
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

	switch req.Category {
	case categoryLaptop:
	case categoryCPU:
	case categoryMonitor:
	case categoryMouse:
	case cateoryKeyboard:
	case categoryPrinter:
	default:
		return customerrors.ErrorInvalidAssetCategory
	}

	return nil
}
