package services

import (
	"errors"

	"github.com/denchick/news/models"
	"github.com/denchick/news/store"
)

// NewsService ...
type NewsService struct {
	store *store.Store
}

// NewNewsService ...
func NewNewsService(store *store.Store) *NewsService {
	return &NewsService{store}
}

// GetNewsByTag ...
func (service *NewsService) GetNewsByTag(tag string, count int) ([]*models.Article, error) {
	return nil, errors.New("Not implemented")
}

// GetNewsByName ...
func (service *NewsService) GetNewsByName(name string) ([]*models.Article, error) {
	return nil, errors.New("Not implemented")
}
