package contract

import "github.com/vkhichar/asset-management/domain"

type CreateAssetResponse struct {
	CreatedObject domain.Asset `json:"object"`
}
