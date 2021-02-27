package repositories

import (
	"github.com/denchick/news/models"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

// FeedsRepository ...
type FeedsRepository struct {
	db *pg.DB
}

// NewFeedsRepository ...
func NewFeedsRepository(db *pg.DB) *FeedsRepository {
	return &FeedsRepository{db}
}

// GetFeeds ...
func (repo *FeedsRepository) GetFeeds() ([]*models.Feed, error) {
	var feeds []*models.Feed
	err := repo.db.Model(&feeds).Select()
	
	if err != nil { 
		return nil, errors.Wrap(err, "store.repository.GetFeeds")
	}
	return feeds, nil
}