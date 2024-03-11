package repository

import (
	"context"
	"fmt"
	"github.com/beebeewijaya-tech/go-budget/internal/db"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
)

type UserRepository struct {
	Database *db.Database
}

func NewUserRepository(db *db.Database) *UserRepository {
	return &UserRepository{
		Database: db,
	}
}

func (u *UserRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	query := "INSERT INTO users (id, email, password) VALUES($1, $2, $3)"

	_, err := u.Database.Client.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("error when inserting to users %v", err)
	}

	return user, nil
}

func (u *UserRepository) Get(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	query := "SELECT id, email, password FROM users WHERE email=$1"

	err := u.Database.Client.QueryRowxContext(ctx, query, email).StructScan(&user)
	if err != nil {
		return domain.User{}, fmt.Errorf("error when getting user %v", err)
	}

	return user, nil
}
