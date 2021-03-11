package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

const (
	getUserByEmailQuery = "SELECT id, name, email, password, is_admin FROM users WHERE email= $1"
	selectAllUsers      = "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users"
	getUserById         = "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users WHERE id=$1"
	deleteUserById      = "DELETE FROM users WHERE id=$1"
)

type UserRepository interface {
	FindUser(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	DeleteUser(ctx context.Context, id int) (*domain.User, error)
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

func (repo *userRepo) DeleteUser(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User

	err := repo.db.Get(&user, getUserById, id)

	if err == sql.ErrNoRows {
		fmt.Printf("Repository: No users present")
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	tx := repo.db.MustBegin()

	tx.MustExec(deleteUserById, id)

	defer func() {
		errOnCommit := tx.Commit()

		if errOnCommit != nil {
			fmt.Printf("Repo: Error while commiting in table. Error: %s", errOnCommit)
			tx.Rollback()
		}
	}()
	return &user, nil
}
