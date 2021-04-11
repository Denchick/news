package repositories

import (
	"github.com/denchick/news/models"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

type SubcategoriesRepository struct {
	db *pg.DB
}

func NewSubcategoryRepository(db *pg.DB) *SubcategoriesRepository {
	return &SubcategoriesRepository{db}
}

func (repo *SubcategoriesRepository) GetSubcategories(categoryID uint) ([]*models.DBSubcategory, error) {
	var subcategories []*models.DBSubcategory

	_, err := repo.db.Query(&subcategories, `
		SELECT DISTINCT subcategory_id as id, name
		FROM feeds_categories
		JOIN subcategories on subcategories.id = feeds_categories.subcategory_id
		WHERE category_id = ?;
	`, categoryID)

	if err != nil {
		return nil, errors.Wrap(err, "store.repository.GetSubcategory")
	}
	return subcategories, nil
}
