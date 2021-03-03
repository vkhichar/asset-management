package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/service"
)

func LoginHandler(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Set Content-Type for response
		w.Header().Set("Content-Type", "application/json")

		var req contract.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			fmt.Printf("handler: error while decoding request for login: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: invalid request for email: %s", req.Email)

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			w.Write(responseBytes)
			return
		}

		user, token, err := userService.Login(r.Context(), req.Email, req.Password)
		if err == service.ErrInvalidEmailPassword {
			fmt.Printf("handler: invalid email or password for email: %s", req.Email)

			w.WriteHeader(http.StatusUnauthorized)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid email or password"})
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while logging in for email: %s, error: %s", req.Email, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.LoginResponse{IsAdmin: user.IsAdmin, Token: token})
		w.Write(responseBytes)
	}
}

//create user handler
func CreateUserHandler(userService service.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Context-Type", "application/json")

		var req contract.CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {

			fmt.Printf("handler: error while decoding request for create user: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: invalid request for email: %s", req.Email)

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: err.Error()})
			w.Write(responseBytes)
			return
		}

		user, err := userService.CreateUser(r.Context(), req.Name, req.Email, req.Password, req.IsAdmin)

		if err == service.ErrInvalidEmailPassword {
			fmt.Printf("handler: invalid email or password for email: %s", req.Email)

			w.WriteHeader(http.StatusUnauthorized)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "invalid email or password"})
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while creating user for email: %s, error: %s", req.Email, err.Error())

			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, _ := json.Marshal(contract.CreateUserResponse{Name: user.Name, Email: user.Email, Password: user.Password, IsAdmin: user.IsAdmin})
		w.Write(responseBytes)

	}

}
