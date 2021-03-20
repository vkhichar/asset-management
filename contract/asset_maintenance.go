package contract

import (
	"errors"
	"fmt"
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
	StartedAt   string  `json:"started_at"`
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
	var endedAt string
	if domain.EndedAt != nil {
		endedAt = domain.EndedAt.Format(DateFormat)
	}
	return MaintenanceActivityResp{
		Id:          domain.ID,
		Description: domain.Description,
		Cost:        domain.Cost,
		StartedAt:   domain.StartedAt.Format(DateFormat),
		EndedAt:     endedAt,
		AssetId:     domain.AssetId,
	}
}

func (activity UpdateMaintenanceActivityReq) ToDomain() (*domain.MaintenanceActivity, error) {
	var endedAt *time.Time
	if activity.EndedAt != "" {
		date, _ := time.Parse(DateFormat, activity.EndedAt)
		endedAt = &date
	}
	startedAt, _ := time.Parse(DateFormat, activity.StartedAt)
	return &domain.MaintenanceActivity{
		Cost:        activity.Cost,
		EndedAt:     endedAt,
		Description: activity.Description,
		StartedAt:   startedAt,
	}, nil
}

func (req UpdateMaintenanceActivityReq) Validate() error {
	if req.Cost < 0.0 {
		return errors.New("cost cannot be negative")
	}

	if req.Description = strings.TrimSpace(req.Description); req.Description == "" {
		return errors.New("missing description")
	}

	currentTime := time.Now()

	if req.StartedAt = strings.TrimSpace(req.StartedAt); req.StartedAt == "" {
		return errors.New("missing date started_at")
	}
	startedAt, err := parseAndVerifyDate(currentTime, req.StartedAt, "started_at")
	if err != nil {
		return err
	}

	if req.EndedAt = strings.TrimSpace(req.EndedAt); req.EndedAt != "" {
		endedAt, err := parseAndVerifyDate(currentTime, strings.TrimSpace(req.EndedAt), "ended_at")
		if err != nil {
			return err
		}
		if endedAt.Before(*startedAt) {
			return errors.New("invalid data ended_at: cannot be before started_at")
		}
	}
	return nil
}

func parseAndVerifyDate(currentTime time.Time, date string, jsonKey string) (*time.Time, error) {
	parsedDate, err := time.Parse(DateFormat, date)
	if err != nil {
		fmt.Println("parsing date error ", err.Error())
		return nil, errors.New(fmt.Sprintf("invalid date %s", jsonKey))
	}

	if parsedDate.After(currentTime) {
		return nil, errors.New(fmt.Sprintf("date %s cannot be in future", jsonKey))
	}

	return &parsedDate, nil
}
