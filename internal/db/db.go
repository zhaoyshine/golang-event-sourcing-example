package db

import (
	"fmt"
	"os"

	"eventsourcing/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func MustInit(cfg *config.Config) {
	dbInfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
		cfg.DB.User, cfg.DB.Database)

	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		os.Exit(1)
	}

	DB = db
}
