package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/service"
)

func AuthenticationHandler(tokenService service.TokenService, next http.HandlerFunc, adminApi bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := ReadTokenFromRequest(r)
		if err != nil {
			WriteErrorResponse(w, err)
			return
		}
		_, err = VerifyToken(tokenService, token, adminApi)
		if err != nil {
			WriteErrorResponse(w, err)
			return
		}
		next.ServeHTTP(w, r)
	})

}

func ReadTokenFromRequest(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", customerrors.ErrMissingToken
	}
	if !strings.Contains(header, "Bearer ") {
		return "", customerrors.ErrInvalidToken
	}
	tokenStr := header[len("Bearer ")-1:]
	if len(tokenStr) == 0 {
		return "", customerrors.ErrMissingToken
	}
	return tokenStr, nil
}

func VerifyToken(tokenService service.TokenService, token string, adminApi bool) (*service.Claims, error) {

	claims, err := tokenService.ValidateToken(token)
	if err != nil {
		fmt.Printf("handler: %s\n", err.Error())
		return nil, customerrors.ErrInvalidToken
	}
	if adminApi && !claims.IsAdmin {
		fmt.Printf("handler: non-admin user\n")
		return nil, customerrors.ErrForbidden
	}
	return claims, nil
}
