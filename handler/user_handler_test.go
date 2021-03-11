package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/handler"
	mockService "github.com/vkhichar/asset-management/service/mocks"
)

func TestUserHandler_DeleteUserHandler_When_DeleteUserReturnsError(t *testing.T) {
	id := 1
	request, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}

	mockUserService.On("DeleteUser", mock.Anything, id).Return(nil, errors.New("something went wrong"))
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.DeleteUserHandler(mockUserService)).Methods("DELETE")
	r.ServeHTTP(resp, request)

	expectedErr := `{"error":"something went wrong"}`

	assert.JSONEq(t, string(expectedErr), resp.Body.String())
}

func TestUserHandler_DeleteUserHandler_When_DeleteUserReturnsNil(t *testing.T) {
	id := 1
	request, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}

	mockUserService.On("DeleteUser", mock.Anything, id).Return(nil, nil)
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.DeleteUserHandler(mockUserService)).Methods("DELETE")
	r.ServeHTTP(resp, request)

	expectedErr := `{"error":"no user found"}`

	assert.JSONEq(t, string(expectedErr), resp.Body.String())
}

func TestUserHandler_DeleteUserHandler_When_DeleteUserHasErrorWhileParsingId(t *testing.T) {
	request, err := http.NewRequest("DELETE", "/users/AB", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.DeleteUserHandler(mockUserService)).Methods("DELETE")
	r.ServeHTTP(resp, request)

	expectedErr := `{"error":"Enter id in valid format"}`

	assert.JSONEq(t, string(expectedErr), resp.Body.String())
}

func TestUserHandler_DeleteUserHandler_When_Success(t *testing.T) {
	id := 1
	request, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	timeNow := time.Now()

	user := domain.User{
		ID:        1,
		Name:      "Dummy",
		Email:     "dummy@email",
		Password:  "12345",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	resp := httptest.NewRecorder()
	mockUserService := &mockService.MockUserService{}

	mockUserService.On("DeleteUser", mock.Anything, id).Return(&user, nil)
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", handler.DeleteUserHandler(mockUserService)).Methods("DELETE")
	r.ServeHTTP(resp, request)

	expectedUser, _ := json.Marshal(user)

	assert.JSONEq(t, string(expectedUser), resp.Body.String())
}
