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

	feeds, err := s.Feeds.GetFeeds()
	if err != nil {
		return errors.Wrap(err, "store.FeedsRepository.GetFeeds")
	}

	lg.Debug().Msgf("Got %d feeds", len(feeds))

	feedParser := feedparser.NewFeedParser(lg)
	for _, feed:= range feeds {
		articles, err := feedParser.Parse(feed)
		if err != nil {
			lg.Logger.Info().Msgf("Can't parse %s, skipped", feed.URL)
			continue
		}
		lg.Debug().Msgf("Parsed %d articles from %s", len(articles), feed.URL)
		err = m.News.SaveNews(articles)
		if err != nil {
			return errors.Wrap(err, "can't save news")
		}
		lg.Debug().Msgf("Saved %d news.", len(articles))

	}
	return nil
}
