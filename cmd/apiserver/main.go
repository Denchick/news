package main

import (
	"log"
	"net/http"
	"time"

	"github.com/denchick/news/controllers"
	"github.com/denchick/news/internal/config"
	"github.com/denchick/news/manager"
	"github.com/denchick/news/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Get()

	// Init store
	store, err := store.NewStore()
	if err != nil {
		return errors.Wrap(err, "store.NewStore")
	}
	
	// Init manager
	manager, err := manager.NewManager(store)
	if err != nil {
		return errors.Wrap(err, "manager.NewManager")
	}

	// Initialize Echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Init controllers
	v1 := e.Group("/v1")

	newsController := controllers.NewNewsController(manager)
	newsRoutes := v1.Group("/news")
	newsRoutes.GET("", newsController.Get)

	pingRoutes := v1.Group("/_ping")
	pingRoutes.GET("", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Start server
	s := &http.Server{
		Addr:         cfg.HTTPAddr,
		ReadTimeout:  30 * time.Minute,
		WriteTimeout: 30 * time.Minute,
	}
	e.Logger.Print(e.StartServer(s))

	return nil
}