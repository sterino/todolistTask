package store

import (
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func Migrate(db SQLX) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db.Client.DB, "migrations/postgres"); err != nil {
		return err
	}

	return nil
}
