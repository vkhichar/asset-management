package contract

import (
	"github.com/vkhichar/asset-management/domain"
)

type UpdateAssetEvent struct {
	EvenType string
	Data     *domain.Asset
}
