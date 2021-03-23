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
		ID:       104,
		Name:     "Nikhil",
		Email:    "nikhil@email",
		Password: "1234",
		IsAdmin:  true,
	}
	config.Init()
	repository.InitDB()

	userRepo := repository.NewUserRepository()

	returnuser, err := userRepo.CreateUser(ctx, user)
	if err != nil {
		fmt.Printf("userRepo Mock:error while inserting data %s", err.Error())
		return
	}

	assert.NotNil(t, returnuser)
	assert.NoError(t, err)
}

func TestUserRepo_GetUserByID_When_GetUserByID_ReturnUserExist(t *testing.T) {

	ctx := context.Background()

	ID := 12

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	userRepo := repository.NewUserRepository()
	tx.MustExec("DELETE FROM users WHERE id=$1", "1")
	tx.MustExec("INSERT INTO users (id, name, email, password,is_admin) VALUES ($1, $2, $3, $4, $5)", "1", "gourav", "gourav@gmail.com", "12345", true)
	tx.Commit()
	returnUser, err := userRepo.GetUserByID(ctx, ID)
	assert.NotEmpty(t, returnUser)
	assert.NoError(t, err)
}

func TestUserRepo_GetUserByID_When_GetUserByID_ReturnUserNotExist(t *testing.T) {
	ctx := context.Background()

	ID := 2
	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	userRepo := repository.NewUserRepository()
	tx := db.MustBegin()
	tx.MustExec("DELETE FROM users WHERE id=$1", "2")
	tx.Commit()
	returnUser, err := userRepo.GetUserByID(ctx, ID)
	fmt.Println(err)
	assert.Empty(t, returnUser)
	assert.Nil(t, err)
}
