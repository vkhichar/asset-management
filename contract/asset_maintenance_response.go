package contract

import (
	"github.com/google/uuid"
)

type AssetMaintaintenanceResponse struct {
	ID          int       `json:"id"`
	AssetId     uuid.UUID `json:"asset_id"`
	Cost        float64   `json:"cost"`
	StartedAt   string    `json:"started_at"`
	Description string    `json:"description"`
}
