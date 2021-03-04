package contract

import (
	"time"

	"github.com/vkhichar/asset-management/domain"
)

type AssetMaintain struct {
	AssetId     domain.UUID `json:"asset_id"`
	Cost        int         `json:"cost"`
	StartedAt   time.Time   `json:"started_at"`
	Description string      `json:"description"`
}
