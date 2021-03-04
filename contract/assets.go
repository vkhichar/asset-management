package contract

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
)

type Asset struct {

	Id             uuid.UUID       `json:"id"` //uuid package read

	Status         string          `json:"status"`
	Category       string          `json:"category"`
	PurchaseAt     time.Time       `json:"purchase_at"`
	PurchaseCost   float64         `json:"purchase_cost"`
	Name           string          `json:"name"`
<<<<<<< HEAD
	Specifications json.RawMessage `json:"specifications"`
}

func DomainToContractassets(d *domain.Asset) Asset {

=======
	Specifications json.RawMessage `json:"specifications"` //encoding json read
}

func DomainToContractassets(d *domain.Asset) Asset {
>>>>>>> ASSET MASTER R
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
