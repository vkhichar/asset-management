package contract

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
)

type Asset struct {
	Id uuid.UUID `json:"id"` //uuid package read

	Status         string          `json:"status"`
	Category       string          `json:"category"`
	PurchaseAt     time.Time       `json:"purchase_at"`
	PurchaseCost   float64         `json:"purchase_cost"`
	Name           string          `json:"name"`
	Specifications json.RawMessage `json:"specifications"`
}

func DomainToContractassets(d *domain.Asset) Asset {

	u := Asset{
		Id:             d.Id,
		Status:         d.Status,
		Category:       d.Category,
		PurchaseAt:     d.PurchaseAt,
		PurchaseCost:   d.PurchaseCost,
		Name:           d.Name,
		Specifications: d.Specifications,
	}
	return u
}
