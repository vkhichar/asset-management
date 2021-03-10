package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

func TestUserRepository_ListUsers_When_Success(t *testing.T) {
	ctx := context.Background()
	var userExpected []domain.User

	db := repository.GetDB()
	tx := db.MustBegin()
	tx.MustExec("insert into users (name,email,password,is_admin) values ($1,$2,$3,$3)", "Jan Doe", "jandoe@gmail.com", "12345", true)
	tx.MustExec("insert into users (name,email,password,is_admin) values ($1,$2,$3,$4)", "Alisa Ray", "alisaray@gmail.com", "hello", false)

	db.Select(&userExpected, "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users")
	userRepo := repository.NewUserRepository()

	user, _ := userRepo.ListUsers(ctx)

	assert.Equal(t, userExpected, user)
}
