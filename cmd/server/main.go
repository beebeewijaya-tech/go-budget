package main

import (
	"fmt"
	"github.com/beebeewijaya-tech/go-budget/internal/controllers"
	"github.com/beebeewijaya-tech/go-budget/internal/db"
	"github.com/beebeewijaya-tech/go-budget/internal/delivery/http"
	"github.com/beebeewijaya-tech/go-budget/internal/repository"
	"github.com/beebeewijaya-tech/go-budget/internal/usecase"
	"github.com/spf13/viper"
	"log"
)

func Run() error {
	config := viper.New()
	config.AddConfigPath("env")
	config.SetConfigType("json")
	config.SetConfigName("config")
	err := config.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %w", err)
	}

	database := db.NewDatabase(config)

	// repository
	userRepo := repository.NewUserRepository(database)
	budgetRepo := repository.NewBudgetRepository(database)
	expenseRepo := repository.NewExpenseRepository(database)

	// usecase
	userUsecase := usecase.NewUserUsecase(userRepo)
	budgetUsecase := usecase.NewBudgetUsecase(budgetRepo)
	expenseUsecase := usecase.NewExpenseUsecase(expenseRepo, budgetRepo)

	// controllers
	userController := controllers.NewUserController(userUsecase)
	budgetController := controllers.NewBudgetController(budgetUsecase)
	expenseController := controllers.NewExpenseController(expenseUsecase)

	// routers
	routers := http.NewServer(userController, budgetController, expenseController)

	return routers.Run(":9000")
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
