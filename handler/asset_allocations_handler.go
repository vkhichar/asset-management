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

func CreateAssetAllocationHandler(assetAllocation service.AssetAllocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}

func AssetDeAllocationHandler(assetAllocation service.AssetAllocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, err := uuid.Parse(params["asset_id"])
		if err != nil {
			fmt.Printf("asset allocation handler: error while parsing string into UUID: %s", err.Error())
			resopnseBytes, err := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			if err != nil {
				fmt.Printf(err.Error())
				return
			}
			w.Write(resopnseBytes)
			return
		}

		validID := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

		if !validID.MatchString(id.String()) {
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

		returnedMsg, err := assetAllocation.AssetDeallocation(r.Context(), id)

		if err == customerrors.ErrDeallocatedAlready {
			fmt.Printf("asset allocation handler: error asset deallocated already: %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "asset deallocated already"})
			if err != nil {
				fmt.Printf("error while marshalling error message: %s", err.Error())
				return
			}
			w.Write(responseBytes)
			return
		}

		if err != nil {

			fmt.Printf("asset allocation handler: error while deallocating error: %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if err != nil {
				fmt.Printf("error while marshalling error message:%s", err.Error())
				return
			}
			w.Write(responseBytes)
			return
		}

		responseBytes, err := json.Marshal(returnedMsg)
		if err != nil {
			fmt.Printf("asset allocation handler: error while marshaling return msg: %s", err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
		return

	}
}
