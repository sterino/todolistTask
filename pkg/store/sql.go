package store

import (
	_ "database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
	"log"
	"todoList/internal/config"
)

type SQLX struct {
	Client *sqlx.DB
}

func NewSQL(cfg config.Config) (store SQLX, err error) {
	psqlUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	log.Printf("url: %v", psqlUrl)

	store.Client, err = sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return
	}
	log.Printf("Connected to database with URL: %s", psqlUrl)

	return

}
