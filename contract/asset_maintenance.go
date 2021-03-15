package contract

import (
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
)

const (
	DateFormat = "2006-01-02" // yyyy-mm-dd
	DateRegex  = "^[0-9]{4}-[0-9]{2}-[0-9]{2}$"
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
		StartedAt:   domain.StartedAt.Format(DateFormat),
		EndedAt:     domain.EndedAt.Format(DateFormat),
		AssetId:     domain.AssetId,
	}
}

func (activity UpdateMaintenanceActivityReq) ToDomain() (*domain.MaintenanceActivity, error) {
	endedAt, err := time.Parse(DateFormat, activity.EndedAt)
	if err != nil {
		return nil, err
	}
	return &domain.MaintenanceActivity{
		Cost:        activity.Cost,
		EndedAt:     endedAt,
		Description: activity.Description,
	}, nil
}

func (req UpdateMaintenanceActivityReq) Validate() bool {
	if req.Cost == 0.0 || strings.TrimSpace(req.EndedAt) == "" || strings.TrimSpace(req.Description) == "" {
		return false
	}

	if matched, _ := regexp.MatchString(DateRegex, req.EndedAt); !matched {
		return false
	}

	return true
}
