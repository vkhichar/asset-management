package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/asset-management/domain"
)

var EmailInvalid = errors.New("Invalid email")

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

const (
	getUserByEmailQuery = "SELECT id, name, email, password, is_admin FROM users WHERE email= $1"
	selectAllUsers      = "SELECT id, name, email, password, is_admin, created_at, updated_at FROM users"
	createUserByQuery   = "INSERT INTO users (name, email, password,is_admin) VALUES ($1, $2, $3, $4) RETURNING id, name, email, password, is_admin"
)

type UserRepository interface {
	FindUser(ctx context.Context, email string) (*domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
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
	var user1 domain.User
	email := user.Email

	valid := emailRegex.MatchString(email)

	if valid == false {

		return nil, EmailInvalid

	}
	err := repo.db.Get(&user1, createUserByQuery, user.Name, user.Email, user.Password, user.IsAdmin)

	if err != nil {

		return nil, err

	}

	return &user1, nil
}
