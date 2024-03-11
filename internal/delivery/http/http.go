package http

import (
	"github.com/beebeewijaya-tech/go-budget/internal/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Router            *echo.Echo
	userController    *controllers.UserController
	budgetController  *controllers.BudgetController
	expenseController *controllers.ExpenseController
}

func NewServer(
	userController *controllers.UserController,
	budgetController *controllers.BudgetController,
	expenseController *controllers.ExpenseController,
) *Server {
	s := &Server{
		userController:    userController,
		budgetController:  budgetController,
		expenseController: expenseController,
	}

	s.Router = echo.New()
	s.registerMiddleware()
	s.registerAuthRoutes()
	s.registerBusinessRoutes()

	return s
}

func (s Server) registerMiddleware() {
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
}

func (s Server) registerAuthRoutes() {
	s.Router.POST("/v1/login", s.userController.Login)
	s.Router.POST("/v1/register", s.userController.Register)
}

func (s Server) registerBusinessRoutes() {
	pvt := s.Router.Group("")
	pvt.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("and0LXRlc3QK"),
	}))
	pvt.POST("/v1/budget", s.budgetController.CreateBudget)
	pvt.GET("/v1/budget", s.budgetController.ListBudget)

	pvt.POST("/v1/expense", s.expenseController.CreateExpense)
	pvt.GET("/v1/expense", s.expenseController.ListExpense)
}

func (s Server) Run(address string) error {
	return s.Router.Start(address)
}
