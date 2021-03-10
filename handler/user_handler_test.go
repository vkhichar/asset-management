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

	"github.com/stretchr/testify/assert"
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
	mockService.On("CreateUser", ctx, user).Return(user, err)

	CreateUserHandle := http.HandlerFunc(handler.CreateUserHandler(mockService))
	CreateUserHandle.ServeHTTP(response, req)

	fmt.Printf("%s\n", response.Body.String())
	assert.NotNil(t, &user)
	assert.NoError(t, err)

}

func TestUserHandler_GetUserById_When_ReturnError(t *testing.T) {

	ctx := context.Background()
	ID := "14"
	responseByte, _ := json.Marshal(ID)
	requestReader := bytes.NewReader(responseByte)
	req, err := http.NewRequest("GET", "/users/{ID}", requestReader)
	if err != nil {
		fmt.Printf("TestHandler: error while newRequest %s", err.Error())
		t.FailNow()
	}

	response := httptest.NewRecorder()
	expectedError := string(`{"error":"invalid request"}`)
	mockService := &mockService.MockUserService{}
	mockService.On("GetUserByID", ctx, ID).Return(nil, errors.New("invalid request"))

	CreateUserHandle := http.HandlerFunc(handler.CreateUserHandler(mockService))
	CreateUserHandle.ServeHTTP(response, req)

	fmt.Println(response.Body.String())

	assert.JSONEq(t, expectedError, response.Body.String())

}

func TestUserHandler_GetUserByID_Success(t *testing.T) {

	ctx := context.Background()
	var user *domain.User
	ID := "1"
	responseByte, _ := json.Marshal(ID)
	requestReader := bytes.NewReader(responseByte)
	req, err := http.NewRequest("GET", "/users/{ID}", requestReader)
	if err != nil {
		fmt.Printf("TestHandler: error while newRequest %s", err.Error())
		t.FailNow()
	}

	response := httptest.NewRecorder()

	mockService := &mockService.MockUserService{}
	mockService.On("GetUserByID", ctx, ID).Return(nil, err)

	CreateUserHandle := http.HandlerFunc(handler.CreateUserHandler(mockService))
	CreateUserHandle.ServeHTTP(response, req)

	assert.NotNil(t, &user)
	assert.NoError(t, err)

}
