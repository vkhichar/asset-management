package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/service"
)

func CreateMaintenanceHandler(assetMaintenanceService service.AssetMaintenanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		assetId, eror := uuid.Parse(vars["asset_id"])
		// Set Content-Type for response
		w.Header().Set("Content-Type", "application/json")
		if eror != nil {
			fmt.Printf("handler:incorrect asset id")
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid asset id"})
			w.Write(responseBytes)
			return
		}
		var req contract.AssetMaintenanceReq
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			fmt.Printf("handler: error while decoding request for creating maintenance activity for assets: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: description or cost have invalid format")

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			w.Write(responseBytes)
			return
		}
		createAssetMaintenance, err := req.ConvertReqFormat(assetId)
		if err != nil {
			fmt.Printf("handler: incorrect date format: %s", err.Error())

			w.WriteHeader(http.StatusNotFound)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "incorrect date format "})
			w.Write(responseBytes)
			return
		}

		assetMaintenance, err := assetMaintenanceService.CreateAssetMaintenance(r.Context(), createAssetMaintenance)

		if err != nil {
			fmt.Printf("handler: error: %s", err.Error())

			w.WriteHeader(http.StatusNotFound)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "This asset id does not exist... "})
			w.Write(responseBytes)
			return
		}
		resp := contract.NewAssetMaintenanceResponse(assetMaintenance)
		w.WriteHeader(http.StatusCreated)
		responseBytes, _ := json.Marshal(resp)
		w.Write(responseBytes)
	}
}

func DetailedMaintenanceActivityHandler(service service.AssetMaintenanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		assetMaintenance, err := service.DetailedMaintenanceActivity(r.Context(), id)
		if err != nil {
			fmt.Printf("handler: error: %s", err.Error())

			w.WriteHeader(http.StatusNotFound)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "This maintenance id does not exist... "})
			w.Write(responseBytes)
			return
		}
		resp := contract.NewDetailAssetMaintenanceActivityResponse(assetMaintenance)
		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(resp)
		w.Write(responseBytes)
	}
}
