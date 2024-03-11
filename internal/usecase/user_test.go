package usecase

import (
	"context"
	"errors"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/beebeewijaya-tech/go-budget/internal/domain/mock_repo"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

type testCase struct {
	name    string
	payload domain.User
	repo    func(r *mock_repo.MockUserRepository)
	error   bool
}

func TestUserUsecase_Register(t *testing.T) {
	testCases := []testCase{
		{
			name: "ok",
			payload: domain.User{
				ID:       uuid.New().String(),
				Email:    "abc@domain.com",
				Password: "abcdefg",
			},
			repo: func(r *mock_repo.MockUserRepository) {
				r.EXPECT().Create(gomock.Any(), gomock.Any()).
					Return(domain.User{
						ID:       uuid.New().String(),
						Email:    "abc@domain.com",
						Password: "abcdefg",
					}, nil)
			},
			error: false,
		},
		{
			name: "error",
			payload: domain.User{
				ID:       uuid.New().String(),
				Email:    "",
				Password: "",
			},
			repo: func(r *mock_repo.MockUserRepository) {
				r.EXPECT().Create(gomock.Any(), gomock.Any()).
					Return(domain.User{}, errors.New("rejected by database"))
			},
			error: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := mock_repo.NewMockUserRepository(ctrl)

			tc.repo(repo)

			u := NewUserUsecase(repo)
			require.NotNil(t, u)

			user, err := u.Register(context.Background(), tc.payload)
			if tc.error {
				require.Error(t, err, "rejected by database")
				require.NotEmpty(t, err)
				require.Empty(t, user)
			} else {
				require.NotEmpty(t, user)
				require.Empty(t, err)
			}
		})
	}
}

func TestUserUsecase_Login(t *testing.T) {
	testCases := []testCase{
		{
			name: "ok",
			payload: domain.User{
				ID:       uuid.New().String(),
				Email:    "abc@domain.com",
				Password: "abcdefg",
			},
			repo: func(r *mock_repo.MockUserRepository) {
				r.EXPECT().Get(gomock.Any(), gomock.Any()).
					Return(domain.User{
						Email:    "abc@domain.com",
						Password: "$2a$14$.KLxQGiMmNTLgIspflpzvOZueyUUSQLOjtB7C8KrzgWowrqJauO1O",
					}, nil)
			},
			error: false,
		},
		{
			name: "error when get user",
			payload: domain.User{
				Email:    "ddf@gmail.com",
				Password: "sdffddf",
			},
			repo: func(r *mock_repo.MockUserRepository) {
				r.EXPECT().Get(gomock.Any(), gomock.Any()).
					Return(domain.User{}, errors.New("user not found"))
			},
			error: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			repo := mock_repo.NewMockUserRepository(ctrl)

			tc.repo(repo)

			u := NewUserUsecase(repo)
			require.NotNil(t, u)

			user, err := u.Login(context.Background(), tc.payload)
			if tc.error {
				require.Error(t, err, "user not found")
				require.NotEmpty(t, err)
				require.Empty(t, user)
			} else {
				require.NotEmpty(t, user)
				require.NotEmpty(t, user.Token)
				require.Empty(t, err)
			}
		})
	}
}
