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
	mockevent := &mockService.MockEventService{}
	mockUserRepo.On("FindUser", ctx, email).Return(nil, errors.New("some db error"))

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)
	user, token, err := userService.Login(ctx, email, "1234")

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
	mockevent := &mockService.MockEventService{}

	mockUserRepo.On("FindUser", ctx, email).Return(&user, nil)
	claims := &service.Claims{UserID: 1, IsAdmin: true}
	mockTokenService.On("GenerateToken", claims).Return("generated-token", nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)
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
	mockevent := &mockService.MockEventService{}

	mockUserRepo.On("CreateUser", ctx, user).Return(nil, errors.New("some db error"))
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)
	newUser, err := userService.CreateUser(ctx, user)

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
	mockevent := &mockService.MockEventService{}

	mockUserRepo.On("CreateUser", ctx, user).Return(&user, nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)
	dbUser, err := userService.CreateUser(ctx, user)

	assert.NoError(t, err)
	assert.Equal(t, &user, dbUser)
}

func TestUserService_GetUserById_When_ReturnError_User_not_exist(t *testing.T) {
	ctx := context.Background()
	ID := 5
	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockUserRepo.On("GetUserByID", ctx, ID).Return(nil, nil)
	mockevent := &mockService.MockEventService{}

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)
	newUser, err := userService.GetUserByID(ctx, ID)

	if err == nil {
		fmt.Printf("Error while creating user")
		t.FailNow()
	}

	assert.Equal(t, "User does not exist", err.Error())
	assert.Error(t, err)
	assert.Nil(t, newUser)

}

func TestUserService_GetUserById_When_ReturnError(t *testing.T) {
	ctx := context.Background()
	ID := 5
	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockUserRepo.On("GetUserByID", ctx, ID).Return(nil, errors.New("invalid request"))
	mockevent := &mockService.MockEventService{}

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)
	newUser, err := userService.GetUserByID(ctx, ID)

	if err == nil {
		fmt.Printf("Error while creating user")
		t.FailNow()
	}

	assert.Equal(t, "invalid request", err.Error())
	assert.Error(t, err)
	assert.Nil(t, newUser)

}
func TestUserService_GetUserByID_Success(t *testing.T) {
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
	mockUserRepo.On("GetUserByID", ctx, user.ID).Return(&user, nil)
	mockevent := &mockService.MockEventService{}

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)
	dbUser, err := userService.GetUserByID(ctx, user.ID)

	assert.NoError(t, err)
	assert.Equal(t, &user, dbUser)
}

func TestEventService_PostUserEvent_ReturnsSuccess(t *testing.T) {
	ctx := context.Background()

	user := domain.User{
		ID:       1,
		Name:     "Dummy",
		Email:    "dummy@email",
		Password: "12345",
		IsAdmin:  false,
	}
	event_id := "122"
	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockUserRepo.On("CreateUser", ctx, user).Return(&user, nil)
	mockevent := &mockService.MockEventService{}
	mockevent.On("PostUserEvent", ctx, &user).Return(event_id, nil)
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)

	obj, err := userService.CreateUser(ctx, user)

	assert.NotNil(t, obj)
	assert.NoError(t, err)

}

func TestEventService_PostUserEvent_ReturnsError(t *testing.T) {
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
	mockevent := &mockService.MockEventService{}
	mockevent.On("PostUserEvent", ctx, &user).Return("", errors.New("some error in event service"))
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)

	obj, err := userService.CreateUser(ctx, user)

	assert.NotNil(t, obj)
	assert.Error(t, err)

}
