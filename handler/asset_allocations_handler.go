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

func CreateAssetAllocationHandler(assetAllocationService service.AssetAllocationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		assetId, err := uuid.Parse(vars["asset_id"])

		if err != nil {
			fmt.Printf("handler:incorrect asset id")
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid asset id"})
			w.Write(responseBytes)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		var request contract.CreateAssetAllocationRequestForJson
		claims := r.Context().Value("claims")
		allocatedBy := claims.(*service.Claims).UserID

		err = json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			fmt.Printf("handler: error while decoding request for creating maintenance activity for assets: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		req := contract.NewAssetAllocationRequest(assetId, allocatedBy, request)

		assetAllocation, err := assetAllocationService.CreateAssetAllocation(r.Context(), req)

		if err != nil {
			if err == customerrors.UserNotExist {
				fmt.Println("handler: user for this id does not exist")
				w.WriteHeader(http.StatusBadRequest)
				responseBytes, err := json.Marshal(contract.ErrorResponse{customerrors.UserNotExist.Error()})
				if err != nil {
					fmt.Printf("handler: error while converting error to json, error:%s", err)
					return
				}
				w.Write(responseBytes)
				return
			}
			if err == customerrors.AssetDoesNotExist {
				fmt.Println("handler: asset for this id does not exist")
				w.WriteHeader(http.StatusBadRequest)
				responseBytes, err := json.Marshal(contract.ErrorResponse{customerrors.AssetDoesNotExist.Error()})
				if err != nil {
					fmt.Printf("handler: error while converting error to json, error:%s", err)
					return
				}
				w.Write(responseBytes)
				return
			}
			if err == customerrors.AdminDoesNotExist {
				fmt.Println("handler: admin id is incorrect")
				w.WriteHeader(http.StatusBadRequest)
				responseBytes, err := json.Marshal(contract.ErrorResponse{customerrors.AdminDoesNotExist.Error()})
				if err != nil {
					fmt.Printf("handler: error while converting error to json, error:%s", err)
					return
				}
				w.Write(responseBytes)
				return
			}
			if err == customerrors.AssetCannotBeAllocated {
				fmt.Println("handler: this asset is either retired or is under maintenance")
				w.WriteHeader(http.StatusBadRequest)
				responseBytes, err := json.Marshal(contract.ErrorResponse{customerrors.AssetCannotBeAllocated.Error()})
				if err != nil {
					fmt.Printf("handler: error while converting error to json, error:%s", err)
					return
				}
				w.Write(responseBytes)
				return
			}
			if err == customerrors.AssetAlreadyAllocated {
				fmt.Println("handler: this asset is already allocated to another user")
				w.WriteHeader(http.StatusBadRequest)
				responseBytes, err := json.Marshal(contract.ErrorResponse{customerrors.AssetAlreadyAllocated.Error()})
				if err != nil {
					fmt.Printf("handler: error while converting error to json, error:%s", err)
					return
				}
				w.Write(responseBytes)
				return
			}
			fmt.Printf("handler: error while asset allocation, error= %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if err != nil {
				fmt.Printf("handler: Error while converting error to json, error:%s", err)
				return
			}
			w.Write(responseBytes)
			return
		}

		resp := contract.NewCreateAssetAllocationResponse(assetAllocation)
		responseBytes, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("handler: error while marshaling")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBytes)
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
