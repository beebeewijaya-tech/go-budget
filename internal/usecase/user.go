package usecase

import (
	"context"
	"fmt"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/beebeewijaya-tech/go-budget/internal/utilities"
	"github.com/google/uuid"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(
	repo domain.UserRepository,
) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (s *UserUsecase) Register(ctx context.Context, user domain.User) (domain.User, error) {
	var err error
	user.ID = uuid.New().String()
	user.Password, err = utilities.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to hash password: %v", err)
	}

	u, err := s.repo.Create(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("error when create user on the repo: %v", err)
	}

	u.Password = ""

	return u, nil
}

func (s *UserUsecase) Login(ctx context.Context, user domain.User) (domain.User, error) {
	u, err := s.repo.Get(ctx, user.Email)
	if err != nil {
		return domain.User{}, fmt.Errorf("error when create user on the repo: %v", err)
	}

	checkPassword := utilities.CheckPasswordHash(user.Password, u.Password)
	if !checkPassword {
		return domain.User{}, fmt.Errorf("error when comparing user password: %v", err)
	}

	token := utilities.GenerateToken(u)
	if token == "" {
		return domain.User{}, fmt.Errorf("error when generating token: %v", err)
	}

	u.Password = ""
	u.Token = token

	return u, nil
}
