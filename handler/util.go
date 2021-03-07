package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/service"
)

func ReadTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", customerrors.ErrMissingToken
	}
	var tokenStr string
	if !strings.Contains(header, "Bearer ") {
		return "", customerrors.ErrInvalidToken
	}

	tokenStr = header[len("Bearer ")-1:]
	if len(tokenStr) == 0 {
		return "", customerrors.ErrMissingToken
	}

	return tokenStr, nil
}

/**
*
 */

func VerifyToken(r *http.Request, w http.ResponseWriter, adminApi bool) (*service.Claims, error) {
	token, err := ReadTokenFromRequest(r)

	if err != nil {
		return nil, err
	}

	claims, err := deps.tokenService.ValidateToken(token)

	if err != nil {
		fmt.Printf("handler: invalid token: %s", err.Error())
		return nil, customerrors.ErrInvalidToken
	}

	if adminApi && !claims.IsAdmin {
		fmt.Printf("handler: non-admin user")
		return nil, customerrors.ErrForbidden
	}

	return claims, nil
}

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
	w.WriteHeader(statusCode)
	responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
	w.Write(responseBytes)
}
