package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/config"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

func TestUserRepository_DeleteUser_When_DeleteUserReturnsError(t *testing.T) {
	ctx := context.Background()

	id := 1
	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("delete from users")

	userRepo := repository.NewUserRepository()

	user, err := userRepo.DeleteUser(ctx, id)

	assert.Nil(t, user)
	assert.Nil(t, err)
}

func TestUserRepository_DeleteUsers_When_Success(t *testing.T) {
	ctx := context.Background()
	var userExpected domain.User
	id := 1

	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("delete from users")
	tx.MustExec("insert into users (id, name,email,password,is_admin) values ($1,$2,$3,$4,$5)", id, "Jan Doe", "jandoe@gmail.com", "12345", true)
	tx.Commit()
	db.Get(&userExpected, "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users WHERE id = $1", id)

	userRepo := repository.NewUserRepository()
	user, err := userRepo.DeleteUser(ctx, id)

	assert.Equal(t, &userExpected, user)
	assert.Nil(t, err)
}
