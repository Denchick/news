package main

import (
	"github.com/denchick/news/config"
	"github.com/denchick/news/logger"
)

// TODO read feed urls from db
func getFeeds() []string {
	return []string{"https://vas3k.ru/rss/", "https://vas3k.club/posts.rss"}
}

func main() {
	cfg := config.Get()
	lg := logger.Get(cfg.LogLevel)

	// Init store
	// store, err := store.NewStore()
	// if err != nil {
	// 	return errors.Wrap(err, "store.NewStore")
	// }
	
	// Init manager
	// manager, err := service.NewManager(store)
	// if err != nil {
	// 	return errors.Wrap(err, "manager.NewManager")
	// }

	urls := getFeeds()
	feedParser := NewFeedParser(lg)
	for _, url := range urls {
		articles, err := feedParser.Parse(url)
		if err != nil {
			return
		}
		// TODO store save
		lg.Debug().Msgf("Successfully parsed %d articles from %s", len(articles), url)
	}
}
