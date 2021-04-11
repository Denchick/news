package repositories

import (
	"github.com/denchick/news/models"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

type FeedsRepository struct {
	db *pg.DB
}

func NewFeedsRepository(db *pg.DB) *FeedsRepository {
	return &FeedsRepository{db}
}

func (repo *FeedsRepository) GetFeeds() ([]*models.DBFeed, error) {
	var feeds []*models.DBFeed
	err := repo.db.Model(&feeds).Select()

	if err != nil {
		return nil, errors.Wrap(err, "store.repository.GetFeeds")
	}
	return feeds, nil
}

func (repo *FeedsRepository) GetSubcategoryFeeds(subcategory *models.DBSubcategory) ([]*models.DBFeed, error) {
	feeds := make([]*models.DBFeed, 0)

	_, err := repo.db.Query(&feeds, `
		SELECT feed_id as id, url, name
		FROM feeds_categories
		JOIN feeds ON feeds.id = feeds_categories.feed_id
		WHERE subcategory_id = ?;
	`, subcategory.ID)
	
	if err != nil {
		return nil, errors.Wrap(err, "store.repository.GetFeedsFromCategory")
	}
	return feeds, nil
}
