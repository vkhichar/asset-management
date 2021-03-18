package contract

import (
	"encoding/json"
)

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
