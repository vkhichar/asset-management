package contract

import (
	"github.com/vkhichar/asset-management/domain"
)

type CreateAssetMaintenanceEventRequest struct {
	EventType string
	Data      *domain.MaintenanceActivity
}
