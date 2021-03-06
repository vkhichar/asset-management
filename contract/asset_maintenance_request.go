package contract

import (
	"errors"
	"strings"
	"time"

	"github.com/vkhichar/asset-management/domain"
)

type AssetMaintenanceReq struct {
	Cost        float64 `json:"cost"`
	StartedAt   string  `json:"started_at"`
	Description string  `json:"description"`
}

func (req AssetMaintenanceReq) ConvertReqFormat() (domain.MaintenanceActivity, error) {
	t, err := time.Parse("02-01-2006", req.StartedAt)
	if err != nil {
		return domain.MaintenanceActivity{}, err
	}
	tempcreateassetmaintenance := domain.MaintenanceActivity{
		Cost:        req.Cost,
		StartedAt:   t,
		Description: req.Description,
	}

	return tempcreateassetmaintenance, nil

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
