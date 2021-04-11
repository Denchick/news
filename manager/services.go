package manager

import "github.com/denchick/news/models"

// NewsService ...
type NewsService interface {
	SaveNews(articles []*models.DBArticle) error
	GetNews(category string) ([]*models.NewsSubcategory, error)
	GetNewsByName(name string) ([]*models.DBArticle, error)
}