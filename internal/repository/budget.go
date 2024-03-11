package repository

import (
	"context"
	"fmt"
	"github.com/beebeewijaya-tech/go-budget/internal/db"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
)

type BudgetRepository struct {
	db *db.Database
}

func NewBudgetRepository(db *db.Database) *BudgetRepository {
	return &BudgetRepository{
		db: db,
	}
}

func (b *BudgetRepository) Create(ctx context.Context, budget domain.Budget) (domain.Budget, error) {
	query := "INSERT INTO budgets (id, name, current, max, user_id) VALUES($1, $2, $3, $4, $5)"
	_, err := b.db.Client.ExecContext(ctx, query, budget.ID, budget.Name, budget.Current, budget.Max, budget.UserID)
	if err != nil {
		return domain.Budget{}, fmt.Errorf("error when create budget to the db %v", err)
	}

	return budget, err
}

func (b *BudgetRepository) Get(ctx context.Context, id string) (domain.Budget, error) {
	panic("create")
}

func (b *BudgetRepository) Delete(ctx context.Context, id string) error {
	panic("create")
}

func (b *BudgetRepository) List(ctx context.Context, userId string, limit, offset int) ([]domain.Budget, error) {
	var budgets []domain.Budget
	query := "SELECT id, name, current, max FROM budgets WHERE user_id = $1 LIMIT $2 OFFSET $3"

	rows, err := b.db.Client.QueryxContext(ctx, query, userId, limit, offset)
	if err != nil {
		return []domain.Budget{}, fmt.Errorf("error when querying to budgets list %v", err)
	}

	for rows.Next() {
		var budget domain.Budget
		err = rows.StructScan(&budget)
		if err != nil {
			return []domain.Budget{}, fmt.Errorf("error when extract struct from list %v", err)
		}

		budgets = append(budgets, budget)
	}

	return budgets, nil
}

func (b *BudgetRepository) Update(ctx context.Context, budget domain.Budget) (domain.Budget, error) {
	query := "UPDATE budgets SET current = current + $1 WHERE id = $2"
	_, err := b.db.Client.ExecContext(ctx, query, budget.Current, budget.ID)
	if err != nil {
		return domain.Budget{}, fmt.Errorf("error when update budget to the db %v", err)
	}

	return budget, err
}
