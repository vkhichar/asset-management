package service

import (
	"context"
	"fmt"

	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"

	"github.com/vkhichar/asset-management/domain"
	"github.com/vkhichar/asset-management/repository"
)

type UserService interface {
	Login(ctx context.Context, email, password string) (user *domain.User, token string, err error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	UpdateUser(ctx context.Context, id int, req contract.UpdateUserRequest) (user *domain.User, err error)
	DeleteUser(ctx context.Context, id int) (*domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	tokenSvc TokenService
	eventSvc EventService
}

func NewUserService(repo repository.UserRepository, ts TokenService, es EventService) UserService {
	return &userService{
		userRepo: repo,
		tokenSvc: ts,
		eventSvc: es,
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
	//create user service
	return nil, nil
}

func (service *userService) UpdateUser(ctx context.Context, id int, req contract.UpdateUserRequest) (*domain.User, error) {
	user, err := service.userRepo.UpdateUser(ctx, id, req)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return user, customerrors.UserDoesNotExist
	}

	eventId, errEvent := service.eventSvc.PostUserEvent(ctx, user)
	if errEvent != nil {
		fmt.Printf("Service: Error while creating event. Error: %s", errEvent.Error())
		return user, errEvent
	} else {
		fmt.Println("New event created:", eventId)
	}

	return user, nil
}

func (service *userService) DeleteUser(ctx context.Context, id int) (*domain.User, error) {
	user, err := service.userRepo.DeleteUser(ctx, id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return user, customerrors.NoUserExistForDelete
	}

	return user, nil
}
