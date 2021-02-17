package store

import (
	"github.com/denchick/news/models"
)

// NewsRepository is a store for News
//go:generate mockery --dir . --name NewsRepository --output ./mocks
type NewsRepository interface {
	FindByName(name string) ([]*models.Article, error)
	FindSimilar(name string) ([]*models.Article, error)
}