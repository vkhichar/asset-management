package contract

import "github.com/vkhichar/asset-management/domain"

type CreateUserEvent struct {
	EventType string       `json:"event_type"`
	Data      *domain.User `json:"data"`
}
