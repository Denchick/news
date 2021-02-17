package repositories

import (
	"errors"

	"github.com/denchick/news/models"
	"github.com/go-pg/pg/v10"
)

// NewsRepository ...
type NewsRepository struct {
	db *pg.DB
}

// NewNewsRepository ...
func NewNewsRepository(db *pg.DB) *NewsRepository {
	return &NewsRepository{db}
}

// FindByName ...
func (repo *NewsRepository) FindByName(name string) ([]*models.Article, error) {
	return nil, errors.New("Not implemented")
}

// FindSimilar ...
func (repo *NewsRepository) FindSimilar(name string) ([]*models.Article, error) {
	return nil, errors.New("Not implemented")
}
