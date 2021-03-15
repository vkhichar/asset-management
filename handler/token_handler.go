package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/service"
)

func GenerateTokenHandler(tokenService service.TokenService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var req contract.TokenClaims
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		claims := service.Claims{
			UserID:  req.UserId,
			IsAdmin: req.IsAdmin,
		}
		token, err := tokenService.GenerateToken(&claims)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(contract.TokenResponse{Token: token})
		w.Write(res)
	}
}

func VerifyTokenHandler(tokenService service.TokenService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		token, err := ReadTokenFromRequest(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			w.Write(responseBytes)
			return
		}

		claims, err := tokenService.ValidateToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		res, _ := json.Marshal(contract.TokenClaims{UserId: claims.UserID, IsAdmin: claims.IsAdmin})
		w.Write(res)
	}
}

func ReadTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", errors.New("missing token")
	}
	if !strings.Contains(header, "Bearer ") {
		return "", errors.New("missing token")
	}
	tokenStr := header[len("Bearer "):]
	if len(tokenStr) == 0 {
		return "", errors.New("missing token")
	}
	return tokenStr, nil
}
