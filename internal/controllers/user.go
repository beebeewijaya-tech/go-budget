package controllers

import (
	"github.com/beebeewijaya-tech/go-budget/internal/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	userUsecase domain.UserUsecase
}

func NewUserController(
	userUsecase domain.UserUsecase,
) *UserController {
	return &UserController{
		userUsecase: userUsecase,
	}
}

func (u UserController) Register(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	us, err := u.userUsecase.Register(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, us)
}

func (u UserController) Login(c echo.Context) error {
	var user domain.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	us, err := u.userUsecase.Login(c.Request().Context(), user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, us)
}
