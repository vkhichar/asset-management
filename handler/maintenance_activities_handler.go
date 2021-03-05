package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/service"
)

func CreateMaintenanceHandler(assetMaintenanceService service.AssetMaintenanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		AssetID, err1 := uuid.Parse(vars["assetId"])
		if err1 != nil {
			fmt.Printf("handler:incorrect asset id")
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid asset id"})
			w.Write(responseBytes)
			return
		}
		// Set Content-Type for response
		w.Header().Set("Content-Type", "application/json")

		var req contract.AssetMaintenance
		req.AssetId = AssetID
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
		createassetmaintenance := domain.MaintenanceActivity{
			AssetId:     AssetID,
			Cost:        req.Cost,
			StartedAt:   req.StartedAt,
			Description: req.Description,
		}

		assetmaintenance, err := assetMaintenanceService.CreateAssetMaintenance(r.Context(), createassetmaintenance)
		if err != nil {
			fmt.Printf("handler: error: %s", err.Error())

			w.WriteHeader(http.StatusNotFound)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "This asset id does not exist... "})
			w.Write(responseBytes)
			return
		}
		w.WriteHeader(http.StatusCreated)
		responseBytes, _ := json.Marshal(contract.AssetMaintaintenanceResponse{ID: assetmaintenance.ID, AssetId: assetmaintenance.AssetId, Cost: assetmaintenance.Cost, StartedAt: assetmaintenance.StartedAt, Description: assetmaintenance.Description})
		w.Write(responseBytes)
	}
}
