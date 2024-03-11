package controllers

import (
	"fmt"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/beebeewijaya-tech/go-budget/internal/utilities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ExpenseController struct {
	expenseUsecase domain.ExpenseUsecase
}

func NewExpenseController(expenseUsecase domain.ExpenseUsecase) *ExpenseController {
	return &ExpenseController{
		expenseUsecase: expenseUsecase,
	}
}

func (e ExpenseController) CreateExpense(c echo.Context) error {
	var expense domain.Expense
	if err := c.Bind(&expense); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when binding request %v", err).Error())
	}

	user := utilities.GetUser(c)
	expense.UserID = user.ID

	ex, err := e.expenseUsecase.Create(c.Request().Context(), expense)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, ex)
}

func (e ExpenseController) ListExpense(c echo.Context) error {
	var expense domain.ExpenseRequest
	if err := c.Bind(&expense); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when binding request %v", err).Error())
	}

	user := utilities.GetUser(c)

	ex, err := e.expenseUsecase.List(c.Request().Context(), user.ID, expense.Limit, expense.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, ex)
}
