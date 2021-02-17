package store

import (
	"errors"

	"github.com/denchick/news/config"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // required
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Usefull migrate commands
// $ migrate create -ext sql -dir migrations *name*
// $ migrate -path migrations -database postgres://postgres:postgres@0.0.0.0/db_news?sslmode=disable force 20210207232148

// runPgMigrations runs Postgres migrations
func runPgMigrations() error {
	cfg := config.Get()
	if cfg.PgMigrationsPath == "" {
		return nil
	}
	if cfg.PgURL == "" {
		return errors.New("No cfg.PgURL provided")
	}
	m, err := migrate.New(
		cfg.PgMigrationsPath,
		cfg.PgURL,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}