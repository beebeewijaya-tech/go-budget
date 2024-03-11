package repository

import (
	"context"
	"fmt"
	"github.com/beebeewijaya-tech/go-budget/internal/db"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
)

type ExpenseRepository struct {
	db *db.Database
}

func NewExpenseRepository(db *db.Database) *ExpenseRepository {
	return &ExpenseRepository{
		db: db,
	}
}

func (e *ExpenseRepository) Create(ctx context.Context, expense domain.Expense) (domain.Expense, error) {
	query := "INSERT INTO expenses (id, name, description, amount, budget_id, user_id) VALUES($1, $2, $3, $4, $5, $6)"
	_, err := e.db.Client.ExecContext(ctx, query, expense.ID, expense.Name, expense.Description, expense.Amount, expense.Budget, expense.UserID)
	if err != nil {
		return domain.Expense{}, fmt.Errorf("error when create expense to the db %v", err)
	}

	return expense, err
}

func (e *ExpenseRepository) Get(ctx context.Context, id string) (domain.Expense, error) {
	panic("unimplemented")
}

func (e *ExpenseRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func (e *ExpenseRepository) List(ctx context.Context, userId string, limit, offset int) ([]domain.Expense, error) {
	var expenses []domain.Expense
	query := "SELECT id, name, description, amount, budget_id, user_id FROM expenses WHERE user_id = $1 LIMIT $2 OFFSET $3"

	rows, err := e.db.Client.QueryxContext(ctx, query, userId, limit, offset)
	if err != nil {
		return []domain.Expense{}, fmt.Errorf("error when querying to expenses list %v", err)
	}

	for rows.Next() {
		var expense domain.Expense
		err = rows.StructScan(&expense)
		if err != nil {
			return []domain.Expense{}, fmt.Errorf("error when extract struct from list %v", err)
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
