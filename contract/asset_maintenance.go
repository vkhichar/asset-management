package contract

import (
	"time"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
)

const (
	DATE_FORMAT = "2006-01-02" // yyyy-mm-dd
)

type AssetMaintain struct {
	Cost        float64   `json:"cost"`
	StartedAt   time.Time `json:"started_at"`
	Description string    `json:"description"`
}

type UpdateMaintenanceActivityReq struct {
	Cost        float64 `json:"cost"`
	EndedAt     string  `json:"ended_at"`
	Description string  `json:"description"`
}

type MaintenanceActivityResp struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Cost        float64   `json:"cost"`
	StartedAt   string    `json:"started_at"`
	EndedAt     string    `json:"ended_at"`
	AssetId     uuid.UUID `json:"asset_id"`
}

func NewMaintenanceActivityResp(domain domain.MaintenanceActivity) MaintenanceActivityResp {
	return MaintenanceActivityResp{
		Id:          domain.ID,
		Description: domain.Description,
		Cost:        domain.Cost,
		StartedAt:   domain.StartedAt.Format(DATE_FORMAT),
		EndedAt:     domain.EndedAt.Format(DATE_FORMAT),
		AssetId:     domain.AssetId,
	}
}

func (activity UpdateMaintenanceActivityReq) ConvertToDomain() domain.MaintenanceActivity {
	endedAt, _ := time.Parse(DATE_FORMAT, activity.EndedAt)

	return domain.MaintenanceActivity{
		Cost:        activity.Cost,
		EndedAt:     endedAt,
		Description: activity.Description,
	}
}
