package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

func TestUserRepo_CreateUser_When_CreateUserReturnSuccess(t *testing.T) {

	ctx := context.Background()
	user := domain.User{
		ID:       13,
		Name:     "Nikhil",
		Email:    "nikhil@email",
		Password: "1234",
		IsAdmin:  true,
	}
	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	var newUser domain.User
	userRepo := repository.NewUserRepository()

	returnuser, err := userRepo.CreateUser(ctx, user)
	if err != nil {
		fmt.Printf("userRepo Mock:error while inserting data %s", err.Error())
		return
	}

	db.Get(&newUser, "SELECT * FROM users WHERE id= $1", 13)

	assert.Equal(t, &newUser, returnuser)
	assert.NoError(t, err)
}

func TestUserRepo_GetUserByID_When_GetUserByID_ReturnUserExist(t *testing.T) {

	ctx := context.Background()

	ID := 5
	config.Init()
	repository.InitDB()
	userRepo := repository.NewUserRepository()

	returnUser, err := userRepo.GetUserByID(ctx, ID)
	assert.NotEmpty(t, returnUser)
	assert.NoError(t, err)
}
