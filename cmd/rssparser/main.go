package main

import (
	"log"

	"github.com/denchick/news/cmd/rssparser/feedparser"
	"github.com/denchick/news/internal/config"
	"github.com/denchick/news/internal/logger"
	"github.com/denchick/news/manager"
	"github.com/denchick/news/store"
	"github.com/pkg/errors"
)

// TODO read feed urls from db
func getFeeds() []string {
	return []string{"https://vas3k.ru/rss/"}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg := config.Get()
	lg := logger.Get(cfg.LogLevel)

	// Init store
	s, err := store.NewStore()
	if err != nil {
		return errors.Wrap(err, "store.NewStore")
	}
	
	// Init manager
	m, err := manager.NewManager(s)
	if err != nil {
		return errors.Wrap(err, "manager.NewManager")
	}

	urls := getFeeds()
	feedParser := feedparser.NewFeedParser(lg)
	for _, url := range urls {
		articles, err := feedParser.Parse(url)
		if err != nil {
			lg.Logger.Info().Msgf("Can't parse %s, skipped", url)
			continue
		}
		err = m.News.SaveNews(articles)
		if err != nil {
			lg.Logger.Err(err)
			continue
		}

		lg.Debug().Msgf("Successfully parsed %d articles from %s", len(articles), url)
	}
	return nil
}
