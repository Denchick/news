package manager

import "github.com/denchick/news/models"

// NewsService ...
type NewsService interface {
	SaveNews(articles []*models.Article) error
	GetNewsByTag(tag string, count int) ([]*models.Article, error)
	GetNewsByName(name string) ([]*models.Article, error)
}