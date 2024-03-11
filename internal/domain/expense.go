package domain

import "context"

// Expense - is responsible to be a sharing entities between package
// The purpose for this is to standardize Expense object
type Expense struct {
	ID          string  `json:"id" params:"id"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Budget      string  `json:"budget" db:"budget_id"`
	UserID      string  `json:"user_id" db:"user_id""`
}

// ExpenseRequest - is responsible to be a request object for binding
type ExpenseRequest struct {
	ID    string `params:"id"`
	Limit int    `query:"limit"`
	Page  int    `query:"page"`
}

// ExpenseRepository - is responsible to be a blueprint that will explain the whole Expense Service
// That storing into the data store
//
//go:generate mockgen -package=mock_repo -destination mock_repo/expense.go . ExpenseRepository
type ExpenseRepository interface {
	Create(ctx context.Context, expense Expense) (Expense, error)
	Get(ctx context.Context, id string) (Expense, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, userId string, limit, offset int) ([]Expense, error)
}

// ExpenseUsecase - will responsible to be a blueprint that will handle business logic of the data
// Responsible for Expense
//
//go:generate mockgen -package=mock_usecase -destination mock_usecase/expense.go . ExpenseUsecase
type ExpenseUsecase interface {
	Create(ctx context.Context, expense Expense) (Expense, error)
	GetByID(ctx context.Context, id string) (Expense, error)
	DeleteByID(ctx context.Context, id string) error
	List(ctx context.Context, userId string, limit, page int) ([]Expense, error)
}
