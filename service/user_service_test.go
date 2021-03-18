package service_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
	mockRepo "github.com/vkhichar/asset-management/repository/mocks"
	"github.com/vkhichar/asset-management/service"
	mockService "github.com/vkhichar/asset-management/service/mocks"
	"gopkg.in/h2non/gock.v1"
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

	mockevent := &mockService.MockEventService{}
	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}

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

func TestUserService_ListUsers_When_ListUsersReturnsError(t *testing.T) {
	ctx := context.Background()

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}

	mockUserRepo.On("ListUsers", ctx).Return(nil, errors.New("Some db error"))

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	user, err := userService.ListUsers(ctx)

	assert.Error(t, err)
	assert.Equal(t, "Some db error", err.Error())
	assert.Nil(t, user)
}

func TestUserService_ListUsers_When_Success(t *testing.T) {
	ctx := context.Background()
	users := []domain.User{
		{
			ID:       1,
			Name:     "Jan Doe",
			Email:    "jandoe@gmail.com",
			Password: "12345",
			IsAdmin:  true,
		},
		{
			ID:       2,
			Name:     "Alisa Ray",
			Email:    "alisaray@gmail.com",
			Password: "hello",
			IsAdmin:  false,
		},
		{
			ID:       3,
			Name:     "Tom Walters",
			Email:    "tomwalters@gmail.com",
			Password: "tom123",
			IsAdmin:  false,
		},
	}

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}

	mockUserRepo.On("ListUsers", ctx).Return(users, nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	usersDb, err := userService.ListUsers(ctx)

	assert.NoError(t, err)
	assert.Equal(t, users, usersDb)
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

	event_id := "122"
	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockUserRepo.On("CreateUser", ctx, user).Return(&user, nil)
	mockevent := &mockService.MockEventService{}
	mockevent.On("PostCreateUserEvent", ctx, &user).Return(event_id, nil)
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)

	obj, err := userService.CreateUser(ctx, user)

	assert.NotNil(t, obj)
	assert.NoError(t, err)

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
		fmt.Printf("Error while getting info of user")
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
		fmt.Printf("Error while getting info of user")
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

func TestEventService_PostCreateUserEvent_ReturnsSuccess(t *testing.T) {
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
	mockevent.On("PostCreateUserEvent", ctx, &user).Return(event_id, nil)
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)

	obj, err := userService.CreateUser(ctx, user)

	assert.Equal(t, &user, obj)
	assert.NotNil(t, obj)
	assert.NoError(t, err)

}

func TestEventService_PostCreateUserEvent_ReturnsError(t *testing.T) {
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
	mockevent.On("PostCreateUserEvent", ctx, &user).Return("", errors.New("some error in event service"))
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockevent)

	obj, err := userService.CreateUser(ctx, user)

	assert.NotNil(t, obj)
	assert.Error(t, err)
	assert.EqualError(t, err, "some error in event service")
}

func TestUserService_ListUsers_When_ListUsersReturnsNil(t *testing.T) {
	ctx := context.Background()

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}

	mockUserRepo.On("ListUsers", ctx).Return(nil, nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	user, err := userService.ListUsers(ctx)

	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUserService_UpdateUser_When_UpdateUserReturnsError(t *testing.T) {
	ctx := context.Background()

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}

	id := 1
	name := "Fatema Moaiyadi"
	password := "hello123"
	timeNow := time.Now()

	req := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	user := &domain.User{
		ID:        1,
		Name:      "Fatema Moaiyadi",
		Email:     "jandoe@gmail.com",
		Password:  "hello123",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	mockUserRepo.On("UpdateUser", ctx, id, req).Return(nil, errors.New("User of given id does not exist"))

	mockEventService.On("PostUpdateUserEvent", ctx, user).Return("", nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	user, err := userService.UpdateUser(ctx, id, req)

	assert.Error(t, err)
	assert.Equal(t, "User of given id does not exist", err.Error())
	assert.Nil(t, user)
}

func TestUserService_UpdateUser_When_Success(t *testing.T) {
	ctx := context.Background()

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}

	id := 1
	name := "Fatema Moaiyadi"
	password := "hello123"
	timeNow := time.Now()

	req := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	user := &domain.User{
		ID:        1,
		Name:      "Fatema Moaiyadi",
		Email:     "jandoe@gmail.com",
		Password:  "hello123",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	mockUserRepo.On("UpdateUser", ctx, id, req).Return(user, nil)
	mockEventService.On("PostUpdateUserEvent", ctx, user).Return("", nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	userFromDb, err := userService.UpdateUser(ctx, id, req)

	assert.Nil(t, err)
	assert.Equal(t, user, userFromDb)
}

func TestUserService_UpdateUser_When_UpdateUserReturnsNil(t *testing.T) {
	ctx := context.Background()

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}

	id := 4
	name := "Fatema Moaiyadi"
	password := "hello123"
	timeNow := time.Now()

	req := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	user := &domain.User{
		ID:        1,
		Name:      "Fatema Moaiyadi",
		Email:     "jandoe@gmail.com",
		Password:  "hello123",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	mockUserRepo.On("UpdateUser", ctx, id, req).Return(nil, nil)
	mockEventService.On("PostUpdateUserEvent", ctx, user).Return("", nil)

	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	userFromDb, err := userService.UpdateUser(ctx, id, req)

	assert.Nil(t, userFromDb)
	assert.NotNil(t, err)
}

func TestUserService_UpdateUser_When_PostUserEventReturnsError(t *testing.T) {
	ctx := context.Background()

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	timeNow := time.Now()

	user := &domain.User{
		ID:        1,
		Name:      "Fatema Moaiyadi",
		Email:     "jandoe@gmail.com",
		Password:  "hello123",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	id := 4
	name := "Fatema Moaiyadi"
	password := "hello123"

	req := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	mockUserRepo.On("UpdateUser", ctx, id, req).Return(user, nil)
	mockEventService.On("PostUpdateUserEvent", ctx, user).Return("", errors.New("Error while creating event"))

	userFromService, err := userService.UpdateUser(ctx, id, req)

	assert.NotNil(t, userFromService)
	assert.Nil(t, err)
}

func TestUserService_UpdateUser_When_PostUserEventReturnsId(t *testing.T) {
	ctx := context.Background()

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	timeNow := time.Now()

	user := &domain.User{
		ID:        1,
		Name:      "Fatema Moaiyadi",
		Email:     "jandoe@gmail.com",
		Password:  "hello123",
		IsAdmin:   true,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	id := 4
	eventId := "22"
	name := "Fatema Moaiyadi"
	password := "hello123"

	req := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	mockUserRepo.On("UpdateUser", ctx, id, req).Return(user, nil)
	mockEventService.On("PostUpdateUserEvent", ctx, user).Return(eventId, nil)

	userFromService, err := userService.UpdateUser(ctx, id, req)

	assert.NotNil(t, userFromService)
	assert.Nil(t, err)
	assert.Equal(t, userFromService, user)
}

func TestUserService_DeleteUser_When_DeleteUserReturnsError(t *testing.T) {
	ctx := context.Background()
	id := 4

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	mockUserRepo.On("DeleteUser", ctx, id).Return(nil, errors.New("Some Internal Server Error"))

	dbUser, err := userService.DeleteUser(ctx, id)

	expectedErr := "Some Internal Server Error"

	assert.NotNil(t, err)
	assert.Nil(t, dbUser)
	assert.Equal(t, expectedErr, err.Error())
}

func TestUserService_DeleteUser_When_DeleteUserReturnsNil(t *testing.T) {
	ctx := context.Background()
	id := 4

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	mockUserRepo.On("DeleteUser", ctx, id).Return(nil, nil)

	dbUser, err := userService.DeleteUser(ctx, id)

	expectedErr := "The user for this id does not exist"

	assert.Equal(t, expectedErr, err.Error())
	assert.Nil(t, dbUser)
}

func TestUserService_DeleteUser_When_Success(t *testing.T) {
	ctx := context.Background()
	id := 1

	user := domain.User{
		ID:       1,
		Name:     "Dummy",
		Email:    "dummy@email",
		Password: "12345",
		IsAdmin:  true,
	}

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockEventService := &mockService.MockEventService{}
	userService := service.NewUserService(mockUserRepo, mockTokenService, mockEventService)

	mockUserRepo.On("DeleteUser", ctx, id).Return(&user, nil)

	dbUser, err := userService.DeleteUser(ctx, id)

	assert.NoError(t, err)
	assert.Equal(t, &user, dbUser)
}

func TestUserService_PostUserEvent_When_HTTPostReturnsSuccess(t *testing.T) {
	ctx := context.Background()
	defer gock.Off()

	service.InitEnv()

	gock.New(config.GetEventServiceUrl()).Post(service.EventResource).
		Reply(200).JSON(map[string]int{"id": 21})

	user := domain.User{
		ID:       1,
		Name:     "Dummy",
		Email:    "dummy@email",
		Password: "12345",
		IsAdmin:  true,
	}

	eventService := service.NewEventService()
	eventId, err := eventService.PostUpdateUserEvent(ctx, &user)

	assert.Nil(t, err)
	assert.Equal(t, "21", eventId)
}

func TestUserService_PostUserEvent_When_HTTPostReturnsError(t *testing.T) {
	ctx := context.Background()

	service.InitEnv()

	gock.New(config.GetEventServiceUrl()).Post("/events").
		Reply(400)

	user := domain.User{
		ID:       1,
		Name:     "Dummy",
		Email:    "dummy@email",
		Password: "12345",
		IsAdmin:  true,
	}

	eventService := service.NewEventService()
	eventId, err := eventService.PostUpdateUserEvent(ctx, &user)

	assert.NotNil(t, err)
	assert.Equal(t, "", eventId)
}
