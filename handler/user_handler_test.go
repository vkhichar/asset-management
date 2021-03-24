package handler_test

import (
	"bytes"
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
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/handler"
	"github.com/vkhichar/asset-management/service"
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
	id := 1
	name := "fatema"
	password := "12345"
	body := fmt.Sprintf(`{"name": "fatema", "password": "12345"}`)
	request, err := http.NewRequest("PUT", "/profile", strings.NewReader(body))
	userReq := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	claims := &service.Claims{
		UserID:  1,
		IsAdmin: false,
	}

	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}
	mockUserService.On("UpdateUser", mock.Anything, id, userReq).Return(nil, errors.New("something went wrong"))
	r := mux.NewRouter()
	r.HandleFunc("/profile", handler.UpdateUsersHandler(mockUserService)).Methods("PUT")
	context := context.WithValue(request.Context(), "claims", claims)
	r.ServeHTTP(resp, request.WithContext(context))

	//r.ServeHTTP(resp, request)
	expectedErr := string(`{"error":"something went wrong"}`)
	assert.JSONEq(t, expectedErr, resp.Body.String())
}

func TestUserHandler_UpdateUsersHandler_When_Success(t *testing.T) {
	id := 1
	name := "fatema"
	password := "12345"
	body := fmt.Sprintf(`{"name": "fatema", "password": "12345"}`)
	request, err := http.NewRequest("PUT", "/profile", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	userReq := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	timeNow := time.Now()

	expectedUserResp, _ := json.Marshal(contract.UpdateUserResponse{
		ID:        1,
		Name:      "fatema",
		Email:     "fatema.m@gmail.com",
		IsAdmin:   true,
		CreatedAt: timeNow.String(),
		UpdatedAt: timeNow.String(),
		Password:  "12345",
	})

	userResp := &domain.User{
		ID:        1,
		Name:      "fatema",
		Email:     "fatema.m@gmail.com",
		Password:  "12345",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}
	claims := &service.Claims{
		UserID:  1,
		IsAdmin: false,
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}
	mockUserService.On("UpdateUser", mock.Anything, id, userReq).Return(userResp, nil)
	r := mux.NewRouter()
	r.HandleFunc("/profile", handler.UpdateUsersHandler(mockUserService)).Methods("PUT")
	context := context.WithValue(request.Context(), "claims", claims)
	r.ServeHTTP(resp, request.WithContext(context))

	assert.JSONEq(t, string(expectedUserResp), resp.Body.String())
}

func TestUserHandler_UpdateUsersHandler_When_Nil(t *testing.T) {
	id := 1
	name := "fatema"
	password := "12345"
	body := fmt.Sprintf(`{"name": "fatema", "password": "12345"}`)
	request, err := http.NewRequest("PUT", "/profile", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	userReq := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}
	mockUserService.On("UpdateUser", mock.Anything, id, userReq).Return(nil, nil)
	r := mux.NewRouter()
	r.HandleFunc("/profile", handler.UpdateUsersHandler(mockUserService)).Methods("PUT")
	r.ServeHTTP(resp, request)
	expectedErr := string(`{"error":"User for this id does not exist"}`)

	assert.JSONEq(t, string(expectedErr), resp.Body.String())
}

func TestUserHandler_DeleteUserHandler_When_DeleteUserReturnsError(t *testing.T) {
	id := 1
	request, err := http.NewRequest("DELETE", "/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}

	mockUserService.On("DeleteUser", mock.Anything, id).Return("", errors.New("something went wrong"))
	r := mux.NewRouter()
	r.HandleFunc("/profile", handler.DeleteUserHandler(mockUserService)).Methods("DELETE")
	r.ServeHTTP(resp, request)

	expectedErr := `{"error":"something went wrong"}`

	assert.JSONEq(t, string(expectedErr), resp.Body.String())
}

func TestUserHandler_DeleteUserHandler_When_DeleteUserReturnsNil(t *testing.T) {
	id := 1
	request, err := http.NewRequest("DELETE", "/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}

	mockUserService.On("DeleteUser", mock.Anything, id).Return("", nil)
	r := mux.NewRouter()
	r.HandleFunc("/profile", handler.DeleteUserHandler(mockUserService)).Methods("DELETE")
	r.ServeHTTP(resp, request)

	expectedErr := `{"error":"no user found"}`

	assert.JSONEq(t, string(expectedErr), resp.Body.String())
}

func TestUserHandler_DeleteUserHandler_When_DeleteUserHasErrorWhileParsingId(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}
	r := mux.NewRouter()
	r.HandleFunc("/profile", handler.DeleteUserHandler(mockUserService)).Methods("DELETE")
	r.ServeHTTP(resp, request)

	expectedErr := `{"error":"Enter id in valid format"}`

	assert.JSONEq(t, string(expectedErr), resp.Body.String())
}

func TestUserHandler_DeleteUserHandler_When_Success(t *testing.T) {
	id := 1
	request, err := http.NewRequest("DELETE", "/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	result := "User successfully deleted"

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}

	mockUserService.On("DeleteUser", mock.Anything, id).Return(result, nil)
	r := mux.NewRouter()
	r.HandleFunc("/profile", handler.DeleteUserHandler(mockUserService)).Methods("DELETE")
	r.ServeHTTP(resp, request)

	expectedResult, _ := json.Marshal(result)

	assert.JSONEq(t, string(expectedResult), resp.Body.String())
}
