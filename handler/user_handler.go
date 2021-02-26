package handler

import (
	"context"
	"net/http"
)

func LoginHandler(deps dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// handling logic

		user, token, err := deps.userService.Login(context.Background(), "example@gmail.com", "password")
		_ = user
		_ = token
		_ = err
	}
}
