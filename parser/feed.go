package main

import (
	"github.com/denchick/news/database"
	"github.com/mmcdole/gofeed"
)

func getFeeds() []string {
	return []string {"https://vas3k.ru/rss/"}
}

func parseFeed(feedUrl string) (articlesArr []database.Article) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedUrl)
	for _, item := range feed.Items {
		article := database.Article{
			Title: item.Title,
			Description: item.Description,
			Link: item.Link,
		}
		articlesArr = append(articlesArr, article)
	}
	return
}
