package domain

import "context"

// User - is responsible to be a sharing entities between package
// The purpose for this is to standardize User object
type User struct {
	ID       string `json:"id,omitempty" params:"id"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

// UserRepository - is responsible to be a blueprint that will explain the whole User Service
// That storing into the data store
//
//go:generate mockgen -package=mock_repo -destination mock_repo/user.go . UserRepository
type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	Get(ctx context.Context, email string) (User, error)
}

// UserUsecase - will responsible to be a blueprint that will handle business logic of the data
// Responsible for User
//
//go:generate mockgen -package=mock_usecase -destination mock_usecase/user.go . UserUsecase
type UserUsecase interface {
	Register(ctx context.Context, user User) (User, error)
	Login(ctx context.Context, user User) (User, error)
}
