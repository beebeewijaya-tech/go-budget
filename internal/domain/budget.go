package domain

import "context"

// Budget - is responsible to be a sharing entities between package
// The purpose for this is to standardize Budget object
type Budget struct {
	ID      string  `json:"id" params:"id"`
	Name    string  `json:"name"`
	Current float64 `json:"current,omitempty"`
	Max     float64 `json:"max,omitempty"`
	UserID  string  `json:"user_id,omitempty"`
}

// BudgetRequest - is responsible to be a query params for budgets when getting the list
type BudgetRequest struct {
	ID    string `params:"id"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
}

// BudgetRepository - is responsible to be a blueprint that will explain the whole Budget Service
// That storing into the data store
//
//go:generate mockgen -package=mock_repo -destination mock_repo/budget.go . BudgetRepository
type BudgetRepository interface {
	Create(ctx context.Context, budget Budget) (Budget, error)
	Get(ctx context.Context, id string) (Budget, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, userId string, limit, offset int) ([]Budget, error)
	Update(ctx context.Context, budget Budget) (Budget, error)
}

// BudgetUsecase - will responsible to be a blueprint that will handle business logic of the data
// Responsible for budget
//
//go:generate mockgen -package=mock_usecase -destination mock_usecase/budget.go . BudgetUsecase
type BudgetUsecase interface {
	Create(ctx context.Context, budget Budget) (Budget, error)
	GetByID(ctx context.Context, id string) (Budget, error)
	DeleteByID(ctx context.Context, id string) error
	List(ctx context.Context, userId string, limit, page int) ([]Budget, error)
}
