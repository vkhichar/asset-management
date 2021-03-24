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
	GetUserByID(ctx context.Context, ID int) (*domain.User, error)
	UpdateUser(ctx context.Context, id int, req contract.UpdateUserRequest) (user *domain.User, err error)
	DeleteUser(ctx context.Context, id int) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
	tokenSvc TokenService
	eventSvc EventService
}

func NewUserService(repo repository.UserRepository, ts TokenService, event EventService) UserService {
	return &userService{
		userRepo: repo,
		tokenSvc: ts,
		eventSvc: event,
	}
}

func (service *userService) Login(ctx context.Context, email string, password string) (*domain.User, string, error) {
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

	id, err := service.eventSvc.PostCreateUserEvent(ctx, entry)
	if err != nil {
		fmt.Printf("user service: error while calling postuserevent: %s", err.Error())
	} else {
		fmt.Println("New event created:", id)
	}
	return entry, nil

}

func (service *userService) GetUserByID(ctx context.Context, ID int) (*domain.User, error) {
	user, err := service.userRepo.GetUserByID(ctx, ID)

	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, customerrors.UserNotExist
	}

	return user, nil
}

func (service *userService) UpdateUser(ctx context.Context, id int, req contract.UpdateUserRequest) (*domain.User, error) {
	user, err := service.userRepo.UpdateUser(ctx, id, req)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return user, customerrors.UserDoesNotExist
	}

	eventId, errEvent := service.eventSvc.PostUpdateUserEvent(ctx, user)
	if errEvent != nil {
		fmt.Printf("Service: Error while creating event. Error: %s", errEvent.Error())
		return user, nil
	} else {
		fmt.Println("New event created:", eventId)
	}

	return user, nil
}

func (service *userService) DeleteUser(ctx context.Context, id int) (string, error) {
	result, err := service.userRepo.DeleteUser(ctx, id)

	if err != nil {
		return "", err
	}

	if result == "" {
		return "", customerrors.UserDoesNotExist
	}

	return result, nil
}
