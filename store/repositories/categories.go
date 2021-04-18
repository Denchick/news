package repositories

import (
	"github.com/denchick/news/models"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

// CategoriesRepository...
type CategoriesRepository struct {
	db *pg.DB
}

// NewCategoryRepository creates new CategoriesRepository
func NewCategoryRepository(db *pg.DB) *CategoriesRepository {
	return &CategoriesRepository{db}
}

// GetCategory retrieves category by name from Postgres
func (repo *CategoriesRepository) GetCategory(name string) (*models.DBCategory, error) {
	category := new(models.DBCategory)
	err := repo.db.Model(category).Where("name = ?", name).Select()

	if err != nil {
		return nil, errors.Wrap(err, "store.repository.GetCategory")
	}
	return category, nil
}
