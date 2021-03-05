package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/contract"
	"github.com/vkhichar/asset-management/customerrors"
	"github.com/vkhichar/asset-management/domain"
)

const (
	getUserByEmailQuery = "SELECT id, name, email, password, is_admin FROM users WHERE email= $1"
	selectAllUsers      = "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users"
	getUserByIDQuery    = "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users WHERE id = $1"
	updateUserColumns   = "UPDATE users SET name = $1, password = $2, updated_at = $3 WHERE id = $4"
)

type UserRepository interface {
	FindUser(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	UpdateUserRepo(ctx context.Context, id int, req contract.UpdateUserRequest) (*domain.User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepository() UserRepository {
	return &userRepo{
		db: GetDB(),
	}
}

func (repo *userRepo) FindUser(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	err := repo.db.Get(&user, getUserByEmailQuery, email)
	if err == sql.ErrNoRows {
		fmt.Printf("repository: couldn't find user for email: %s", email)

		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (repo *userRepo) ListUsers(ctx context.Context) ([]domain.User, error) {
	var user []domain.User
	err := repo.db.Select(&user, selectAllUsers)

	if err == sql.ErrNoRows {
		fmt.Printf("repository: No users present")

		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *userRepo) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	//create user method
	return nil, nil
}

func (repo *userRepo) UpdateUserRepo(ctx context.Context, id int, req contract.UpdateUserRequest) (*domain.User, error) {
	var user domain.User
	var tempUser domain.User
	err := repo.db.Get(&tempUser, getUserByIDQuery, id)

	if err != nil {
		return nil, customerrors.UserDoesNotExist
	}

	if req.Name == nil {
		req.Name = &tempUser.Name
	}

	if req.Password == nil {
		req.Password = &tempUser.Password
	}

	tx := repo.db.MustBegin()

	tx.MustExec(updateUserColumns, *req.Name, *req.Password, time.Now(), id)

	err1 := tx.Commit()
	if err1 != nil {
		fmt.Printf("repo: Error while commiting the database. Error: %s", err1)
		tx.Rollback()
	}
	err = repo.db.Get(&user, getUserByIDQuery, id)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
