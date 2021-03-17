package domain

import (
	"time"

	"github.com/google/uuid"
)

type MaintenanceActivity struct {
	ID          int        `db:"id"`
	AssetId     uuid.UUID  `db:"asset_id"`
	Cost        float64    `db:"cost"`
	StartedAt   time.Time  `db:"started_at"`
	EndedAt     *time.Time `db:"ended_at"`
	Description string     `db:"description"`
}
