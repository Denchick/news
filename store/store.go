package store

import (
	"log"

	"github.com/denchick/news/internal/db"
	"github.com/denchick/news/store/repositories"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

// Store contains all repositories
type Store struct {
	DB   *pg.DB
	News NewsRepository
}

// NewStore creates new store
func NewStore() (*Store, error) {
	pgDB, err := db.Dial()
	if err != nil {
		return nil, errors.Wrap(err, "store.Dial")
	}

	log.Println("Running PostgreSQL migrations...")
	if err := db.RunPgMigrations(); err != nil {
		return nil, errors.Wrap(err, "store.runPgMigrations")
	}

	store := &Store{pgDB, repositories.NewNewsRepository(pgDB)}

	return store, nil
}

