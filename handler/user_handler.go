package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vkhichar/asset-management/customerrors"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
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
			responseBytes, decodingErr := json.Marshal(contract.ErrorResponse{Error: "invalid request"})

			if decodingErr != nil {
				fmt.Fprintf(w, "handler: error while marshaling decoding request for login")
				return
			}
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: invalid request for email: %s", req.Email)
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, validationErr := json.Marshal(contract.ErrorResponse{Error: err.Error()})

			if validationErr != nil {
				fmt.Fprintf(w, "handler: error while marshaling invalid request for email")
				return
			}
			w.Write(responseBytes)
			return
		}

		user, token, err := userService.Login(r.Context(), req.Email, req.Password)
		if err == customerrors.ErrInvalidEmailPassword {
			fmt.Printf("handler: invalid email or password for email: %s", req.Email)
			w.WriteHeader(http.StatusUnauthorized)
			responseBytes, invalidErr := json.Marshal(contract.ErrorResponse{Error: "invalid email or password"})
			if invalidErr != nil {
				fmt.Fprint(w, "handler: error while marshaling invalid email or password for email")
				return
			}
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while logging in for email: %s, error: %s", req.Email, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, loginErr := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if loginErr != nil {
				fmt.Fprintf(w, "handler: error while marshaling logging in for email")
				return
			}
			w.Write(responseBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		responseBytes, statusokErr := json.Marshal(contract.LoginResponse{IsAdmin: user.IsAdmin, Token: token})
		if statusokErr != nil {
			fmt.Fprintf(w, "handler: error while marshaling status ok")
			return
		}
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

		responsebytes, _ := json.Marshal(userResp)
		w.WriteHeader(http.StatusOK)
		w.Write(responsebytes)
	}

}

func CreateUserHandler(userService service.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Context-Type", "application/json")

		var req contract.CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {

			fmt.Printf("handler: error while decoding request for create user: %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, decodingErr := json.Marshal(contract.ErrorResponse{Error: "invalid request"})

			if decodingErr != nil {
				fmt.Fprintf(w, "handler: error while marshaling decoding request for create user")
				return
			}
			w.Write(responseBytes)
			return
		}

		err = req.Validate()
		if err != nil {
			fmt.Printf("handler: invalid request for email: %s", req.Email)

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, validationErr := json.Marshal(contract.ErrorResponse{Error: err.Error()})

			if validationErr != nil {
				fmt.Fprintf(w, "handler: error while marshaling invalid request for email")
				return
			}
			w.Write(responseBytes)
			return
		}

		createdUser := domain.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			IsAdmin:  req.IsAdmin,
		}

		user, err := userService.CreateUser(r.Context(), createdUser)

		if err == service.ErrInvalidEmailPassword {
			fmt.Printf("handler: invalid email or password for email: %s", req.Email)

			w.WriteHeader(http.StatusUnauthorized)
			responseBytes, invalidErr := json.Marshal(contract.ErrorResponse{Error: "invalid email or password"})

			if invalidErr != nil {
				fmt.Fprintf(w, "handler: error while marshaling invalid email or password for email")
				return
			}
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while creating user for email: %s, error: %s", req.Email, err.Error())

			w.WriteHeader(http.StatusInternalServerError)

			responseBytes, createUsererr := json.Marshal(contract.ErrorResponse{Error: "handler:something went wrong"})

			if createUsererr != nil {
				fmt.Fprintf(w, "handler: error while marshaling create user for email error")
				return
			}
			w.Write(responseBytes)
			return
		}

		contractUser := contract.DomaintoContract(user)

		w.WriteHeader(http.StatusOK)
		responseBytes, statusokErr := json.Marshal(contractUser)
		if statusokErr != nil {
			fmt.Fprintf(w, "handler: error while marshaling status ok")
			return
		}
		w.Write(responseBytes)

	}
}
