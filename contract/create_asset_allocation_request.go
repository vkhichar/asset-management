package contract

import (
	"github.com/google/uuid"
)

type CreateAssetAllocationRequest struct {
	ID            int       `json:"id"`
	AssetId       uuid.UUID `json:"asset_id"`
	UserId        int       `json:"user_id"`
	AllocatedBy   string    `json:"allocated_by"`
	AllocatedFrom string    `json:"allocated_from"`
	AllocatedTill string    `json:"allocated_till"`
}
