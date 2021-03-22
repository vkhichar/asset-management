package contract

import "github.com/google/uuid"

type CreateAssetAllocationRequestForJson struct {
	UserId int `json:"user_id"`
}

type CreateAssetAllocationRequest struct {
	AssetId     uuid.UUID `json:"asset_id"`
	UserId      int       `json:"user_id"`
	AllocatedBy int       `json:"allocated_by"`
}

func NewAssetAllocationRequest(assetId uuid.UUID, AllocatedBy int, c CreateAssetAllocationRequestForJson) CreateAssetAllocationRequest {
	req := CreateAssetAllocationRequest{
		AssetId:     assetId,
		UserId:      c.UserId,
		AllocatedBy: AllocatedBy,
	}
	return req
}
