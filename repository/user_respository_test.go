package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/vkhichar/asset-management/domain"
	mockRepo "github.com/vkhichar/asset-management/repository/mocks"
	mockService "github.com/vkhichar/asset-management/service/mocks"
)

func (repo *userRepo) TestUserRepo_CreateUser_When_CreateUserReturnError(t *testing.T) {

	ctx := context.Background()
	user := domain.User{
		ID:       1,
		Name:     "Dummy",
		Email:    "dummy@email",
		Password: "12345",
		IsAdmin:  false,
	}
	newUser := repo.db.Get("INSERT INTO users (name, email, password,is_admin) VALUES ($1, $2, $3, $4) RETURNING id, name, email, password, is_admin, created_at, updated_at", user.Name, user.Email, user.Password, user.IsAdmin)

	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenService := &mockService.MockTokenService{}
	mockUserRepo.On("CreateUser", ctx, user).Return(nil, errors.New("invalid request"))

}
