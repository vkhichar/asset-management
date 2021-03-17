package contract

import "github.com/vkhichar/asset-management/domain"

type UpdateUserEventRequest struct {
	EventType string
	Data      *domain.User
}
