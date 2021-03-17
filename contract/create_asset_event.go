package contract

import (
	"github.com/vkhichar/asset-management/domain"
)

type CreateAssetEvent struct {
	EventType string        `json:"event_type"`
	Data      *domain.Asset `json:"data"`
}
