package contract

import "encoding/json"

type UpdateRequest struct {
	Status         *string         `json:"status"`
	PurchaseCost   *float64        `json:"purchasecost"`
	Specifications json.RawMessage `json:"specifications"`
}
