package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/domain"
	mockRepo "github.com/vkhichar/asset-management/repository/mocks"
	"github.com/vkhichar/asset-management/service"
	mockService "github.com/vkhichar/asset-management/service/mocks"
)

func TestUserService_Login_When_FindUserReturnsError(t *testing.T) {
	ctx := context.Background()
	email := "dummy@email"

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}

	mockUserRepo.On("FindUser", ctx, email).Return(nil, errors.New("some db error"))

	userService := service.NewUserService(mockUserRepo, mockTokenService)
	user, token, err := userService.Login(ctx, email, "1234")

	// fmt.Printf("%s\n", err.Error())
	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Equal(t, "", token)
	assert.Nil(t, user)
}

func TestUserService_Login_Success(t *testing.T) {
	ctx := context.Background()
	email := "dummy@email"
	inputPassword := "12345"
	user := domain.User{
		ID:       1,
		Name:     "Dummy",
		Email:    "dummy@email",
		Password: "12345",
		IsAdmin:  true,
	}

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}

	mockUserRepo.On("FindUser", ctx, email).Return(&user, nil)
	claims := &service.Claims{UserID: 1, IsAdmin: true}
	mockTokenService.On("GenerateToken", claims).Return("generated-token", nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService)
	dbUser, token, err := userService.Login(ctx, email, inputPassword)

	assert.NoError(t, err)
	assert.Equal(t, "generated-token", token)
	assert.Equal(t, &user, dbUser)
}

func TestUserService_CreatUser_CreateUserReturnsError(t *testing.T) {
	ctx := context.Background()

	user := domain.User{
		ID:       1,
		Name:     "Dummy",
		Email:    "dummy@gmail.com",
		Password: "12345",
		IsAdmin:  true,
	}

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}

	mockUserRepo.On("CreateUser", ctx, user).Return(nil, errors.New("some db error"))
	userService := service.NewUserService(mockUserRepo, mockTokenService)
	newUser, err := userService.CreateUser(ctx, user)

	// fmt.Printf("%s\n", err.Error())
	if err == nil {
		fmt.Printf("Error while creating user")
		t.FailNow()
	}

	assert.Error(t, err)
	assert.Equal(t, "some db error", err.Error())
	assert.Nil(t, newUser)

}

func TestUserService_CreateUser_Success(t *testing.T) {
	ctx := context.Background()

	user := domain.User{
		ID:       1,
		Name:     "Dummy",
		Email:    "dummy@email",
		Password: "12345",
		IsAdmin:  false,
	}

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}

	mockUserRepo.On("CreateUser", ctx, user).Return(&user, nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService)
	// dbUser, token, err := userService.Login(ctx, email, inputPassword)
	dbUser, err := userService.CreateUser(ctx, user)

	assert.NoError(t, err)
	assert.Equal(t, &user, dbUser)
}
