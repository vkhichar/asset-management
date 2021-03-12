package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/vkhichar/asset-management/contract"
<<<<<<< HEAD
<<<<<<< HEAD
	"github.com/vkhichar/asset-management/customerrors"
=======
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
>>>>>>> Added update activity Api
=======
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
>>>>>>> Code refactor
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
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "incorrect date format"})
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

		responseBytes, eror := json.Marshal(resp)
		if eror != nil {
			fmt.Printf("handler: error while marshaling")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusCreated)
		w.Write(responseBytes)
	}
}

func DetailedMaintenanceActivityHandler(service service.AssetMaintenanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, errr := strconv.Atoi(mux.Vars(r)["id"])
		if errr != nil {
			fmt.Println("handler: wrong id format")
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "wrong id format"})
			w.Write(responseBytes)
			return
		}
		assetMaintenance, err := service.DetailedMaintenanceActivity(r.Context(), id)

		if err == customerrors.MaintenanceIdDoesNotExist {
			fmt.Println("handler: Maintenance Id does not exist")
			w.WriteHeader(http.StatusNotFound)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "id not found"})
			w.Write(responseBytes)
			return

		}

		if err != nil {
			fmt.Printf("handler: error while searching for maintenance activity,error= %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}
		resp := contract.NewDetailAssetMaintenanceActivityResponse(assetMaintenance)
		responseBytes, eror := json.Marshal(resp)
		if eror != nil {
			fmt.Printf("handler: error while marshaling")
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	}
)

func DeleteMaintenanceActivityHandler(service service.AssetMaintenanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			WriteErrorResponse(w, customerrors.ErrBadRequest)
			return
		}
		err = service.DeleteMaintenanceActivity(r.Context(), id)
		if err != nil {
			WriteErrorResponse(w, errors.New("Something went wrong"))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}
}

func ListMaintenanceActivitiesByAsserId(service service.AssetMaintenanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		assetId, err := uuid.Parse(mux.Vars(r)["asset_id"])
		if err != nil {
			fmt.Println(err)
			WriteErrorResponse(w, customerrors.ErrBadRequest)
			return
		}

		activities, err := service.GetAllForAssetId(r.Context(), assetId)

		if err != nil {
			WriteErrorResponse(w, errors.New("Something went wrong"))
			return
		}
		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(convertAllActivitiesToContract(activities))
		w.Write(responseBytes)
		return
	}
}

func UpdateMaintenanceActivity(service service.AssetMaintenanceService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			fmt.Println(err)
			WriteErrorResponse(w, customerrors.ErrBadRequest)
			return
		}
		var updateReq contract.UpdateMaintenanceActivityReq
		err = json.NewDecoder(r.Body).Decode(&updateReq)
		if err != nil {
			fmt.Println(err)
			WriteErrorResponse(w, customerrors.ErrBadRequest)
			return
		}

		if !updateReq.Validate() {
			WriteErrorResponse(w, customerrors.ErrBadRequest)
			return
		}

		activityDomain, err := updateReq.ToDomain()
		if err != nil {
			WriteErrorResponse(w, customerrors.ErrBadRequest)
			return
		}

		activityDomain.ID = id

		activity, err := service.UpdateMaintenanceActivity(r.Context(), *activityDomain)

		if err == customerrors.ErrNotFound {
			WriteErrorResponse(w, customerrors.ErrNotFound)
			return
		}

		if err != nil {
			WriteErrorResponse(w, errors.New("Something went wrong"))
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.NewMaintenanceActivityResp(*activity))
		w.Write(responseBytes)

	}
}

func convertAllActivitiesToContract(activities []domain.MaintenanceActivity) []contract.MaintenanceActivityResp {
	res := make([]contract.MaintenanceActivityResp, len(activities))
	for index, value := range activities {
		res[index] = contract.NewMaintenanceActivityResp(value)
	}
	return res
}
