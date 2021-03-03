package domain

import (
	"time"
)

type UUID [16]byte

type assets struct {
	Id             UUID                   `db:"id"` //uuid package read
	Status         string                 `db:"status"`
	Category       string                 `db:"category"`
	Purchaseat     time.Time              `db:"purchase_at"`
	Purchasecost   float64                `db:"purchase_cost"`
	Name           string                 `db:"name"`
	Specifications map[string]interface{} `db:"specifiation"` //encoding json read
}
