package usecase

import (
	"context"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/google/uuid"
)

type ExpenseUsecase struct {
	expenseRepo domain.ExpenseRepository
	budgetRepo  domain.BudgetRepository
}

func NewExpenseUsecase(expenseRepo domain.ExpenseRepository, budgetRepo domain.BudgetRepository) *ExpenseUsecase {
	return &ExpenseUsecase{
		expenseRepo: expenseRepo,
		budgetRepo:  budgetRepo,
	}
}

func (e *ExpenseUsecase) Create(ctx context.Context, expense domain.Expense) (domain.Expense, error) {
	expense.ID = uuid.New().String()
	ex, err := e.expenseRepo.Create(ctx, expense)
	if err != nil {
		return domain.Expense{}, err
	}

	budget := domain.Budget{
		ID:      expense.Budget,
		Current: expense.Amount,
	}
	_, err = e.budgetRepo.Update(ctx, budget)
	if err != nil {
		return domain.Expense{}, err
	}

	return ex, err
}

func (e *ExpenseUsecase) GetByID(ctx context.Context, id string) (domain.Expense, error) {
	panic("unimplemented")
}

func (e *ExpenseUsecase) DeleteByID(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (e *ExpenseUsecase) List(ctx context.Context, userId string, limit, page int) ([]domain.Expense, error) {
	offset := (page - 1) * limit
	ex, err := e.expenseRepo.List(ctx, userId, limit, offset)
	if err != nil {
		return []domain.Expense{}, err
	}

	return ex, err
}
