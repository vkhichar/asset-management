package contract

import (
	"time"
)

type CSVAsset struct {
	ID             string    `csv:"ID"`
	Status         string    `csv:"Stastus"`
	Category       string    `csv:"Category"`
	PurchaseAt     time.Time `csv:"PurchaseAt"`
	PurchaseCost   float64   `csv:"PurchaseCost"`
	Name           string    `csv:"Name"`
	Specifications string    `csv:"Specifications"`
}
