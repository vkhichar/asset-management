package domain

import "time"

type Maintenance struct {
	ID          string    `db:"id"`
	AssetsID    string    `db:"assets_id"`
	Cost        int       `db:"cost"`
	StartedAt   time.Time `db:"started_at"`
	EndedAt     time.Time `db:"ended_at"`
	Description time.Time `db:"description"`
}
