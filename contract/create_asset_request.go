package contract

type CreateAssetRequest struct {
	//AssetID       int    `json:"id"`
	Status         string  `json:"status"`
	Category       string  `json:"category"`
	PurchaseAt     string  `json:"purchase_at"`
	PurchaseCost   float64 `json:"purchase_cost"`
	AssetName      string  `json:"name"`
	Specifications string  `json:"specifications"`
}
