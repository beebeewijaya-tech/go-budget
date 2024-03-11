package controllers

import (
	"fmt"
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/beebeewijaya-tech/go-budget/internal/utilities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BudgetController struct {
	budgetUsecase domain.BudgetUsecase
}

func NewBudgetController(
	budgetUsecase domain.BudgetUsecase,
) *BudgetController {
	return &BudgetController{
		budgetUsecase: budgetUsecase,
	}
}

func (b *BudgetController) CreateBudget(c echo.Context) error {
	var budget domain.Budget
	if err := c.Bind(&budget); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when parsing budget request %v", err))
	}

	user := utilities.GetUser(c)
	budget.UserID = user.ID

	bu, err := b.budgetUsecase.Create(c.Request().Context(), budget)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, bu)
}

func (b *BudgetController) ListBudget(c echo.Context) error {
	var budgetReq domain.BudgetRequest
	if err := c.Bind(&budgetReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Errorf("error when parsing budget request %v", err))
	}

	fmt.Println("budgetReq: ", budgetReq)
	user := utilities.GetUser(c)

	bu, err := b.budgetUsecase.List(c.Request().Context(), user.ID, budgetReq.Limit, budgetReq.Page)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, bu)
}
