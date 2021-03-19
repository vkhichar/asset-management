package contract

import (
	"encoding/json"

	"github.com/vkhichar/asset-management/domain"
)

type CreateAssetMaintenanceEventRequest struct {
	EventType string
	Data      *domain.MaintenanceActivity
}

type EventSubmitRequest struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func NewEventRequest(eventType string, data json.RawMessage) EventSubmitRequest {
	return EventSubmitRequest{
		Type: eventType,
		Data: data,
	}
}
