package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/service"
)

func DeleteAssetHandler(asset service.AssetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		Id, fault := uuid.Parse(vars["Id"])
		if fault != nil {
			fmt.Println("handler: Invalid UUID")

			w.WriteHeader(http.StatusNotFound)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "Invalid UUID"})
			if err != nil {
				fmt.Printf("handler:Something went wrong while Marshaling,%s ", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(responseBytes)
			return

		}

		asset, err := asset.DeleteAsset(r.Context(), Id)
		if err == customerrors.AssetAlreadyDeleted {
			fmt.Println("handler: Asset Already Deleted")

			w.WriteHeader(http.StatusNotFound)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "Asset Already Deleted"})
			if err != nil {
				fmt.Printf("handler:Something went wrong while Marshaling,%s ", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(responseBytes)
			return
		}
		if err == customerrors.NoAssetsExist {
			fmt.Println("handler: No assets exist")

			w.WriteHeader(http.StatusNotFound)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "no asset found"})
			if err != nil {
				fmt.Printf("handler:Something went wrong while Marshaling,%s ", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(responseBytes)
			return
		}
		if err != nil {
			fmt.Printf("handler:Error while Searching for Updated Asset, %s", err.Error())
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		assetRep := contract.DomainToContractassets(asset)
		responseBytes, bug := json.Marshal(assetRep)
		if bug != nil {
			fmt.Println("handler:Something went wrong while marshalingh assetRep")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		return

	}
}

func UpdateAssetHandler(asset service.AssetService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(r)
		Id, fault := uuid.Parse(vars["Id"])

		if fault != nil {
			fmt.Println("handler: Invalid UUID")

			w.WriteHeader(http.StatusNotFound)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "Invalid UUID"})
			if err != nil {
				fmt.Printf("handler:Something went wrong while Marshaling,%s ", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(responseBytes)
			return
		}
		var req contract.UpdateRequest
		issue := json.NewDecoder(r.Body).Decode(&req)
		if issue != nil {
			fmt.Println("handler: Something went wrong", issue.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		asset, err := asset.UpdateAsset(r.Context(), Id, req)
		if err == customerrors.AssetAlreadyDeleted {
			fmt.Println("handler: Asset Already Deleted")

			w.WriteHeader(http.StatusNotFound)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "Asset Already Deleted"})
			if err != nil {
				fmt.Printf("handler:Something went wrong while Marshaling,%s ", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(responseBytes)
			return
		}
		if err == customerrors.NoAssetsExist {
			fmt.Println("handler: No asset found")

			w.WriteHeader(http.StatusNotFound)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "no asset found"})
			if err != nil {
				fmt.Printf("handler:Something went wrong while Marshaling,%s", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler:Error while Searching for Updated Asset, %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		assetRep := contract.DomainToContractassets(asset)
		responseBytes, bug := json.Marshal(assetRep)
		if bug != nil {
			fmt.Println("handler:Something went wrong while marshalingh assetRep")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		return

	}

}

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
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
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

		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		return

	}

}
