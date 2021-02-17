package manager

import "github.com/denchick/news/models"

// NewsService ...
type NewsService interface {
	GetNewsByTag(tag string, count int) ([]*models.Article, error)
	GetNewsByName(name string) ([]*models.Article, error)
}