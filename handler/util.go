package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
)

func WriteErrorResponse(w http.ResponseWriter, err error) {
	var statusCode int
	switch err {
	case customerrors.ErrMissingToken:
		statusCode = http.StatusBadRequest
	case customerrors.ErrInvalidToken:
		statusCode = http.StatusUnauthorized
	case customerrors.ErrForbidden:
		statusCode = http.StatusForbidden
	case customerrors.ErrBadRequest:
		statusCode = http.StatusBadRequest
	case customerrors.ErrNotFound:
		statusCode = http.StatusNotFound
	default:
		statusCode = http.StatusInternalServerError

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
	w.Write(responseBytes)
}
