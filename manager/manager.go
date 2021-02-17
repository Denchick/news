package manager

import (
	"errors"

	"github.com/denchick/news/manager/services"
	"github.com/denchick/news/store"
)

// Manager is just a collection of all services we have in the project
type Manager struct {
	News NewsService
}

// NewManager creates new service manager
func NewManager(store *store.Store) (*Manager, error) {
	if store == nil {
		return nil, errors.New("No store provided")
	}
	return &Manager{
		News: services.NewNewsService(store),
	}, nil
}
