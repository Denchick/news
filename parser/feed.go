package main

import (
	"github.com/mmcdole/gofeed"
	"server/articles"
)

func getFeeds() []string {
	return []string {"https://vas3k.ru/rss/"}
}

func parseFeed(feedUrl string) (articlesArr []articles.Article) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedUrl)
	for _, item := range feed.Items {
		article := articles.Article{
			Title: item.Title,
			Description: item.Description,
			Link: item.Link,
		}
		articlesArr = append(articlesArr, article)
	}
	return
}
