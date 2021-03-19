package handler

import (
	"net/http"

	"github.com/vkhichar/asset-management/service"
)

func CreateAssetAllocationHandler(assetAllocation service.AssetAllocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
