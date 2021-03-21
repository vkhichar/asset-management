package contract

import (
	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
)

type CreateAssetAllocationResponse struct {
	ID            int       `json:"id"`
	AssetId       uuid.UUID `json:"asset_id"`
	UserId        int       `json:"user_id"`
	AllocatedBy   string    `json:"allocated_by"`
	AllocatedFrom string    `json:"allocated_from"`
}

func NewCreateAssetAllocationResponse(domain *domain.AssetAllocations) CreateAssetAllocationResponse {
	resp := CreateAssetAllocationResponse{
		ID:            domain.ID,
		AssetId:       domain.AssetId,
		UserId:        domain.UserId,
		AllocatedBy:   domain.AllocatedBy,
		AllocatedFrom: domain.AllocatedFrom.String(),
	}
	return resp
}
