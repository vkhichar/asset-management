package contract

import (
	"time"

	"github.com/google/uuid"
)

type AssetMaintaintenanceResponse struct {
	ID          int       `json:"id"`
	AssetId     uuid.UUID `json:"asset_id"`
	Cost        float64   `json:"cost"`
	StartedAt   time.Time `json:"started_at"`
	Description string    `json:"description"`
}
