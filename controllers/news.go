package controllers

import (
	"net/http"

	"github.com/denchick/news/manager"
	"github.com/labstack/echo/v4"
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
	return c.NoContent(http.StatusOK)
}