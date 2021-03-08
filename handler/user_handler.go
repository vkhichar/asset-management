package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
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
		if err == customerrors.ErrInvalidEmailPassword {
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

func ListUsersHandler(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		user, err := userService.ListUsers(r.Context())

		if err == customerrors.NoUsersExist {
			fmt.Println("handler: No users exist")
			w.WriteHeader(http.StatusNotFound)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "no user found"})
			w.Write(responseBytes)
			return

		}

		if err != nil {
			fmt.Printf("handler: error while searching for user,error= %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, _ := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			w.Write(responseBytes)
			return
		}

		userResp := make([]contract.User, 0)
		for _, u := range user {
			userResp = append(userResp, contract.DomainToContract(&u))
		}
		responsebytes, err := json.Marshal(userResp)
		w.WriteHeader(http.StatusOK)
		w.Write(responsebytes)
	}
}

func DeleteUserHandler(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, _ := strconv.Atoi(params["id"])

		user, err := userService.DeleteUserService(r.Context(), id)

		if err == customerrors.NoUserExist {
			fmt.Println("Handler: No user of given id exist")
			w.WriteHeader(http.StatusNotFound)
			responseBytes, err1 := json.Marshal(contract.ErrorResponse{Error: "no user found"})
			if err1 != nil {
				fmt.Printf("handler: Error while converting error to json. Error: %s", err)
				return
			}
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while searching for user,error= %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, err1 := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if err1 != nil {
				fmt.Printf("handler: Error while converting error to json.Error: %s", err)
				return
			}
			w.Write(responseBytes)
			return
		}
		responsebytes, err1 := json.Marshal(user)
		if err1 != nil {
			fmt.Printf("handler: Error while converting user to json.Error: %s", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responsebytes)
	}
}
