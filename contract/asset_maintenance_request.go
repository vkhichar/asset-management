package contract

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vkhichar/asset-management/domain"
)

type AssetMaintenanceReq struct {
	Cost        float64 `json:"cost"`
	StartedAt   string  `json:"started_at"`
	Description string  `json:"description"`
}

func (req AssetMaintenanceReq) ConvertReqFormat(assetId uuid.UUID) (domain.MaintenanceActivity, error) {
	t, err := time.Parse("02-01-2006", req.StartedAt)
	if err != nil {
		maintenanceAct := domain.MaintenanceActivity{}
		return maintenanceAct, err
	}
	tempCreateAssetMaintenance := domain.MaintenanceActivity{
		AssetId:     assetId,
		Cost:        req.Cost,
		StartedAt:   t,
		Description: req.Description,
	}

	return tempCreateAssetMaintenance, nil

}

func (req AssetMaintenanceReq) Validate() error {
	if req.Cost < 0 {
		return errors.New("cost cannot be negative")
	}

	if strings.TrimSpace(req.Description) == "" {
		return errors.New("description is required")
	}

	return nil
}
