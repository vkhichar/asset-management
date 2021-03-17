package contract

import (
	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
)

type AssetMaintenanceResponse struct {
	ID          int       `json:"id"`
	AssetId     uuid.UUID `json:"asset_id"`
	Cost        float64   `json:"cost"`
	StartedAt   string    `json:"started_at"`
	Description string    `json:"description"`
}

func NewAssetMaintenanceResponse(domain *domain.MaintenanceActivity) AssetMaintenanceResponse {
	return AssetMaintenanceResponse{
		ID:          domain.ID,
		AssetId:     domain.AssetId,
		Cost:        domain.Cost,
		StartedAt:   domain.StartedAt.Format("02-01-2006"),
		Description: domain.Description,
	}
}

type DetailAssetMaintenanceActivityResponse struct {
	ID          int       `json:"id"`
	AssetId     uuid.UUID `json:"asset_id"`
	Cost        float64   `json:"cost"`
	StartedAt   string    `json:"started_at"`
	EndedAt     string    `json:"ended_at"`
	Description string    `json:"description"`
}

func NewDetailAssetMaintenanceActivityResponse(domain *domain.MaintenanceActivity) DetailAssetMaintenanceActivityResponse {
	return DetailAssetMaintenanceActivityResponse{
		ID:          domain.ID,
		AssetId:     domain.AssetId,
		Cost:        domain.Cost,
		StartedAt:   domain.StartedAt.Format("02-01-2006"),
		EndedAt:     domain.EndedAt.Format("02-01-2006"),
		Description: domain.Description,
	}
}
