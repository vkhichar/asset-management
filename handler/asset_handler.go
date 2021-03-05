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
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "no asset found"})
			if err != nil {
				fmt.Printf("handler: Something went wrong while Marshaling: %s", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(responseBytes)
			return
		}
		if err != nil {
			fmt.Printf("handler:Error while Searching for Assets, %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		assetResp := make([]contract.Asset, 0)
		for _, u := range asset {
			assetResp = append(assetResp, contract.DomainToContractassets(&u))
		}

		responseBytes, err := json.Marshal(assetResp)

		if err != nil {

			fmt.Printf("handler: Something Went Wrong while Marshaling assets: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(responseBytes)
		return

	}
}

func CreateAssetHandler(assetService service.AssetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Set Content-Type for response
		w.Header().Set("Content-Type", "application/json")

		var req contract.CreateAssetRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Printf("handler: error while decoding request for create asset: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: invalid request for create asset: Check for proper fields ")
			fmt.Printf(err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			w.Write(responseBytes)
			return
		}

		toAsset, _ := req.ConvertToAsset()

		returnedAsset, err := assetService.CreateAsset(r.Context(), toAsset)

		if err != nil {
			fmt.Printf("handler: error while creating asset, error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.CreateAssetResponse{ID: returnedAsset.AssetID, Status: returnedAsset.Status, Category: returnedAsset.Category, PurchaseAt: returnedAsset.PurchaseAt.String(), PurchaseCost: returnedAsset.PurchaseCost, Name: returnedAsset.AssetName, Specifications: returnedAsset.Specifications})
		w.Write(responseBytes)
		return
	}
}
