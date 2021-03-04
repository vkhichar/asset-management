package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/service"
)

func ListAssetHandler(asset service.AssetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		asset, err := asset.ListAssets(r.Context())
		if err == customerrors.NoAssetsExist {
			fmt.Println("handler: No assets exist")
			w.WriteHeader(http.StatusNotFound)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "no asset found"})
			w.Write(responseBytes)
			return
		}
		if err != nil {
			fmt.Printf("handler:Error while Searching for Assets, %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		assetResp := make([]contract.Asset, 0)
		for _, u := range asset {
			assetResp = append(assetResp, contract.DomainToContractassets(&u))
		}
		responseBytes, err := json.Marshal(assetResp)
		w.Write(responseBytes)
		return

	}

}