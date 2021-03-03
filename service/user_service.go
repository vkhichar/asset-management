package service

import (
	"context"
	"errors"

	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

var ErrInvalidEmailPassword = errors.New("invalid email or password")

var NoUsersExist = errors.New("No users exist at present")

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

<<<<<<< HEAD
func (service *userService) ListUser(ctx context.Context) ([]domain.User, error) {
	user, err := service.userRepo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, NoUsersExist
	}

	return user, nil
=======
func (service *userService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	return nil, nil
}

func (service *userService) ListUsers(ctx context.Context) ([]domain.User, error) {
	return nil, nil
>>>>>>> a5591a3f4381c2750a8ae4cf1cd8709e2d7c0b87
}
