package contract

import "time"

type AssetMaintain struct {
	Cost        int       `json:"cost"`
	StartedAt   time.Time `json:"started_at"`
	Description string    `json:"description"`
}
