package contract

import (
	"encoding/json"

	"github.com/google/uuid"
)

type GetAssetResponse struct {
	ID             uuid.UUID       `json:"id"`
	Status         string          `json:"status"`
	Category       string          `json:"category"`
	PurchaseAt     string          `json:"purchase_at"`
	PurchaseCost   float64         `json:"purchase_cost"`
	Name           string          `json:"name"`
	Specifications json.RawMessage `json:"specifications"`
}
