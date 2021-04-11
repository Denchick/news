package store

import (
	"github.com/denchick/news/models"
)

// NewsRepository is a store for News
//go:generate mockery --dir . --name NewsRepository --output ./mocks
type NewsRepository interface {
	Create(*models.DBArticle) error
	GetByName(name string) ([]*models.DBArticle, error)
	GetFromFeed(feed *models.DBFeed) ([]*models.DBArticle, error)
	GetSimilar(name string) ([]*models.DBArticle, error)
}

// FeedsRepository is a store for Feeds
//go:generate mockery --dir . --name FeedsRepository --output ./mocks
type FeedsRepository interface {
	GetFeeds() ([]*models.DBFeed, error)
	GetSubcategoryFeeds(subcategory *models.DBSubcategory) ([]*models.DBFeed, error)
}

type CategoriesRepository interface {
	GetCategory(name string) (*models.DBCategory, error)
}

type SubcategoriesRepository interface {
	GetSubcategories(categoryID uint) ([]*models.DBSubcategory, error)
}