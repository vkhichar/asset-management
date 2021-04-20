package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/service"
)

func AuthenticationHandler(tokenService service.TokenService, next http.HandlerFunc, adminApi bool) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := ReadToken(r)
		if err != nil {
			WriteErrorResponse(w, err)
			return
		}
		claims, err := VerifyToken(tokenService, token, adminApi)
		if err != nil {
			WriteErrorResponse(w, err)
			return
		}

		context := context.WithValue(r.Context(), "claims", claims)

		next.ServeHTTP(w, r.WithContext(context))
	})

}

func ReadToken(r *http.Request) (string, error) {
	header := r.Header.Get("Authorization")
	if header == "" {
		return "", customerrors.ErrMissingToken
	}
	if !strings.Contains(header, "Bearer ") {
		return "", customerrors.ErrInvalidToken
	}
	tokenStr := header[len("Bearer "):]
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

func UserAuthenticationHandler(tokenService service.TokenService, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := ReadToken(r)
		if err != nil {
			WriteErrorResponse(w, err)
			return
		}
		claims, err := tokenService.ValidateToken(token)
		if err != nil {
			fmt.Printf("handler: %s\n", err.Error())
			return
		}

		context := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(context))
	})
}
