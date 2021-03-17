package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

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
		Id, errParse := uuid.Parse(vars["Id"])
		if errParse != nil {
			fmt.Printf("handler: Invalid UUID %s", errParse.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "Invalid UUID"})
			if err != nil {
				fmt.Printf("handler:Something went wrong while Marshaling,%s ", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.Write(responseBytes)
			return

		}

		asset, err := asset.DeleteAsset(r.Context(), Id)

		if err != nil {
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
			fmt.Printf("handler:Error while Searching for Updated Asset, %s", err.Error())
			responseBytes, errMarshal := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if errMarshal != nil {
				fmt.Printf("handler: Error while Marshaling %s", errMarshal.Error())
			}
			w.Write(responseBytes)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		assetRep := contract.DomainToContractassets(asset)
		responseBytes, errMarshal := json.Marshal(assetRep)
		if errMarshal != nil {
			fmt.Printf("handler:Something went wrong while marshalingh assetRep %s", errMarshal.Error())
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
		Id, errParse := uuid.Parse(vars["Id"])

		if errParse != nil {
			fmt.Printf("handler: Invalid UUID %s", errParse.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "Invalid UUID"})
			if err != nil {
				fmt.Printf("handler:Something went wrong while Marshaling,%s ", err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.Write(responseBytes)
			return
		}
		var req contract.UpdateRequest
		errDecoder := json.NewDecoder(r.Body).Decode(&req)
		if errDecoder != nil {
			fmt.Printf("handler: Something went wrong %s", errDecoder.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		asset, err := asset.UpdateAsset(r.Context(), Id, req)

		if err != nil {
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
			fmt.Printf("handler:Error while Searching for Updated Asset, %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, errMarshal := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if errMarshal != nil {
				fmt.Printf("handler: Error while Marshaling %s", errMarshal.Error())
			}
			w.Write(responseBytes)
			return
		}

		assetRep := contract.DomainToContractassets(asset)
		responseBytes, errMarshal := json.Marshal(assetRep)
		if errMarshal != nil {
			fmt.Printf("handler:Something went wrong while marshalingh assetRep %s", errMarshal.Error())
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
		w.Write(responseBytes)
		return

	}
}

func CreateAssetHandler(assetService service.AssetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		var req contract.CreateAssetRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Printf("handler: error while decoding request for create asset: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: invalid request for create asset: Check for proper fields ")
			fmt.Printf(err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			w.Write(responseBytes)
			return
		}

		toAsset, err := req.ConvertToAsset()

		if err != nil {
			fmt.Printf("handler: error while converting to object of type domain.Asset, error: %s", err.Error())
			return
		}

		returnedAsset, err := assetService.CreateAsset(r.Context(), &toAsset)

		if err != nil {
			fmt.Printf("handler: error while creating asset, error: %s", err.Error())
			fmt.Printf(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, err := json.Marshal(contract.CreateAssetResponse{ID: returnedAsset.Id, Status: returnedAsset.Status, Category: returnedAsset.Category, PurchaseAt: returnedAsset.PurchaseAt.String(), PurchaseCost: returnedAsset.PurchaseCost, Name: returnedAsset.Name, Specifications: returnedAsset.Specifications})
		if err != nil {
			fmt.Printf("asset_handler: error while marshalling, %s", err.Error())
			return
		}
		w.Write(responseBytes)
		return
	}
}

func GetAssetHandler(assetService service.AssetService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)

		id, err := uuid.Parse(params["id"])
		if err != nil {
			fmt.Printf("asset handler: Error while parsing string into JSON: %s", err.Error())
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			w.Write(responseBytes)
			return
		}

		validID := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

		if validID.MatchString(id.String()) == false {
			fmt.Printf("handler: invalid request for get asset: Check for proper id ")
			fmt.Printf("Invalid UUID format")

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "Invalid UUID format"})
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			w.Write(responseBytes)
			return
		}

		returnedAsset, err := assetService.GetAsset(r.Context(), id)

		if err != nil {
			if err == customerrors.NoAssetsExist {
				fmt.Printf("handler: asset does not exist")
				w.WriteHeader(http.StatusNotFound)
				responseBytes, err := json.Marshal(contract.ErrorResponse{Error: err.Error()})
				if err != nil {
					fmt.Printf(err.Error())
					return
				}
				w.Write(responseBytes)
				return
			}

			fmt.Printf("handler: error while creating asset, error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			w.Write(responseBytes)
			return
		}

		responseBytes, err := json.Marshal(contract.GetAssetResponse{ID: returnedAsset.Id, Status: returnedAsset.Status, Category: returnedAsset.Category, PurchaseAt: returnedAsset.PurchaseAt.String(), PurchaseCost: returnedAsset.PurchaseCost, Name: returnedAsset.Name, Specifications: returnedAsset.Specifications})
		if err != nil {
			fmt.Printf("asset_handler: error while marshalling, %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		return
	}
}
