package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/service"
)

var (
	ErrMissingToken = errors.New("missing token")
	ErrInvalidToken = errors.New("invalid or expired token")
	ErrForbidden    = errors.New("forbidden")
	ErrBadRequest   = errors.New("bad request")
)

func ReadTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("Missing token")
	}
	tokenStr := strings.Split(header, " ")
	if len(tokenStr) != 2 {
		return "", errors.New("Missing token")
	}

	return tokenStr[1], nil
}

func verifyToken(r *http.Request, w http.ResponseWriter, adminApi bool) (*service.Claims, error) {
	token, err := ReadTokenFromRequest(r)

	if err != nil {
		fmt.Printf("handler: invalid token: %s", err.Error())
		return nil, ErrMissingToken
	}

	claims, err := deps.tokenService.ValidateToken(token)

	if err != nil {
		fmt.Printf("handler: invalid token: %s", err.Error())
		return nil, ErrInvalidToken
	}

	if adminApi && !claims.IsAdmin {
		fmt.Printf("handler: non-admin user")
		return nil, ErrForbidden
	}

	return claims, nil
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	var statusCode int
	switch err {
	case ErrMissingToken:
		statusCode = http.StatusBadRequest
	case ErrInvalidToken:
		statusCode = http.StatusUnauthorized
	case ErrForbidden:
		statusCode = http.StatusForbidden
	case ErrBadRequest:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError

	}
	w.WriteHeader(statusCode)
	responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
	w.Write(responseBytes)
}
