package usecase

import (
	"context"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/google/uuid"
)

type BudgetUsecase struct {
	budgetRepo domain.BudgetRepository
}

func NewBudgetUsecase(
	budgetRepo domain.BudgetRepository,
) *BudgetUsecase {
	return &BudgetUsecase{
		budgetRepo: budgetRepo,
	}
}

func (b *BudgetUsecase) Create(ctx context.Context, budget domain.Budget) (domain.Budget, error) {
	budget.ID = uuid.New().String()
	bu, err := b.budgetRepo.Create(ctx, budget)
	if err != nil {
		return domain.Budget{}, err
	}

	return bu, err
}

func (b *BudgetUsecase) GetByID(ctx context.Context, id string) (domain.Budget, error) {
	panic("create")
}

func (b *BudgetUsecase) DeleteByID(ctx context.Context, id string) error {
	panic("create")
}

func (b *BudgetUsecase) List(ctx context.Context, userId string, limit, page int) ([]domain.Budget, error) {
	offset := (page - 1) * limit
	bu, err := b.budgetRepo.List(ctx, userId, limit, offset)
	if err != nil {
		return []domain.Budget{}, err
	}

	return bu, err
}
