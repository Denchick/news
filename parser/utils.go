package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/denchick/news/database"
	"strings"
)

func createArticle(title string, description string, link string) database.Article {
	return database.Article{
		Title: strings.TrimSpace(title),
		Description: getCleanedDescription(description, 500),
		Link: strings.TrimSpace(link),
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
