package repositories

import (
	"github.com/denchick/news/models"
	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
)

type CategoriesRepository struct {
	db *pg.DB
}

func NewCategoryRepository(db *pg.DB) *CategoriesRepository {
	return &CategoriesRepository{db}
}

func (repo *CategoriesRepository) GetCategory(name string) (*models.DBCategory, error) {
	category := new(models.DBCategory)
	err := repo.db.Model(category).Where("name = ?", name).Select()
	
	if err != nil { 
		return nil, errors.Wrap(err, "store.repository.GetCategory")
	}
	return category, nil
}