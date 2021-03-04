package contract

import (
	"encoding/json"
)

type CreateAssetRequest struct {
	//AssetID       int    `json:"id"`
	Status         string  `json:"status"`
	Category       string  `json:"category"`
	PurchaseAt     string  `json:"purchase_at"`
	PurchaseCost   float64 `json:"purchase_cost"`
	AssetName      string  `json:"name"`
	Specifications json.RawMessage  `json:"specifications"`
}
