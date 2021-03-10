package contract

import (
	"encoding/json"
	"errors"
	"regexp"

	"github.com/google/uuid"
)

type GetAssetRequest struct {
	//AssetID       int    `json:"id"`
	Status         string          `json:"status"`
	Category       string          `json:"category"`
	PurchaseAt     string          `json:"purchase_at"`
	PurchaseCost   float64         `json:"purchase_cost"`
	AssetName      string          `json:"name"`
	Specifications json.RawMessage `json:"specifications"`
}

func (req GetAssetRequest) Validate(Id uuid.UUID) error {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	if r.MatchString(Id.String()) == false {
		return errors.New("get asset contract: Invalid asset id format")
	}
	return nil
}
