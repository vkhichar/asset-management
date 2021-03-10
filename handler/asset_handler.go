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
		vars := mux.Vars(r)
		Id, fault := uuid.Parse(vars["Id"])
		if fault != nil {

			fmt.Println("handler: Something went wrong while converting string to uuid")
		}

		w.Header().Set("Content-Type", "application/json")

		asset, err := asset.DeleteAsset(r.Context(), Id)
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

		vars := mux.Vars(r)
		Id, fault := uuid.Parse(vars["Id"])

		if fault != nil {

			fmt.Println("handler: Something went wrong while converting string to uuid")
		}

		w.Header().Set("Content-Type", "application/json")

		var req contract.UpdateRequest
		issue := json.NewDecoder(r.Body).Decode(&req)
		var m map[string]interface{}
		f := json.Unmarshal([]byte(req.Specifications), &m)
		if f != nil {
			fmt.Println("hello")
		}

		fmt.Println(m)

		if issue != nil {
			fmt.Println("handler: Something went wrong", issue.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		asset, err := asset.UpdateAsset(r.Context(), Id, req)
		if err == customerrors.AssetAlreadyDeleted {
			fmt.Println("handler: Asset Already Deleted")

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
		if err == customerrors.NoAssetsExist {
			fmt.Println("handler: No assets exist")

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
				fmt.Printf("handler:Something went wrong while Marshaling,%s ", err.Error())
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
			fmt.Printf("handler:Something Went Wrong while Marshaling assetRepo,%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		return

	}

}
