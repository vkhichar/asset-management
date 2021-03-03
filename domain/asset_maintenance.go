package domain

import (
	"time"
)

type UUID [16]byte
type MaintenanceActivity struct {
	ID          int       `db:"id"`
	AssetsID    UUID      `db:"assets_id"`
	Cost        int       `db:"cost"`
	StartedAt   time.Time `db:"started_at"`
	EndedAt     time.Time `db:"ended_at"`
	Description string    `db:"description"`
}
