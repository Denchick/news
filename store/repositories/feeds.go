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

func (repo *FeedsRepository) GetFeedsFromCategory(category *models.DBCategory) ([]*models.DBFeed, error) {
	feeds := make([]*models.DBFeed, 0)

	err := repo.db.Model((*models.DBFeedCategory)(nil)).
		ColumnExpr("feed_id as id, url, name"). //todo тт ошибка
		Join("JOIN feeds ON feeds.id = db_feed_category.feed_id").
		Where("category_id = ?", category.ID).
		Select(&feeds)
	if err != nil {
		return nil, errors.Wrap(err, "store.repository.GetFeedsFromCategory")
	}
	return feeds, nil
}
