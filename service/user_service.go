package service

import (
	"context"
	"errors"

	"github.com/vkhichar/asset-management/customerrors"

	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

var ErrInvalidEmailPassword = errors.New("invalid email or password")

type UserService interface {
	Login(ctx context.Context, email, password string) (user *domain.User, token string, err error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
}

var ExtraError = errors.New("invalid email")

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
		return nil, "", customerrors.ErrInvalidEmailPassword
	}

	if user.Password != password {
		return nil, "", customerrors.ErrInvalidEmailPassword
	}

	claims := &Claims{UserID: user.ID, IsAdmin: user.IsAdmin}
	token, err := service.tokenSvc.GenerateToken(claims)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (service *userService) ListUsers(ctx context.Context) ([]domain.User, error) {
	user, err := service.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return user, customerrors.NoUsersExist
	}

	return user, nil
}

func (service *userService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {

	entry, err := service.userRepo.CreateUser(ctx, user)

	if err != nil {
		return nil, err
	}

	return entry, nil

}
