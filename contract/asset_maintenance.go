package contract

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type AssetMaintenance struct {
	AssetId     uuid.UUID `json:"asset_id"`
	Cost        float64   `json:"cost"`
	StartedAt   time.Time `json:"started_at"`
	Description string    `json:"description"`
}

func (req AssetMaintenance) Validate() error {
	if req.Cost < 0 {
		return errors.New("cost cannot be negative")
	}

	if strings.TrimSpace(req.Description) == "" {
		return errors.New("description is required")
	}

	return nil
}
