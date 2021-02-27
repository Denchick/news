package repositories

import (
	"github.com/pkg/errors"
	"github.com/denchick/news/models"
	"github.com/go-pg/pg/v10"
)

// NewsRepository ...
type NewsRepository struct {
	db *pg.DB
}

// NewNewsRepository ...
func NewNewsRepository(db *pg.DB) *NewsRepository {
	return &NewsRepository{db}
}

// BulkCreate ...
func (repo *NewsRepository) BulkCreate(articles []*models.Article) error {
	_, err := repo.db.Model(&articles).Insert(&articles)
	if err != nil {
		return errors.Wrap(err, "store.repositories.BulkCreate")
	}
	return nil
}

// GetByName ... 
func (repo *NewsRepository) GetByName(name string) ([]*models.Article, error) {
	var articles []*models.Article
	err := repo.db.Model(&articles).
		Where("name = ?", name).
		Limit(10).
		Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, errors.Wrap(err, "store.repository.GetByName")
	}
	return articles, nil
}
// GetSimilar ...
func (repo *NewsRepository) GetSimilar(name string) ([]*models.Article, error) {
	var articles []*models.Article
	err := repo.db.Model(&articles).
		OrderExpr("? <-> name", name).
		Limit(3).
		Select()

	if err != nil && err != pg.ErrNoRows {
		return nil, errors.Wrap(err, "store.repository.GetByName")
	}
	return articles, nil
}