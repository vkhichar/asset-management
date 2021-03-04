package service

import (
	"context"

	"github.com/vkhichar/asset-management/errorspkg"

	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type UserService interface {
	Login(ctx context.Context, email, password string) (user *domain.User, token string, err error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
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
		return nil, "", errorspkg.ErrInvalidEmailPassword
	}

	if user.Password != password {
		return nil, "", errorspkg.ErrInvalidEmailPassword
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
		return user, errorspkg.NoUsersExist
	}

	return user, nil
}

func (service *userService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	//create user service
	return nil, nil
}
