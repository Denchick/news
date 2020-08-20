package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/denchick/news/database"
	"github.com/mmcdole/gofeed"
	"strings"
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
			Description: getCleanedDescription(item.Description, 500),
			Link: item.Link,
		}
		articlesArr = append(articlesArr, article)
	}
	return
}

func getCleanedDescription(text string, limit int) (result string) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(text))
	text = strings.TrimSpace(doc.Text())
	runes := []rune(text)
	if len(runes) >= limit {
		return string(runes[:limit]) + "..."
	}
	return text
}