package main

import (
	"github.com/denchick/news/database"
	"github.com/mmcdole/gofeed"
	"log"
)

func getFeeds() []string {
	return []string{"https://vas3k.ru/rss/", "https://meduza.io/rss/all"}
}

func parseFeed(feedUrl string) (articlesArr []database.Article) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedUrl)
	for _, item := range feed.Items {
		text := item.Description
		if text == "" {
			text = item.Content
		}
		article := createArticle(item.Title, text, item.Link)
		log.Printf("Parsed %s", article.Link)
		articlesArr = append(articlesArr, article)
	}
	return
}
