package contract

import (
	"time"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
)

const (
	DATE_FORMAT = "2006/01/02"
)

type AssetMaintain struct {
	Cost        float64   `json:"cost"`
	StartedAt   time.Time `json:"started_at"`
	Description string    `json:"description"`
}

type MaintenanceActivityResp struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Cost        float64   `json:"cost"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
	AssetId     uuid.UUID `json:"asset_id"`
}

func NewMaintenanceActivityResp(domain domain.MaintenanceActivity) MaintenanceActivityResp {
	return MaintenanceActivityResp{
		Id:          domain.ID,
		Description: domain.Description,
		Cost:        domain.Cost,
		StartedAt:   domain.StartedAt,
		EndedAt:     domain.EndedAt,
		AssetId:     domain.AssetId,
	}
}
