package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	Id             uuid.UUID       `db:"id"` //uuid package read
	Status         string          `db:"status"`
	Category       string          `db:"category"`
	PurchaseAt     time.Time       `db:"purchase_at"`
	PurchaseCost   float64         `db:"purchase_cost"`
	Name           string          `db:"name"`
	Specifications json.RawMessage `db:"specifications"` //encoding json read
}
