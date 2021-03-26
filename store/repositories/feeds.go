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

func (repo *FeedsRepository) GetFeeds() ([]*models.Feed, error) {
	var feeds []*models.Feed
	err := repo.db.Model(&feeds).Select()

	if err != nil {
		return nil, errors.Wrap(err, "store.repository.GetFeeds")
	}
	return feeds, nil
}

func (repo *FeedsRepository) GetFeedsFromCategory(category *models.Category) ([]*models.Feed, error) {
	feeds := make([]*models.Feed, 0)
	err := repo.db.Model((*models.FeedCategory)(nil)).
		ColumnExpr("feed_id as id, url").
		Join("JOIN feeds ON feeds.id = feed_category.feed_id").
		Where("feed_id = ?", category.ID).
		Select(&feeds)
	if err != nil {
		return nil, errors.Wrap(err, "store.repository.GetFeedsFromCategory")
	}
	return feeds, nil
}
