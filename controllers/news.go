package controllers

import (
	"net/http"

	"github.com/denchick/news/manager"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// NewsController ...
type NewsController struct {
	services *manager.Manager
}

// NewNewsController creates a new news controller
func NewNewsController(services *manager.Manager) *NewsController {
	return &NewsController{services}
}

// Get ...
func (controller *NewsController) Get(c echo.Context) error {
	category := c.QueryParam("category")
	if len(category) == 0 {
		return c.NoContent(http.StatusBadRequest)
	}

	newsBySubcategories, err := controller.services.News.GetNews(category)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			errors.Wrap(err, "could not get news"),
		)
	}
	if len(newsBySubcategories) == 0 {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, newsBySubcategories)
}
