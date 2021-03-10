package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/handler"
	mockService "github.com/vkhichar/asset-management/service/mocks"
)

func TestUserHandler_ListUsersHandler_When_ListUsersReturnsError(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	expectedErr := string(`{"error":"something went wrong"}`)

	mockUserService := &mockService.MockUserService{}
	mockUserService.On("ListUsers", ctx).Return(nil, errors.New("something went wrong"))

	handlerTest := http.HandlerFunc(handler.ListUsersHandler(mockUserService))
	handlerTest.ServeHTTP(rr, req)

	assert.JSONEq(t, expectedErr, rr.Body.String())
}

func TestUserHandler_ListUsersHandler_When_Success(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequest("GET", "/users", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	timeNow := time.Now()
	users := []domain.User{
		{
			ID:        1,
			Name:      "Jan Doe",
			Email:     "jandoe@gmail.com",
			Password:  "12345",
			IsAdmin:   true,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		{
			ID:        2,
			Name:      "Alisa Ray",
			Email:     "alisaray@gmail.com",
			Password:  "hello",
			IsAdmin:   false,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
	}

	expectedUsers, _ := json.Marshal([]contract.User{
		{
			ID:        1,
			Name:      "Jan Doe",
			Email:     "jandoe@gmail.com",
			IsAdmin:   true,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
		{
			ID:        2,
			Name:      "Alisa Ray",
			Email:     "alisaray@gmail.com",
			IsAdmin:   false,
			CreatedAt: timeNow,
			UpdatedAt: timeNow,
		},
	})

	mockUserService := &mockService.MockUserService{}
	mockUserService.On("ListUsers", ctx).Return(users, nil)

	handleTest := handler.ListUsersHandler(mockUserService)
	handleTest.ServeHTTP(rr, req)

	assert.JSONEq(t, string(expectedUsers), rr.Body.String())
}

func TestUserHandler_ListUsersHandler_When_ListUsersReturnsNil(t *testing.T) {
	ctx := context.Background()
	req, err := http.NewRequest("GET", "/users", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}
	mockUserService.On("ListUsers", ctx).Return(nil, nil)

	handleTest := handler.ListUsersHandler(mockUserService)
	handleTest.ServeHTTP(rr, req)

	expectedErr := `{"error" : "no user found"}`

	assert.JSONEq(t, string(expectedErr), rr.Body.String())
}

func TestUserHandler_UpdateUsersHandler_When_UpdateUsersReturnsError(t *testing.T) {
	// ctx := context.Background()

	id := 1
	name := "fatema"
	password := "12345"

	// reqBody, _ := json.Marshal(map[string]string{
	// 	"name":     name,
	// 	"password": password,
	// })

	body := fmt.Sprintf(`{"name": "fatema", "password": "12345"}`)

	//ioReader(reqBody)

	//fmt.Println(string(reqBody))

	request, err := http.NewRequest("PUT", "/users/1", strings.NewReader(body))

	userReq := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}
	mockUserService.On("UpdateUser", request.Context(), id, userReq).Return(nil, errors.New("something went wrong"))

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.UpdateUsersHandler(mockUserService)).Methods("PUT")
	r.ServeHTTP(resp, request)

	// handlerTest := http.HandlerFunc(handler.UpdateUsersHandler(mockUserService))
	// handlerTest.ServeHTTP(resp, request)

	expectedErr := string(`{"error":"something went wrong"}`)

	assert.JSONEq(t, expectedErr, resp.Body.String())
}
