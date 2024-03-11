package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase(config *viper.Viper) *Database {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		config.GetString("database.host"),
		config.GetInt32("database.port"),
		config.GetString("database.name"),
		config.GetString("database.username"),
		config.GetString("database.password"),
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("error when connecting to database %v", err)
	}

	return &Database{
		Client: db,
	}
}
