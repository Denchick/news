package store

import (
	"github.com/denchick/news/models"
)

// NewsRepository is a store for News
//go:generate mockery --dir . --name NewsRepository --output ./mocks
type NewsRepository interface {
	BulkCreate([]*models.Article) error
	GetByName(name string) ([]*models.Article, error)
	GetFromFeed(feed *models.Feed) ([]*models.Article, error)
	GetSimilar(name string) ([]*models.Article, error)
}

// FeedsRepository is a store for Feeds
//go:generate mockery --dir . --name FeedsRepository --output ./mocks
type FeedsRepository interface {
	GetFeeds() ([]*models.Feed, error)
	GetFeedsFromCategory(category *models.Category) ([]*models.Feed, error)
}

type CategoriesRepository interface {
	GetCategory(name string) (*models.Category, error)
}