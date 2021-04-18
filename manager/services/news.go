package services

import (
	"github.com/denchick/news/models"
	"github.com/denchick/news/store"
	"github.com/pkg/errors"
)

// NewsService...
type NewsService struct {
	store *store.Store
}

// NewNewsService creates new NewsService
func NewNewsService(store *store.Store) *NewsService {
	return &NewsService{store}
}

// SaveNews...
func (service *NewsService) SaveNews(articles []*models.DBArticle) error {
	for _, article := range articles {
		err := service.store.Articles.Create(article)
		if err != nil {
			return errors.Wrap(err, "manager.services.SaveNews")
		}
	}
	return nil
}

// GetNews...
func (service *NewsService) GetNews(categoryName string) ([]*models.NewsSubcategory, error) {
	category, err := service.store.Categories.GetCategory(categoryName)
	if err != nil {
		return nil, errors.Wrap(err, "manager.services.GetNews")
	}

	subcategories, err := service.store.Subcategories.GetSubcategories(category.ID)
	if err != nil {
		return nil, errors.Wrap(err, "manager.services.GetNews")
	}

	newsSubcategory := make([]*models.NewsSubcategory, 0)
	for _, subcategory := range subcategories {

		feeds, err := service.store.Feeds.GetSubcategoryFeeds(subcategory)
		if err != nil {
			return nil, errors.Wrap(err, "manager.services.GetNews")
		}

		newsColumns := make([]*models.NewsColumn, 0)
		for _, feed := range feeds {
			articles, err := service.store.Articles.GetFromFeed(feed)
			if err != nil {
				return nil, errors.Wrap(err, "manager.services.GetNews")
			}
			newsColumns = append(newsColumns, &models.NewsColumn{URL: feed.URL, FeedName: feed.Name, Articles: service.toWeb(articles)})
		}
		newsSubcategory = append(newsSubcategory, &models.NewsSubcategory{Name: subcategory.Name, Columns: newsColumns})
	}

	return newsSubcategory, nil
}

// GetNewsByName...
func (service *NewsService) GetNewsByName(name string) ([]*models.DBArticle, error) {
	return service.store.Articles.GetSimilar(name)
}

func (service *NewsService) toWeb(oldArticles []*models.DBArticle) []*models.WebArticle {
	var newArticles []*models.WebArticle
	for _, article := range oldArticles {
		newArticles = append(newArticles, article.ToWebArticle())
	}
	return newArticles
}
