package service

import (
	"context"
	"errors"

	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

var ErrInvalidEmailPassword = errors.New("invalid email or password")

type UserService interface {
	Login(ctx context.Context, email, password string) (user *domain.User, token string, err error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) (users []domain.User, err error)
}

type userService struct {
	userRepo repository.UserRepository
	tokenSvc TokenService
}

func NewUserService(repo repository.UserRepository, ts TokenService) UserService {
	return &userService{
		userRepo: repo,
		tokenSvc: ts,
	}
}

func (service *userService) Login(ctx context.Context, email, password string) (*domain.User, string, error) {
	user, err := service.userRepo.FindUser(ctx, email)
	if err != nil {
		return nil, "", err
	}

	if user == nil {
		return nil, "", ErrInvalidEmailPassword
	}

	if user.Password != password {
		return nil, "", ErrInvalidEmailPassword
	}

	claims := &Claims{UserID: user.ID, IsAdmin: user.IsAdmin}
	token, err := service.tokenSvc.GenerateToken(claims)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (service *userService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return nil, nil
}

func (service *userService) ListUsers(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}
