package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

		responseBytes, statusokErr := json.Marshal(contract.LoginResponse{IsAdmin: user.IsAdmin, Token: token})
		if statusokErr != nil {
			fmt.Fprintf(w, "handler: error while marshaling status ok")
			return
		}
		w.WriteHeader(http.StatusOK)

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

		if err == customerrors.ErrInvalidEmailPassword {
			fmt.Printf("handler: invalid email or password for email: %s", req.Email)

			w.WriteHeader(http.StatusBadRequest)
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

		responseBytes, statusokErr := json.Marshal(contractUser)
		if statusokErr != nil {
			fmt.Printf("handler: error while marshaling status ok")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

		w.Write(responseBytes)

	}
}

func GetUserByIDHandler(userService service.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Context-Type", "application/json")
		params := mux.Vars(r)

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Printf("handler: invalid request for id %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, decodingErr := json.Marshal(contract.ErrorResponse{Error: "invalid request"})

			if decodingErr != nil {
				fmt.Fprintf(w, "handler: error while marshaling decoding request for get user by id")
				return
			}
			w.Write(responseBytes)
			return
		}
		user, err := userService.GetUserByID(r.Context(), id)

		if err == customerrors.UserNotExist {
			fmt.Printf("handler: user does not exist : %s", err.Error())

			w.WriteHeader(http.StatusBadRequest)
			responseBytes, invalidErr := json.Marshal(contract.ErrorResponse{Error: "user does not exist "})

			if invalidErr != nil {
				fmt.Fprintf(w, "handler: error while marshaling invalid user id")
				return
			}
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while getting user by id:error: %s", err.Error())

			w.WriteHeader(http.StatusInternalServerError)

			responseBytes, createUsererr := json.Marshal(contract.ErrorResponse{Error: "handler:something went wrong"})

			if createUsererr != nil {
				fmt.Fprintf(w, "handler: error while marshaling user id error")
				return
			}
			w.Write(responseBytes)
			return
		}

		contractUser := contract.DomaintoContractUserID(user)

		responseBytes, statusokErr := json.Marshal(contractUser)
		if statusokErr != nil {
			fmt.Printf("handler: error while marshaling status ok")
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
		w.WriteHeader(http.StatusOK)

		w.Write(responseBytes)
		return

	}

}

func UpdateUsersHandler(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, errInConversion := strconv.Atoi(params["id"])

		if errInConversion != nil {
			fmt.Printf("handler: Error while parameter conversion. Error: %s", errInConversion)
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "Error while parameter conversion"})
			w.Write(responseBytes)

			if err != nil {
				fmt.Printf("handler: Error while Marshal,%s", err)
			}
			return

		}

		var req contract.UpdateUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)

		if err != nil {
			fmt.Printf("handler:Error while decoding request for update, %s", err)
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "invalid request"})
			w.Write(responseBytes)

			if err != nil {
				fmt.Printf("handler: Error while Marshal,%s", err)
			}
			return
		}

		user, err := userService.UpdateUser(r.Context(), id, req)

		if err == customerrors.UserDoesNotExist {
			fmt.Println("handler: User for this id does not exist")
			w.WriteHeader(http.StatusNotFound)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "User for this id does not exist"})
			if err != nil {
				fmt.Printf("handler: Error while converting error to json, error:%s", err)
				return
			}
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while searching for user,error= %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			responseBytes, err := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if err != nil {
				fmt.Printf("handler: Error while converting error to json, error:%s", err)
				return
			}
			w.Write(responseBytes)
			return
		}
		resp := contract.DomainToContractUpdate(user)

		responsebytes, err := json.Marshal(resp)
		if err != nil {
			fmt.Printf("handler: Error while converting to json, error:%s", err)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responsebytes)
	}
}

func DeleteUserHandler(userService service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		params := mux.Vars(r)
		id, parseErr := strconv.Atoi(params["id"])

		if parseErr != nil {
			fmt.Println("Handler: Error while parsing Id")
			w.WriteHeader(http.StatusBadRequest)
			responseBytes, jsonErr := json.Marshal(contract.ErrorResponse{Error: "Enter id in valid format"})
			if jsonErr != nil {
				fmt.Printf("handler: Error while converting error to json. Error: %s", jsonErr)
				return
			}
			w.Write(responseBytes)
			return

		}

		user, err := userService.DeleteUser(r.Context(), id)

		if err == customerrors.NoUserExistForDelete {
			fmt.Println("Handler: No user of given id exist")
			w.WriteHeader(http.StatusNotFound)
			responseBytes, jsonErr := json.Marshal(contract.ErrorResponse{Error: "no user found"})
			if jsonErr != nil {
				fmt.Printf("handler: Error while converting error to json. Error: %s", jsonErr)
				return
			}
			w.Write(responseBytes)
			return
		}

		if err != nil {
			fmt.Printf("handler: error while searching for user,error= %s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)

			responseBytes, jsonErr := json.Marshal(contract.ErrorResponse{Error: "something went wrong"})
			if jsonErr != nil {
				fmt.Printf("handler: Error while converting error to json.Error: %s", jsonErr)
				return
			}
			w.Write(responseBytes)
			return
		}
		responsebytes, jsonErr := json.Marshal(user)
		if jsonErr != nil {
			fmt.Printf("handler: Error while converting user to json.Error: %s", jsonErr)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(responsebytes)
	}
}
