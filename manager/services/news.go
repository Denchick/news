package services

import (
	"github.com/pkg/errors"
	"github.com/denchick/news/models"
	"github.com/denchick/news/store"
)

type NewsService struct {
	store *store.Store
}

func NewNewsService(store *store.Store) *NewsService {
	return &NewsService{store}
}

func (service *NewsService) SaveNews(articles []*models.DBArticle) error {
	for _, article := range articles {	
		err := service.store.News.Create(article)
		if err != nil {
			return errors.Wrap(err, "manager.services.SaveNews")
		}
	}
	return nil
}

func (service *NewsService) GetNews(categoryName string) ([]*models.ArticleGroups, error) {
	category, err := service.store.Categories.GetCategory(categoryName)
	if err != nil {
		return nil, errors.Wrap(err, "manager.services.GetNews")
	}	

	feeds, err := service.store.Feeds.GetFeedsFromCategory(category)
	if err != nil {
		return nil, errors.Wrap(err, "manager.services.GetNews")
	}	

	articleGroups := make([]*models.ArticleGroups, 0)
	for _, feed := range feeds {	
		articles, err := service.store.News.GetFromFeed(feed)
		if err != nil {
			return nil, errors.Wrap(err, "manager.services.GetNews")
		}
		articleGroups = append(articleGroups, &models.ArticleGroups{Link: feed.URL, FeedName: feed.Name, Articles: articles})
	}

	return articleGroups, nil
}

func (service *NewsService) GetNewsByName(name string) ([]*models.DBArticle, error) {
	return service.store.News.GetSimilar(name)
}
