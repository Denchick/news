package store

import (
	"github.com/denchick/news/models"
)

// ArticlesRepository is a store for articles
//go:generate mockery --dir . --name ArticlesRepository --output ./mocks
type ArticlesRepository interface {
	Create(*models.DBArticle) error
	GetByName(name string) ([]*models.DBArticle, error)
	GetFromFeed(feed *models.DBFeed) ([]*models.DBArticle, error)
	GetSimilar(name string) ([]*models.DBArticle, error)
}

// FeedsRepository is a store for feeds
//go:generate mockery --dir . --name FeedsRepository --output ./mocks
type FeedsRepository interface {
	GetFeeds() ([]*models.DBFeed, error)
	GetSubcategoryFeeds(subcategory *models.DBSubcategory) ([]*models.DBFeed, error)
}

// CategoriesRepository is a store for categories
//go:generate mockery --dir . --name CategoriesRepository --output ./mock
type CategoriesRepository interface {
	GetCategory(name string) (*models.DBCategory, error)
}

// SubcategoriesRepository is a store for subcategories
//go:generate mockery --dir . --name SubcategoriesRepository --output ./mock
type SubcategoriesRepository interface {
	GetSubcategories(categoryID uint) ([]*models.DBSubcategory, error)
}
