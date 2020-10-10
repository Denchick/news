package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/denchick/news/database"
	"github.com/mmcdole/gofeed"
)

func getFeeds() []string {
	return []string{"https://vas3k.ru/rss/", "https://vas3k.club/posts.rss"}
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

func createArticle(title string, description string, link string) database.Article {
	return database.Article{
		Title:       strings.TrimSpace(title),
		Description: getCleanedDescription(description, 500),
		Link:        strings.TrimSpace(link),
	}
}

func getCleanedDescription(text string, limit int) (result string) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(text))
	text = strings.TrimSpace(doc.Text())
	text = strings.ReplaceAll(text, "\n", " ")
	runes := []rune(text)
	if len(runes) >= limit {
		return string(runes[:limit]) + "..."
	}
	return text
}
