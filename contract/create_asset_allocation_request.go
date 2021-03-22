package contract

import "github.com/google/uuid"

type CreateAssetAllocationRequest struct {
	AssetId     uuid.UUID `json:"asset_id"`
	UserId      int       `json:"user_id"`
	AllocatedBy int       `json:"allocated_by"`
}

func NewAssetAllocationRequest(assetId uuid.UUID, AllocatedBy int, userId int) CreateAssetAllocationRequest {
	req := CreateAssetAllocationRequest{
		AssetId:     assetId,
		UserId:      userId,
		AllocatedBy: AllocatedBy,
	}
	return req
}
