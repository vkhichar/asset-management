package domain

import (
	"time"

	"github.com/google/uuid"
)

type AssetAllocations struct {
	ID            int        `db:"id"`
	AssetId       uuid.UUID  `db:"asset_id"`
	UserId        int        `db:"user_id"`
	AllocatedBy   string     `db:"allocated_by"`
	AllocatedFrom time.Time  `db:"allocated_from"`
	AllocatedTill *time.Time `db:"allocated_till"`
}
