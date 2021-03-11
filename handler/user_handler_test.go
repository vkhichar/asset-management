package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/handler"
	mockService "github.com/vkhichar/asset-management/service/mocks"
)

func TestUserHandler_When_CreateUser_ReturnError(t *testing.T) {
	ctx := context.Background()

	user := domain.User{
		ID:       0,
		Name:     "Gourav",
		Email:    "gouv@gmail.com",
		Password: "12345",
		IsAdmin:  false,
	}

	responseByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(responseByte)
	req, err := http.NewRequest("POST", "/users", requestReader)
	if err != nil {
		fmt.Printf("TestHandler: error while newRequest %s", err.Error())
		t.FailNow()
	}

	response := httptest.NewRecorder()
	expectedError := string(`{"error":"handler:something went wrong"}`)
	mockService := &mockService.MockUserService{}
	mockService.On("CreateUser", ctx, user).Return(nil, errors.New("handler:something went wrong"))

	CreateUserHandle := http.HandlerFunc(handler.CreateUserHandler(mockService))
	CreateUserHandle.ServeHTTP(response, req)

	assert.JSONEq(t, expectedError, response.Body.String())

}

func TestUserHandler_When_CreateUser_Success(t *testing.T) {

	ctx := context.Background()
	var user *domain.User

	responseByte, _ := json.Marshal(user)
	requestReader := bytes.NewReader(responseByte)
	req, err := http.NewRequest("POST", "/users", requestReader)
	if err != nil {
		fmt.Printf("TestHandler: error while newRequest %s", err.Error())
		t.FailNow()
	}

	response := httptest.NewRecorder()

	mockService := &mockService.MockUserService{}
	mockService.On("CreateUser", ctx, user).Return(user, nil)

	CreateUserHandle := http.HandlerFunc(handler.CreateUserHandler(mockService))
	CreateUserHandle.ServeHTTP(response, req)

	assert.NotNil(t, &user)
	assert.NoError(t, err)

}

func TestGetUserById_When_ReturnError_UserDoesNotExist(t *testing.T) {

	ID := 5

	req, err := http.NewRequest("GET", "/users/5", nil)
	if err != nil {
		fmt.Printf("TestHandler: error while newRequest %s", err.Error())
		t.FailNow()
	}

	response := httptest.NewRecorder()
	expectedError := string(`{"error":"user does not exist "}`)
	mockService := &mockService.MockUserService{}
	mockService.On("GetUserByID", mock.Anything, ID).Return(nil, nil)

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.GetUserByIDHandler(mockService)).Methods("GET")

	r.ServeHTTP(response, req)

	assert.JSONEq(t, expectedError, response.Body.String())

}

func TestGetUserById_When_ReturnError_SomethingWentWrong(t *testing.T) {

	ID := 5

	req, err := http.NewRequest("GET", "/users/5", nil)
	if err != nil {
		fmt.Printf("TestHandler: error while newRequest %s", err.Error())
		t.FailNow()
	}

	response := httptest.NewRecorder()
	expectedError := string(`{"error":"handler:something went wrong"}`)
	mockService := &mockService.MockUserService{}
	mockService.On("GetUserByID", mock.Anything, ID).Return(nil, errors.New("handler:something went wrong"))

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.GetUserByIDHandler(mockService)).Methods("GET")

	r.ServeHTTP(response, req)

	assert.JSONEq(t, expectedError, response.Body.String())

}

func TestUserHandler_GetUserByID_Success(t *testing.T) {

	ID := 5

	req, err := http.NewRequest("GET", "/users/5", nil)
	if err != nil {
		fmt.Printf("TestHandler: error while newRequest %s", err.Error())
		t.FailNow()
	}

	timeNow := time.Now()
	expectedObj, _ := json.Marshal(contract.GetUserByID{

		Name:      "gourav",
		Email:     "gourav@gmail.com",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	})

	user := domain.User{
		ID:        5,
		Name:      "gourav",
		Email:     "gourav@gmail.com",
		Password:  "12345",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	response := httptest.NewRecorder()

	mockService := &mockService.MockUserService{}
	mockService.On("GetUserByID", mock.Anything, ID).Return(&user, nil)

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.GetUserByIDHandler(mockService)).Methods("GET")

	r.ServeHTTP(response, req)

	assert.JSONEq(t, string(expectedObj), response.Body.String())

}
