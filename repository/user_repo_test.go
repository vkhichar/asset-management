package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

func TestUserRepository_ListUsers_When_Nil(t *testing.T) {
	ctx := context.Background()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("delete from users")

	userRepo := repository.NewUserRepository()

	user, err := userRepo.ListUsers(ctx)
	assert.Nil(t, err)
	assert.Nil(t, user)
}

func TestUserRepository_ListUsers_When_Success(t *testing.T) {
	ctx := context.Background()
	var userExpected []domain.User

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("delete from users")
	tx.MustExec("insert into users (name,email,password,is_admin) values ($1,$2,$3,$4)", "Jan Doe", "jandoe@gmail.com", "12345", true)
	tx.MustExec("insert into users (name,email,password,is_admin) values ($1,$2,$3,$4)", "Alisa Ray", "alisaray@gmail.com", "hello", false)

	db.Select(&userExpected, "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users")
	userRepo := repository.NewUserRepository()

	user, err := userRepo.ListUsers(ctx)
	assert.Nil(t, err)
	assert.Equal(t, userExpected, user)
}

func TestUserRepository_UpdateUsers_When_Nil(t *testing.T) {
	ctx := context.Background()

	name := "fatema"
	password := "12345"

	userReq := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	id := 3

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("delete from users")
	tx.MustExec("insert into users (name,email,password,is_admin) values ($1,$2,$3,$4)", "Jan Doe", "jandoe@gmail.com", "12345", true)

	userRepo := repository.NewUserRepository()

	user, err := userRepo.UpdateUser(ctx, id, userReq)
	expectedErr := "The user for this id does not exist"
	assert.Equal(t, expectedErr, err.Error())
	assert.Nil(t, user)
}

func TestUserRepository_UpdateUsers_When_Success(t *testing.T) {
	ctx := context.Background()
	var userExpected domain.User

	name := "fatema"
	password := "hello"

	userReq := contract.UpdateUserRequest{
		Name:     &name,
		Password: &password,
	}

	id := 12

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("delete from users")
	tx.MustExec("insert into users (id, name,email,password,is_admin) values ($1,$2,$3,$4,$5)", 12, "Jan Doe", "jandoe@gmail.com", "12345", true)
	tx.Commit()

	userRepo := repository.NewUserRepository()

	user, err := userRepo.UpdateUser(ctx, id, userReq)

	db.Get(&userExpected, "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users WHERE id = $1", id)

	assert.Equal(t, &userExpected, user)
	assert.Nil(t, err)
}
